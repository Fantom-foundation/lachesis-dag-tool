package main

import (
	"fmt"
	"math"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Fantom-foundation/go-opera/logger"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type GetLogsGenerator struct {
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

func NewGetLogsGenerator(cfg *Config, ks *keystore.KeyStore) *GetLogsGenerator {
	g := &GetLogsGenerator{
		chainId: big.NewInt(cfg.ChainId),
		ks:      ks,

		Instance: logger.New("GetLogsGen"),
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

func (g *GetLogsGenerator) Start() <-chan *Transaction {
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

func (g *GetLogsGenerator) Stop() {
	g.Lock()
	defer g.Unlock()

	if g.done == nil {
		return
	}

	close(g.done)
	g.work.Wait()
	g.done = nil
}

func (g *GetLogsGenerator) getTPS() float64 {
	tps := atomic.LoadUint32(&g.tps)
	return float64(tps)
}

func (g *GetLogsGenerator) SetTPS(tps float64) {
	x := uint32(math.Ceil(tps))
	atomic.StoreUint32(&g.tps, x)
}

func (g *GetLogsGenerator) background(output chan<- *Transaction) {
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

func (g *GetLogsGenerator) Yield() *Transaction {
	if !g.state.IsReady(g.done) {
		return nil
	}
	tx := g.generate(g.position, &g.state)
	g.Log.Info("generated tx", "position", g.position, "dsc", tx.Dsc)
	g.position++

	return tx
}

func (g *GetLogsGenerator) generate(position uint, state *genState) *Transaction {
	count := uint(len(g.accs))
	contractAddr, _ := state.Session.(*common.Address)
	var (
		maker    TxMaker
		callback TxCallback
		dsc      string
	)

	switch step := (position % 10e7); {

	case step == 0:
		dsc = "logemitter contract creation"
		acc := g.accs[position%count]
		maker = g.logemitterCreateContract(acc)
		state.NotReady(dsc)
		callback = func(r *types.Receipt, e error) {
			// r may be nil, then panic
			state.Session = &r.ContractAddress
			state.Ready()
		}

	case 0 < step && step < 1000:
		seed := position % count
		acc := g.accs[seed]
		dsc = fmt.Sprintf("gen some logs %d", seed)
		maker = g.logemitterEmit(acc, *contractAddr, seed)
		break

	default:
		seed := position % count
		acc := g.accs[seed]
		dsc = fmt.Sprintf("get some logs %d", seed)
		maker = g.logemitterGet(acc, *contractAddr, seed)
		break
	}

	return &Transaction{
		Make:     maker,
		Callback: callback,
		Dsc:      dsc,
	}
}

func (g *GetLogsGenerator) Payer(from accounts.Account, amounts ...*big.Int) *bind.TransactOpts {
	t, err := bind.NewKeyStoreTransactorWithChainID(g.ks, from, g.chainId)
	if err != nil {
		panic(err)
	}

	t.Value = big.NewInt(0)
	for _, amount := range amounts {
		t.Value.Add(t.Value, amount)
	}

	return t
}

func (g *GetLogsGenerator) ReadOnly() *bind.CallOpts {
	return &bind.CallOpts{}
}
