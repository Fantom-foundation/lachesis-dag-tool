package source

import (
	"context"
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
			/*
				select {
				case pm.newPeerCh <- peer:
					pm.wg.Add(1)
					defer pm.wg.Done()
					err := pm.handle(peer)
					if entry != nil {
						s.serverPool.Disconnect(entry)
					}
					return err
				case <-pm.quitSync:
					if entry != nil {
						s.serverPool.Disconnect(entry)
					}
					return p2p.DiscQuitting
				}
			*/
			log.Warn("Connect to node", "addr", peer.RemoteAddr().String())
			if entry != nil {
				s.serverPool.Disconnect(entry)
			}
			return p2p.DiscQuitting
		},
		NodeInfo: func() interface{} {
			return s.NodeInfo()
		},
		PeerInfo: func(id enode.ID) interface{} {
			return nil
		},
	}
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
