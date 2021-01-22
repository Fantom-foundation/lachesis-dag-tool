package source

import (
	"context"
	"fmt"
	"io/ioutil"
	"sync"

	"github.com/Fantom-foundation/go-lachesis/gossip"
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
)

const (
	protocolMaxMsgSize = 10 * 1024 * 1024 // Maximum cap on the size of a protocol message
)

func EventsFromP2p(ctx context.Context, network string, from, to idx.Epoch) <-chan *inter.Event {
	log.Info("Events of epoches", "from", from, "to", to, "network", network)
	output := make(chan *inter.Event, 10)

	svc := newService(network, output)
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
	output  chan<- *inter.Event

	// server
	p2pServer  *p2p.Server
	serverPool *gossip.ServerPool

	peers *gossip.PeerSet

	done chan struct{}
	wg   sync.WaitGroup
}

func newService(network string, output chan<- *inter.Event) *service {
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

		done: make(chan struct{}),
	}

	trustedNodes := []string{}
	db := memorydb.New()
	svc.serverPool = gossip.NewServerPool(db, svc.done, &svc.wg, trustedNodes)

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

	return nil
}

// Stop terminates all goroutines belonging to the service, blocking until they
// are all terminated.
func (s *service) Stop() error {
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
				//case pm.newPeerCh <- peer:
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
	p.Log().Debug("Peer connected", "name", p.Name())

	// Execute the handshake
	var (
		myProgress gossip.PeerProgress // TODO: set
	)
	if err := p.Handshake(s.net.NetworkID, myProgress, s.genesis); err != nil {
		p.Log().Debug("Handshake failed", "err", err)
		return err
	}

	if err := s.peers.Register(p); err != nil {
		p.Log().Warn("Peer registration failed", "err", err)
		return err
	}
	defer func() {
		log.Debug("Removing peer", "name", p.Name())
		if err := s.peers.Unregister(p); err != nil {
			log.Error("Peer removal failed", "name", p.Name(), "err", err)
		}
		p.Peer.Disconnect(p2p.DiscUselessPeer)
	}()

	for {
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

	return nil
}

// NodeInfo retrieves some protocol metadata about the running host node.
func (s *service) NodeInfo() *gossip.NodeInfo {
	return &gossip.NodeInfo{
		Network:     s.net.NetworkID,
		Genesis:     s.genesis,
		Epoch:       1,
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
