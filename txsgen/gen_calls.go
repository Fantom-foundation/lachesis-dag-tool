package main

import (
	"fmt"
	"math"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Fantom-foundation/go-lachesis/logger"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type CallsGenerator struct {
	tps            uint32
	chainId        uint
	accs           *keystore.KeyStore
	position       uint
	generatorState genState

	work sync.WaitGroup
	done chan struct{}
	sync.Mutex

	logger.Instance
}

func NewCallsGenerator(cfg *Config, ks *keystore.KeyStore) *CallsGenerator {
	g := &CallsGenerator{
		chainId: uint(cfg.ChainId),
		accs:    ks,

		Instance: logger.MakeInstance(),
	}

	return g
}

func (g *CallsGenerator) Start() <-chan *Transaction {
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

func (g *CallsGenerator) Stop() {
	g.Lock()
	defer g.Unlock()

	if g.done == nil {
		return
	}

	close(g.done)
	g.work.Wait()
	g.done = nil
}

func (g *CallsGenerator) getTPS() float64 {
	tps := atomic.LoadUint32(&g.tps)
	return float64(tps)
}

func (g *CallsGenerator) SetTPS(tps float64) {
	x := uint32(math.Ceil(tps))
	atomic.StoreUint32(&g.tps, x)
}

func (g *CallsGenerator) background(output chan<- *Transaction) {
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

func (g *CallsGenerator) Yield() *Transaction {
	if !g.generatorState.IsReady(g.done) {
		return nil
	}
	tx := g.generate(g.position, &g.generatorState)
	g.Log.Info("generated tx", "position", g.position, "dsc", tx.Dsc)
	g.position++

	return tx
}

func (g *CallsGenerator) generate(position uint, state *genState) *Transaction {
	accs := g.accs.Accounts()
	count := uint(len(accs))

	var (
		maker    TxMaker
		callback TxCallback
		dsc      string
	)

	ballotAddr := *state.Session.(*common.Address)

	switch step := (position % 100001); {

	case step == 0:
		dsc = "ballot contract creation"
		maker = g.ballotCreateContract(0)
		state.NotReady(dsc)
		callback = func(r *types.Receipt, e error) {
			state.Session = &r.ContractAddress
			state.Ready()
		}

	case 0 < step && step < 100000 && step%2 == 0:
		a := (position / 2) % count
		chose := ballotRandChose()
		dsc = fmt.Sprintf("%d voites for %d", a, chose)
		maker = g.ballotVoite(a, ballotAddr, chose)
		break

	case 0 < step && step < 100000 && step%2 == 1:
		a := (position / 2) % count
		dsc = fmt.Sprintf("count voite logs for %d", a)
		maker = g.ballotCountOfVoites(a, ballotAddr)
		break

	case step == 100000:
		dsc = "ballot winner reading"
		maker = g.ballotWinner(ballotAddr)

	default:
		panic(fmt.Sprintf("unknown step %d", step))
	}

	return &Transaction{
		Make:     maker,
		Callback: callback,
		Dsc:      dsc,
	}
}

func (g *CallsGenerator) Payer(n uint, amounts ...*big.Int) *bind.TransactOpts {
	accs := g.accs.Accounts()
	from := accs[n]

	t, err := bind.NewKeyStoreTransactor(g.accs, from)
	if err != nil {
		panic(err)
	}

	t.Value = big.NewInt(0)
	for _, amount := range amounts {
		t.Value.Add(t.Value, amount)
	}

	return t
}

func (g *CallsGenerator) ReadOnly() *bind.CallOpts {
	return &bind.CallOpts{}
}
