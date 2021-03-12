package main

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Fantom-foundation/go-lachesis/logger"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BalancesGenerator struct {
	tps     uint32
	chainId *big.Int
	ks      *keystore.KeyStore
	amount  *big.Int
	payer   accounts.Account
	accs    []accounts.Account

	position       uint
	generatorState genState

	work sync.WaitGroup
	done chan struct{}
	sync.Mutex

	logger.Instance
}

func NewBalancesGenerator(cfg *Config, ks *keystore.KeyStore, amount int64) *BalancesGenerator {
	g := &BalancesGenerator{
		chainId: big.NewInt(cfg.ChainId),
		ks:      ks,
		amount:  big.NewInt(amount),

		Instance: logger.MakeInstance(),
	}

	var found bool
	for _, acc := range ks.Accounts() {
		if err := ks.Unlock(acc, ""); err != nil {
			panic(err)
		}

		if acc.Address == cfg.Payer {
			g.payer = acc
			found = true
		} else {
			g.accs = append(g.accs, acc)
		}
	}
	if !found {
		panic("payer not found in the keystore")
	}

	return g
}

func (g *BalancesGenerator) Start() <-chan *Transaction {
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

func (g *BalancesGenerator) Stop() {
	g.Lock()
	defer g.Unlock()

	if g.done == nil {
		return
	}

	close(g.done)
	g.work.Wait()
	g.done = nil
}

func (g *BalancesGenerator) getTPS() float64 {
	tps := atomic.LoadUint32(&g.tps)
	return float64(tps)
}

func (g *BalancesGenerator) SetTPS(tps float64) {
	x := uint32(math.Ceil(tps))
	atomic.StoreUint32(&g.tps, x)
}

func (g *BalancesGenerator) background(output chan<- *Transaction) {
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
			if tx == nil {
				return
			}
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

func (g *BalancesGenerator) Yield() *Transaction {
	if !g.generatorState.IsReady(g.done) {
		return nil
	}
	tx := g.generate(g.position, &g.generatorState)
	if tx == nil {
		return nil
	}
	g.Log.Info("generated tx", "position", g.position, "dsc", tx.Dsc)
	g.position++

	return tx
}

func (g *BalancesGenerator) generate(position uint, state *genState) *Transaction {
	count := uint(len(g.accs))

	if position >= count {
		return nil
	}

	from := g.payer
	to := g.accs[position%count]

	// wait every N
	var callback TxCallback
	if position%500 == 0 {
		state.NotReady("wait for init tx")
		callback = func(r *types.Receipt, e error) {
			if e != nil {
				panic(e)
			}
			state.Ready()
		}
	}

	return &Transaction{
		Make:     g.transferTx(from, to, g.amount),
		Dsc:      fmt.Sprintf("%s --> %s", from.Address.String(), to.Address.String()),
		Callback: callback,
	}
}

func (g *BalancesGenerator) transferTx(from, to accounts.Account, amount *big.Int) TxMaker {
	return func(client *ethclient.Client) (*types.Transaction, error) {
		nonce, err := client.PendingNonceAt(context.Background(), from.Address)
		if err != nil {
			return nil, err
		}

		tx := types.NewTransaction(
			uint64(nonce),
			to.Address,
			amount,
			gasLimit,
			gasPrice,
			[]byte{},
		)

		signed, err := g.ks.SignTx(from, tx, g.chainId)
		if err != nil {
			panic(err)
		}

		err = client.SendTransaction(context.Background(), signed)
		return signed, err
	}
}
