package main

import (
	"math"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Fantom-foundation/go-opera/logger"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

type ReadonlyGenerator struct {
	tps     uint32
	chainId *big.Int
	ks      *keystore.KeyStore
	accs    []accounts.Account

	position uint
	state    genState

	work sync.WaitGroup
	done chan struct{}
	sync.Mutex

	logger.Instance
}

func NewReadonlyGenerator(cfg *Config, ks *keystore.KeyStore) *ReadonlyGenerator {
	g := &ReadonlyGenerator{
		chainId: big.NewInt(cfg.ChainId),
		ks:      ks,

		Instance: logger.New("ReadonlyGen"),
	}
	g.state.Log = g.Log

	for _, acc := range ks.Accounts() {
		if acc.Address == cfg.Payer {
			continue
		}
		if err := ks.Unlock(acc, ""); err != nil {
			panic(err)
		}
		g.accs = append(g.accs, acc)
	}

	return g
}

func (g *ReadonlyGenerator) Start() <-chan *Transaction {
	g.Lock()
	defer g.Unlock()

	if g.done != nil {
		return nil
	}
	g.done = make(chan struct{})

	output := make(chan *Transaction, 100)
	g.work.Add(1)
	go g.background(output)

	return output
}

func (g *ReadonlyGenerator) Stop() {
	g.Lock()
	defer g.Unlock()

	if g.done == nil {
		return
	}

	close(g.done)
	g.work.Wait()
	g.done = nil
}

func (g *ReadonlyGenerator) getTPS() float64 {
	tps := atomic.LoadUint32(&g.tps)
	return float64(tps)
}

func (g *ReadonlyGenerator) SetTPS(tps float64) {
	x := uint32(math.Ceil(tps))
	atomic.StoreUint32(&g.tps, x)
}

func (g *ReadonlyGenerator) background(output chan<- *Transaction) {
	defer g.work.Done()
	defer close(output)

	g.Log.Info("started")
	defer g.Log.Info("stopped")

	for {
		begin := time.Now()
		var (
			generating time.Duration
			sending    time.Duration
		)

		tps := g.getTPS()
		for count := tps; count > 0; count-- {
			begin := time.Now()
			tx := g.Yield()
			generating += time.Since(begin)

			begin = time.Now()
			select {
			case output <- tx:
				sending += time.Since(begin)
				continue
			case <-g.done:
				return
			}
		}

		spent := time.Since(begin)
		if spent >= time.Second {
			g.Log.Warn("exceeded performance", "tps", tps, "generating", generating, "sending", sending)
			continue
		}

		select {
		case <-time.After(time.Second - spent):
			continue
		case <-g.done:
			return
		}
	}
}

func (g *ReadonlyGenerator) Yield() *Transaction {
	if !g.state.IsReady(g.done) {
		return nil
	}
	tx := g.generate(g.position, &g.state)
	g.Log.Info("generated tx", "position", g.position, "dsc", tx.Dsc)
	g.position++

	return tx
}

func (g *ReadonlyGenerator) generate(_ uint, _ *genState) *Transaction {
	return &Transaction{
		Make: g.randomSfcCall(),
		Dsc:  "random SFC read",
	}
}
