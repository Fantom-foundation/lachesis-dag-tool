package main

import (
	"fmt"
	"time"

	"github.com/Fantom-foundation/go-lachesis/logger"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Fantom-foundation/lachesis-dag-tool/txsgen/utils"
)

// Nodes pool.
type Nodes struct {
	tps      chan float64
	receipts chan int64
	conns    []*Sender
	Done     chan struct{}
	cfg      *Config
	logger.Instance
}

func NewNodes(cfg *Config, input <-chan *Transaction) *Nodes {
	n := &Nodes{
		tps:      make(chan float64, 1),
		receipts: make(chan int64, 10),
		Done:     make(chan struct{}),
		cfg:      cfg,
		Instance: logger.MakeInstance(),
	}

	for _, url := range cfg.URLs {
		n.add(url)
	}

	n.notifyTPS(0)
	go n.background(input)
	go n.measureTPS()
	return n
}

func (n *Nodes) Count() int {
	return len(n.conns)
}

func (n *Nodes) TPS() <-chan float64 {
	return n.tps
}

func (n *Nodes) notifyTPS(tps float64) {
	select {
	case n.tps <- tps:
		break
	default:
	}
}

func (n *Nodes) measureTPS() {
	var (
		avgbuff       = utils.NewAvgBuff(50)
		txCount int64 = 0
		start         = time.Now()
	)
	for got := range n.receipts {
		txCountGotMeter.Inc(int64(got))
		txCount += got

		dur := time.Since(start).Seconds()
		if dur < 5.0 && txCount < 100 {
			continue
		}

		tps := float64(txCount) / dur
		avgbuff.Push(float64(txCount), dur)
		txTpsMeter.Update(int64(tps))

		avg := avgbuff.Avg()
		n.notifyTPS(avg)
		n.Log.Info("TPS", "current", tps, "avg", avg)

		start = time.Now()
		txCount = 0
	}
}

func (n *Nodes) add(url string) {
	c := NewSender(url)
	c.SetName(fmt.Sprintf("Node-%d", len(n.conns)))
	n.conns = append(n.conns, c)
}

func (n *Nodes) background(input <-chan *Transaction) {
	defer close(n.Done)

	if len(n.conns) < 1 {
		panic("no connections")
	}

	i := 0
	for tx := range input {
		if tx == nil {
			continue
		}
		c := n.conns[i]
		c.Send(n.wrapWithCounter(tx))
		i = (i + 1) % len(n.conns)
	}

	for _, c := range n.conns {
		c.Close()
	}
}

func (n *Nodes) wrapWithCounter(tx *Transaction) *Transaction {
	callback := tx.Callback
	tx.Callback = func(r *types.Receipt, e error) {
		if r != nil {
			n.receipts <- 1
		}
		if callback != nil {
			callback(r, e)
		}
	}

	return tx
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
