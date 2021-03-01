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
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type CallsGenerator struct {
	tps     uint32
	chainId uint
	signer  types.Signer

	instances      uint
	accs           []*Acc
	offset         uint
	position       uint
	generatorState genState

	work sync.WaitGroup
	done chan struct{}
	sync.Mutex

	logger.Instance
}

func NewCallsGenerator(cfg *Config, num, ofTotal uint) *CallsGenerator {
	accs := cfg.Accs.Count / ofTotal
	offset := cfg.Accs.Offset + accs*(num-1)
	g := &CallsGenerator{
		chainId:   uint(cfg.ChainId),
		signer:    types.NewEIP155Signer(big.NewInt(int64(cfg.ChainId))),
		instances: ofTotal,
		accs:      make([]*Acc, accs),
		offset:    offset,

		Instance: logger.MakeInstance(),
	}

	return g
}

func (g *CallsGenerator) Start() (output chan *Transaction) {
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
	x := uint32(math.Ceil(tps / float64(g.instances)))
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

type genState struct {
	ready      chan struct{}
	BallotAddr common.Address
}

func (s *genState) NotReady() {
	s.ready = make(chan struct{})
}

func (s *genState) IsReady(done <-chan struct{}) bool {
	if s.ready == nil {
		return true
	}

	select {
	case <-done:
		return false
	case <-s.ready:
		return true
	}
}

func (s *genState) Ready() {
	close(s.ready)
}

func (g *CallsGenerator) generate(position uint, state *genState) *Transaction {
	// count := uint(len(g.accs))
	var (
		maker    TxMaker
		callback TxCallback
		dsc      string
	)

	switch step := (position % 100001); {

	case step == 0:
		dsc = "ballot contract creation"
		maker = g.ballotCreateContract(0)
		state.NotReady()
		callback = func(r *types.Receipt, e error) {
			state.BallotAddr = r.ContractAddress
			state.Ready()
		}

	case 0 < step && step < 100000 && step%2 == 0:
		a := (position / 2) % uint(len(g.accs))
		chose := ballotRandChose()
		dsc = fmt.Sprintf("%d voites for %d", a, chose)
		maker = g.ballotVoite(a, state.BallotAddr, chose)
		break

	case 0 < step && step < 100000 && step%2 == 1:
		a := (position / 2) % uint(len(g.accs))
		dsc = fmt.Sprintf("count voite logs for %d", a)
		maker = g.ballotCountOfVoites(a, state.BallotAddr)
		break

	case step == 100000:
		dsc = "ballot winner reading"
		maker = g.ballotWinner(state.BallotAddr)

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

func (g *CallsGenerator) ReadOnly() *bind.CallOpts {
	return &bind.CallOpts{}
}
