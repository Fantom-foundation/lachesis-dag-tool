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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/p2p/discv5"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"github.com/ethereum/go-ethereum/p2p/enr"
	"github.com/ethereum/go-ethereum/p2p/nat"
	"github.com/ethereum/go-ethereum/params"
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
		Version: params.Version(),
		P2P: p2p.Config{
			NoDiscovery: true,
			DiscoveryV5: true,
			ListenAddr:  ":5051",
			MaxPeers:    10,
			NAT:         nat.Any(),
			BootstrapNodesV5: parseBootstrapNodes([]string{
				// mainnet
				"enode://a2941866e485442aa6b17d67d77f8a6c4580bb556894cc1618473eff1e18203d8cce50b563cf4c75e408886079b8f067069442ed52e2ac9e556baa3f8fcc525f@3.24.15.66:5050",
				"enode://9ea4e8ad12e1fcfc846d00a626fbf3ca1ee6a5143148b8472dd62a64b876d01031c064a717aaa606de34bcbbb691b50cbdf816acd9e9ee7ce334082b739829b9@52.45.126.213:5050",
				"enode://cba8bd4179052908d069c6bacccf999dd576fdfe6fd29db19b21ae262d96a9ac00a3f5904b772e44b172575b8afbf91c0d1303f1b7b03dfae14e50e234004d36@15.164.136.219:5050",
				"enode://fb904114975d7b238c2d5e46824ac00fdd1133f836e6e5332765e089e70c6dbd4d47c513fcb2ee72774484216634cc9785a54816bfde46d721871d340070b64b@34.245.17.87:5050",
				// testnet
				"enode://563b30428f48357f31c9d4906ca2f3d3815d663b151302c1ba9d58f3428265b554398c6fabf4b806a49525670cd9e031257c805375b9fdbcc015f60a7943e427@3.213.142.230:7946",
				"enode://8b53fe4410cde82d98d28697d56ccb793f9a67b1f8807c523eadafe96339d6e56bc82c0e702757ac5010972e966761b1abecb4935d9a86a9feed47e3e9ba27a6@3.227.34.226:7946",
				"enode://1703640d1239434dcaf010541cafeeb3c4c707be9098954c50aa705f6e97e2d0273671df13f6e447563e7d3a7c7ffc88de48318d8a3cc2cc59d196516054f17e@52.72.222.228:7946",
			}),
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
