package source

import (
	"context"
	"fmt"
	"io/ioutil"
	"sync"
	"time"

	"github.com/Fantom-foundation/go-lachesis/gossip"
	"github.com/Fantom-foundation/go-lachesis/gossip/fetcher"
	"github.com/Fantom-foundation/go-lachesis/gossip/ordering"
	"github.com/Fantom-foundation/go-lachesis/gossip/packsdownloader"
	"github.com/Fantom-foundation/go-lachesis/hash"
	"github.com/Fantom-foundation/go-lachesis/inter"
	"github.com/Fantom-foundation/go-lachesis/inter/idx"
	"github.com/Fantom-foundation/go-lachesis/kvdb/memorydb"
	"github.com/Fantom-foundation/go-lachesis/lachesis"
	"github.com/Fantom-foundation/go-lachesis/lachesis/params"
	"github.com/Fantom-foundation/go-lachesis/version"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discv5"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/enr"
	"github.com/ethereum/go-ethereum/p2p/nat"

	"github.com/Fantom-foundation/lachesis-dag-tool/neo4j"
)

const (
	// the maximum cap on the size of a protocol message
	protocolMaxMsgSize = 10 * 1024 * 1024
	// the maximum number of events in the ordering buffer
	eventsBuffSize = 2048

	softResponseLimitSize = 2 * 1024 * 1024    // Target maximum size of returned events, or other data.
	softLimitItems        = 250                // Target maximum number of events or transactions to request/response
	hardLimitItems        = softLimitItems * 4 // Maximum number of events or transactions to request/response

	maxPackSize      = softResponseLimitSize
	maxPackEventsNum = softLimitItems
)

func EventsFromP2p(ctx context.Context, network string, from, to idx.Epoch, store *neo4j.Store) <-chan *neo4j.EventData {
	log.Info("Events of epoches", "from", from, "to", to, "network", network)
	output := make(chan *neo4j.EventData, 1)

	svc := newService(network, output, from, to, store)
	stack := newStack()
	stack.RegisterProtocols(svc.Protocols())
	stack.RegisterLifecycle(svc)
	svc.p2pServer = stack.Server()

	go func() {
		defer close(output)

		err := stack.Start()
		if err != nil {
			log.Error("Error starting protocol stack", "err", err)
			return
		}
		defer stack.Close()

		<-ctx.Done()
	}()

	return output
}

func newStack() *node.Node {
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		panic(err)
	}

	cfg := node.Config{
		DataDir: tmpDir,
		Name:    "lachesis-listener",
		Version: version.AsString(),
		P2P: p2p.Config{
			NoDiscovery:      true,
			DiscoveryV5:      true,
			ListenAddr:       ":5051",
			MaxPeers:         10,
			NAT:              nat.Any(),
			BootstrapNodesV5: parseBootstrapNodes(params.Bootnodes),
		},
	}

	stack, err := node.New(&cfg)
	if err != nil {
		panic(err)
	}

	return stack
}

type service struct {
	net     lachesis.Config
	genesis common.Hash
	output  chan<- *neo4j.EventData

	// server
	p2pServer  *p2p.Server
	serverPool *gossip.ServerPool

	peers      *gossip.PeerSet
	downloader *packsdownloader.PacksDownloader
	fetcher    *fetcher.Fetcher
	buffer     *ordering.EventBuffer
	store      *neo4j.Store
	status     *status

	done chan struct{}
	wg   sync.WaitGroup
}

