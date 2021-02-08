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
	"github.com/ethereum/go-ethereum/core/types"
)

type TransfersGenerator struct {
	tps     uint32
	chainId uint
	signer  types.Signer

	instances uint
	accs      []*Acc
	offset    uint
	position  uint

	work sync.WaitGroup
	done chan struct{}
	sync.Mutex

	logger.Instance
}

func NewTransfersGenerator(cfg *Config, num, ofTotal uint) *TransfersGenerator {
	accs := cfg.Accs.Count / ofTotal
	offset := cfg.Accs.Offset + accs*(num-1)
	g := &TransfersGenerator{
		chainId:   uint(cfg.ChainId),
		signer:    types.NewEIP155Signer(big.NewInt(int64(cfg.ChainId))),
		instances: ofTotal,
		accs:      make([]*Acc, accs),
		offset:    offset,

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

	g.Log.Info("will use", "accounts", len(g.accs), "from", g.offset, "to", uint(len(g.accs))+g.offset)
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
	x := uint32(math.Ceil(tps / float64(g.instances)))
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
	var (
		maker    TxMaker
		callback TxCallback
		dsc      string
	)

	count := uint(len(g.accs))
	a := position % count
	b := (position + 1) % count

	from := g.accs[a]
	if from == nil {
		from = MakeAcc(a + g.offset)
		g.accs[a] = from
	}
	a += g.offset

	to := g.accs[b]
	if to == nil {
		to = MakeAcc(b + g.offset)
		g.accs[b] = to
	}
	b += g.offset

	nonce := position / count

	maker = g.transferTx(from, to, nonce)
	dsc = fmt.Sprintf("%d-->%d", a, b)

	return &Transaction{
		Make:     maker,
		Callback: callback,
		Dsc:      dsc,
	}
}

func (g *TransfersGenerator) Payer(n uint, amounts ...*big.Int) *bind.TransactOpts {
	from := g.accs[n]
	if from == nil {
		from = MakeAcc(n + g.offset)
		g.accs[n] = from
	}

	t := bind.NewKeyedTransactor(from.Key)

	t.Value = big.NewInt(0)
	for _, amount := range amounts {
		t.Value.Add(t.Value, amount)
	}

	return t
}

func (g *TransfersGenerator) ReadOnly() *bind.CallOpts {
	return &bind.CallOpts{}
}
