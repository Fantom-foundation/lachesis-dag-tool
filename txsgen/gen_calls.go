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

type CallsGenerator struct {
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

func NewCallsGenerator(cfg *Config, ks *keystore.KeyStore) *CallsGenerator {
	g := &CallsGenerator{
		chainId: big.NewInt(cfg.ChainId),
		ks:      ks,

		Instance: logger.MakeInstance(),
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
	if !g.state.IsReady(g.done) {
		return nil
	}
	tx := g.generate(g.position, &g.state)
	g.Log.Info("generated tx", "position", g.position, "dsc", tx.Dsc)
	g.position++

	return tx
}

func (g *CallsGenerator) generate(position uint, state *genState) *Transaction {
	count := uint(len(g.accs))
	ballotAddr, _ := state.Session.(*common.Address)
	var (
		maker    TxMaker
		callback TxCallback
		dsc      string
	)

	switch step := (position % 100001); {

	case step == 0:
		dsc = "ballot contract creation"
		acc := g.accs[position%count]
		maker = g.ballotCreateContract(acc)
		state.NotReady(dsc)
		callback = func(r *types.Receipt, e error) {
			// r may be nil, then panic
			state.Session = &r.ContractAddress
			state.Ready()
		}

	case 0 < step && step < 100000 && step%2 == 0:
		acc := g.accs[(position/2)%count]
		chose := ballotRandChose()
		dsc = fmt.Sprintf("%s voites for %d", acc.Address.String(), chose)
		maker = g.ballotVoite(acc, *ballotAddr, chose)
		break

	case 0 < step && step < 100000 && step%2 == 1:
		acc := g.accs[(position/2)%count]
		dsc = fmt.Sprintf("count voite logs of %s", acc.Address.String())
		maker = g.ballotCountOfVoites(acc, *ballotAddr)
		break

	case step == 100000:
		dsc = "ballot winner reading"
		maker = g.ballotWinner(*ballotAddr)

	default:
		panic(fmt.Sprintf("unknown step %d", step))
	}

	return &Transaction{
		Make:     maker,
		Callback: callback,
		Dsc:      dsc,
	}
}

func (g *CallsGenerator) Payer(from accounts.Account, amounts ...*big.Int) *bind.TransactOpts {
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

func (g *CallsGenerator) ReadOnly() *bind.CallOpts {
	return &bind.CallOpts{}
}