func newService(network string, output chan<- *neo4j.EventData, from, to idx.Epoch, store *neo4j.Store) *service {
	currEpoch := store.GetEpoch("current")
	if from < currEpoch {
		from = currEpoch
	}

	var (
		net     lachesis.Config
		genesis common.Hash
	)
	switch network {
	case "test":
		net = lachesis.TestNetConfig()
		genesis = common.HexToHash("0x5e521b3244c5ffb02f5c54e8f69e6441da12cb7ed8a73d8a1ddcb94008acd597")
	case "main":
		net = lachesis.MainNetConfig()
		genesis = common.HexToHash("0xe7813c633b1023bcc132f45273a69df9f3847f16b3209df08517d20bcef3d580")
	default:
		panic("unknow network " + network)
	}

	svc := &service{
		net:     net,
		genesis: genesis,
		output:  output,

		peers: gossip.NewPeerSet(),
		store: store,

		status: newStatus(from),

		done: make(chan struct{}),
	}

	trustedNodes := []string{}
	db := memorydb.New()
	svc.serverPool = gossip.NewServerPool(db, svc.done, &svc.wg, trustedNodes)

	svc.buffer = ordering.New(eventsBuffSize, ordering.Callback{
		Process: func(event *inter.Event) error {
			if from < event.Epoch {
				from = event.Epoch
				store.SetEpoch("current", from)
			}
			if to > 0 && to < event.Epoch {
				close(output)
				return nil
			}

			var wg sync.WaitGroup
			wg.Add(1)
			output <- &neo4j.EventData{Event: event, Ready: wg.Done}
			wg.Wait()

			if svc.status.IsEpochSealedBy(event.Hash()) {
				peerEpoch := func(peer string) idx.Epoch {
					p := svc.peers.Peer(peer)
					if p == nil {
						return 0
					}
					return p.Progress.Epoch
				}
				svc.downloader.OnNewEpoch(svc.status.CurrEpoch(), peerEpoch)
			}

			return nil
		},
		Drop: func(event *inter.Event, peer string, err error) {
		},
		Exists: func(id hash.Event) bool {
			return svc.store.HasEventHeader(id)
		},
		Get: func(id hash.Event) *inter.EventHeaderData {
			event := svc.store.GetEvent(id)
			if event != nil {
				log.Debug("read from db", "event", id, "parents", event.Parents)
			} else {
				log.Debug("read from db", "event", id, "parents", "not found!")
			}
			return event
		},
		Check: func(event *inter.Event, parents []*inter.EventHeaderData) error {
			return nil
		},
	})

	svc.fetcher = fetcher.New(fetcher.Callback{
		PushEvent: func(e *inter.Event, peer string) {
			// log.Info("+++", "event", e.Hash(), "parents", e.Parents)
			svc.buffer.PushEvent(e, peer)
		},
		OnlyInterested: svc.onlyInterestedEvents,
		DropPeer:       svc.removePeer,
		FirstCheck:     func(*inter.Event) error { return nil },
		HeavyCheck:     nil,
	})
	svc.downloader = packsdownloader.New(svc.fetcher, svc.onlyNotConnectedEvents, svc.removePeer)

	return svc
}

// Start is called after all services have been constructed and the networking
// layer was also initialized to spawn any goroutines required by the service.
func (s *service) Start() error {
	topic := discv5.Topic("lachesis@" + s.genesis.Hex())
	if s.p2pServer.DiscV5 != nil {
		go func(topic discv5.Topic) {
			log.Info("Starting topic registration")
			defer log.Info("Terminated topic registration")

			s.p2pServer.DiscV5.RegisterTopic(topic, s.done)
		}(topic)
	}

	s.serverPool.Start(s.p2pServer, topic)
	s.fetcher.Start()

	return nil
}

// Stop terminates all goroutines belonging to the service, blocking until they
// are all terminated.
func (s *service) Stop() error {
	s.downloader.Terminate()
	s.fetcher.Stop()

	close(s.done)
	s.wg.Wait()
	return nil
}

// Protocols returns protocols the service can communicate on.
func (s *service) Protocols() []p2p.Protocol {
	protos := make([]p2p.Protocol, len(gossip.ProtocolVersions))
	for i, vsn := range gossip.ProtocolVersions {
		protos[i] = s.makeProtocol(vsn)
		protos[i].Attributes = []enr.Entry{&gossip.Enr{}}
	}
	return protos
}

func (s *service) makeProtocol(version uint) p2p.Protocol {
	return p2p.Protocol{
		Name:    "lachesis",
		Version: version,
		Length:  gossip.PackMsg + 1,
		Run: func(p *p2p.Peer, rw p2p.MsgReadWriter) error {
			var entry *gossip.PoolEntry
			peer := gossip.NewPeer(int(version), p, rw)
			if s.serverPool != nil {
				entry = s.serverPool.Connect(peer, peer.Node())
			}
			peer.PoolEntry = entry

			select {
			case <-s.done:
				if entry != nil {
					s.serverPool.Disconnect(entry)
				}
				return p2p.DiscQuitting
			default:
				s.wg.Add(1)
				defer s.wg.Done()
				err := s.handle(peer)
				if entry != nil {
					s.serverPool.Disconnect(entry)
				}
				return err
			}
		},
		NodeInfo: func() interface{} {
			return s.NodeInfo()
		},
		PeerInfo: func(id enode.ID) interface{} {
			return nil
		},
	}
}

