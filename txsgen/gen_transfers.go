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

type TransfersGenerator struct {
	tps     uint32
	chainId *big.Int

	accs     *keystore.KeyStore
	position uint

	work sync.WaitGroup
	done chan struct{}
	sync.Mutex

	logger.Instance
}

func NewTransfersGenerator(cfg *Config, ks *keystore.KeyStore) *TransfersGenerator {
	g := &TransfersGenerator{
		chainId: big.NewInt(cfg.ChainId),
		accs:    ks,

		Instance: logger.MakeInstance(),
	}

	return g
}

func (g *TransfersGenerator) Start() (output chan *Transaction) {
	g.Lock()
	defer g.Unlock()

	if g.done != nil {
		return
	}
	g.done = make(chan struct{})

	output = make(chan *Transaction, 100)
	g.work.Add(1)
	go g.background(output)

	return
}

func (g *TransfersGenerator) Stop() {
	g.Lock()
	defer g.Unlock()

	if g.done == nil {
		return
	}

	close(g.done)
	g.work.Wait()
	g.done = nil
}

func (g *TransfersGenerator) getTPS() float64 {
	tps := atomic.LoadUint32(&g.tps)
	return float64(tps)
}

func (g *TransfersGenerator) SetTPS(tps float64) {
	x := uint32(math.Ceil(tps))
	atomic.StoreUint32(&g.tps, x)
}

func (g *TransfersGenerator) background(output chan<- *Transaction) {
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

func (g *TransfersGenerator) Yield() *Transaction {
	tx := g.generate(g.position)
	g.Log.Info("generated tx", "position", g.position, "dsc", tx.Dsc)
	g.position++

	return tx
}

func (g *TransfersGenerator) generate(position uint) *Transaction {
	accs := g.accs.Accounts()
	count := uint(len(accs))

	var (
		maker    TxMaker
		callback TxCallback
		dsc      string
	)

	from := accs[position%count]
	to := accs[(position+1)%count]

	nonce := position / count

	maker = g.transferTx(from, to, nonce)
	dsc = fmt.Sprintf("%s --> %s", from.Address, to.Address)

	return &Transaction{
		Make:     maker,
		Callback: callback,
		Dsc:      dsc,
	}
}

func (g *TransfersGenerator) transferTx(from, to accounts.Account, nonce uint) TxMaker {
	amount := big.NewInt(1e6)

	tx := types.NewTransaction(
		uint64(nonce),
		to.Address,
		amount,
		gasLimit,
		gasPrice,
		[]byte{},
	)

	signed, err := g.accs.SignTx(from, tx, g.chainId)
	if err != nil {
		panic(err)
	}

	return func(client *ethclient.Client) (*types.Transaction, error) {
		err := client.SendTransaction(context.Background(), signed)
		return signed, err
	}
}