// handle is the callback invoked to manage the life cycle of a peer. When
// this function terminates, the peer is disconnected.
func (s *service) handle(p *gossip.Peer) error {
	if s.peers.Len() >= 3 {
		return p2p.DiscTooManyPeers
	}
	p.Log().Info("Peer connected", "name", p.Name())

	progress := gossip.PeerProgress{Epoch: s.status.CurrEpoch()}
	if err := p.Handshake(s.net.NetworkID, progress, s.genesis); err != nil {
		p.Log().Debug("Handshake failed", "err", err)
		return err
	}

	if err := s.peers.Register(p); err != nil {
		p.Log().Warn("Peer registration failed", "err", err)
		return err
	}
	defer s.removePeer(p.Uid)

	for {
		select {
		case <-s.done:
			return p2p.DiscQuitting
		default:
		}
		if err := s.handleMsg(p); err != nil {
			p.Log().Debug("Message handling failed", "err", err)
			return err
		}
	}
}

// handleMsg is invoked whenever an inbound message is received from a remote
// peer. The remote connection is torn down upon returning any error.
func (s *service) handleMsg(p *gossip.Peer) error {
	// Read the next message from the remote peer, and ensure it's fully consumed
	msg, err := p.ReadMsg()
	if err != nil {
		return err
	}
	if msg.Size > protocolMaxMsgSize {
		return errResp(gossip.ErrMsgTooLarge, "%v > %v", msg.Size, protocolMaxMsgSize)
	}
	defer msg.Discard()

	peerDwnlr := s.downloader.Peer(p.Uid)

	switch msg.Code {

	case gossip.EthStatusMsg:
		// Status messages should never arrive after the handshake
		return errResp(gossip.ErrExtraStatusMsg, "uncontrolled status message")
	case gossip.NewEventHashesMsg:
		break
	case gossip.EvmTxMsg:
		break
	case gossip.GetEventsMsg:
		break
	case gossip.GetPackInfosMsg:
		break
	case gossip.GetPackMsg:
		break

	case gossip.ProgressMsg:
		var progress gossip.PeerProgress
		if err := msg.Decode(&progress); err != nil {
			return errResp(gossip.ErrDecode, "%v: %v", msg, err)
		}
		p.SetProgress(progress)

		// notify downloader about new peer's epoch
		_ = s.downloader.RegisterPeer(packsdownloader.Peer{
			ID:               p.Uid,
			Epoch:            p.Progress.Epoch,
			RequestPack:      p.RequestPack,
			RequestPackInfos: p.RequestPackInfos,
		}, s.status.CurrEpoch())
		peerDwnlr := s.downloader.Peer(p.Uid)
		if peerDwnlr != nil && progress.LastPackInfo.Index > 0 {
			_ = peerDwnlr.NotifyPackInfo(p.Progress.Epoch, progress.LastPackInfo.Index, progress.LastPackInfo.Heads, time.Now())
		}

	case gossip.PackInfosMsg:
		if peerDwnlr == nil {
			break
		}

		var infos gossip.PackInfosData
		if err := msg.Decode(&infos); err != nil {
			return errResp(gossip.ErrDecode, "%v: %v", msg, err)
		}
		if err := checkLenLimits(len(infos.Infos), infos); err != nil {
			return err
		}

		log.Warn("GOT", "infos", fmt.Sprintf("%#v", infos))

		// notify about number of packs this peer has
		_ = peerDwnlr.NotifyPacksNum(infos.Epoch, infos.TotalNumOfPacks)

		for _, info := range infos.Infos {
			if len(info.Heads) == 0 {
				return errResp(gossip.ErrEmptyMessage, "%v", msg)
			}
			s.status.AddHeaders(infos.Epoch, info.Heads)
			// Mark the hashes as present at the remote node
			for _, id := range info.Heads {
				p.MarkEvent(id)
			}
			// Notify downloader about new packInfo
			_ = peerDwnlr.NotifyPackInfo(infos.Epoch, info.Index, info.Heads, time.Now())
		}

	case gossip.PackMsg:
		if peerDwnlr == nil {
			break
		}

		var pack gossip.PackData
		if err := msg.Decode(&pack); err != nil {
			return errResp(gossip.ErrDecode, "%v: %v", msg, err)
		}
		if err := checkLenLimits(len(pack.IDs), pack); err != nil {
			return err
		}
		if len(pack.IDs) == 0 {
			return errResp(gossip.ErrDecode, "%v: %v", msg, err)
		}

		// Mark the hashes as present at the remote node
		for _, id := range pack.IDs {
			p.MarkEvent(id)
		}
		// Notify downloader about new pack
		_ = peerDwnlr.NotifyPack(pack.Epoch, pack.Index, pack.IDs, time.Now(), p.RequestEvents)

	case gossip.EventsMsg:
		if s.fetcher.Overloaded() {
			break
		}
		var events []*inter.Event
		if err := msg.Decode(&events); err != nil {
			return errResp(gossip.ErrDecode, "%v: %v", msg, err)
		}
		if err := checkLenLimits(len(events), events); err != nil {
			return err
		}

		// Mark the hashes as present at the remote node
		for _, e := range events {
			p.MarkEvent(e.Hash())
		}

		_ = s.fetcher.Enqueue(p.Uid, events, time.Now(), p.RequestEvents)

	default:
		err := errResp(gossip.ErrInvalidMsgCode, "%v", msg.Code)
		return err

	}

	return nil
}

func (s *service) removePeer(id string) {
	// Short circuit if the peer was already removed
	peer := s.peers.Peer(id)
	if peer == nil {
		return
	}
	log.Debug("Removing peer", "peer", id)

	// Unregister the peer from the downloader and peer set
	_ = s.downloader.UnregisterPeer(id)
	if err := s.peers.Unregister(id); err != nil {
		log.Error("Peer removal failed", "peer", id, "err", err)
	}
	// Hard disconnect at the networking layer
	if peer != nil {
		peer.Peer.Disconnect(p2p.DiscUselessPeer)
	}
}

func (s *service) onlyNotConnectedEvents(ids hash.Events) hash.Events {
	if len(ids) == 0 {
		return ids
	}
	notConnected := make(hash.Events, 0, len(ids))
	for _, id := range ids {
		if s.store.HasEventHeader(id) {
			continue
		}
		notConnected.Add(id)
	}
	return notConnected
}

func (s *service) onlyInterestedEvents(ids hash.Events) hash.Events {
	if len(ids) == 0 {
		return ids
	}

	epoch := s.status.CurrEpoch()

	interested := make(hash.Events, 0, len(ids))
	for _, id := range ids {
		if id.Epoch() != epoch {
			continue
		}
		if s.buffer.IsBuffered(id) || s.store.HasEventHeader(id) {
			continue
		}
		interested.Add(id)
	}
	return interested
}

// NodeInfo retrieves some protocol metadata about the running host node.
func (s *service) NodeInfo() *gossip.NodeInfo {
	return &gossip.NodeInfo{
		Network:     s.net.NetworkID,
		Genesis:     s.genesis,
		Epoch:       s.status.CurrEpoch(),
		NumOfBlocks: 0,
	}
}

func parseBootstrapNodes(urls []string) []*discv5.Node {
	nodes := make([]*discv5.Node, 0, len(urls))
	for _, url := range urls {
		if url == "" {
			continue
		}
		node, err := discv5.ParseNode(url)
		if err != nil {
			log.Error("Bootstrap URL invalid", "enode", url, "err", err)
			continue
		}
		nodes = append(nodes, node)
	}
	return nodes
}

func errResp(code int, format string, v ...interface{}) error {
	return fmt.Errorf("%v - %v", code, fmt.Sprintf(format, v...))
}

func checkLenLimits(size int, v interface{}) error {
	if size <= 0 {
		return errResp(gossip.ErrEmptyMessage, "%v", v)
	}
	if size > hardLimitItems {
		return errResp(gossip.ErrMsgTooLarge, "%v", v)
	}
	return nil
}
