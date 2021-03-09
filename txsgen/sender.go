package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Fantom-foundation/go-lachesis/evmcore"
	"github.com/Fantom-foundation/go-lachesis/logger"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Sender struct {
	url       string
	input     chan *Transaction
	callbacks map[common.Hash]TxCallback
	headers   chan *types.Header

	done chan struct{}
	work sync.WaitGroup

	logger.Instance
}

func NewSender(url string) *Sender {
	s := &Sender{
		url:       url,
		input:     make(chan *Transaction, 10),
		callbacks: make(map[common.Hash]TxCallback),
		headers:   make(chan *types.Header, 1),
		done:      make(chan struct{}),

		Instance: logger.MakeInstance(),
	}

	s.work.Add(1)
	go s.background()

	return s
}

func (s *Sender) Close() {
	if s.done == nil {
		return
	}
	close(s.done)
	s.done = nil

	s.work.Wait()
	close(s.input)
}

func (s *Sender) Send(tx *Transaction) {
	s.input <- tx
}

func (s *Sender) background() {
	defer s.work.Done()
	s.Log.Info("started")
	defer s.Log.Info("stopped")

	var (
		client *ethclient.Client
		err    error
		tx     *Transaction
		sbscr  ethereum.Subscription
	)

	disconnect := func() {
		if sbscr != nil {
			sbscr.Unsubscribe()
			sbscr = nil
		}
		if client != nil {
			client.Close()
			client = nil
			s.Log.Error("disonnect from", "url", s.url)
		}
	}
	defer disconnect()

	for {
		// client connect
		for client == nil {
			client = s.connect()
			sbscr = s.subscribe(client)
			if sbscr == nil {
				disconnect()
			}
			select {
			case <-s.done:
				return
			case <-time.After(time.Second):
			}
		}

		// input header
		for tx == nil {
			select {
			case b := <-s.headers:
				err = s.onNewHeader(client, b)
				if err != nil {
					disconnect()
				}
			case <-s.done:
				return
			case tx = <-s.input:
			}
		}

		// output tx

		var (
			t      *types.Transaction
			txHash common.Hash
		)
		err = try(func() error {
			t, err = tx.Make(client)
			return err
		})
		if t != nil {
			txHash = t.Hash()
		}
		if err == nil {
			if tx.Callback != nil {
				s.callbacks[txHash] = tx.Callback
			}
			txCountSentMeter.Inc(1)
			s.Log.Info("tx sending ok", "hash", txHash, "dsc", tx.Dsc)
			tx = nil
			continue
		}

		if tx.Callback != nil {
			tx.Callback(nil, err)
		}

		switch err.Error() {
		case fmt.Sprintf("known transaction: %x", txHash),
			evmcore.ErrNonceTooLow.Error(),
			evmcore.ErrReplaceUnderpriced.Error():
			s.Log.Warn("tx sending skip", "hash", txHash, "dsc", tx.Dsc, "cause", err)
			tx = nil
			continue
		default:
			s.Log.Error("tx sending err", "hash", txHash, "dsc", tx.Dsc, "err", err)
			disconnect()
			s.delay()
			continue
		}
	}
}

func (s *Sender) connect() *ethclient.Client {
	client, err := ethclient.Dial(s.url)
	if err != nil {
		s.Log.Error("connect to", "url", s.url, "err", err)
		s.delay()
		return nil
	}
	s.Log.Info("connect to", "url", s.url)
	return client
}

func (s *Sender) subscribe(client *ethclient.Client) ethereum.Subscription {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var (
		sbscr ethereum.Subscription
		err   error
	)
	try(func() error {
		sbscr, err = client.SubscribeNewHead(ctx, s.headers)
		return err
	})
	if err != nil {
		s.Log.Error("subscribe to", "url", s.url, "err", err)
		s.delay()
		return nil
	}
	s.Log.Info("subscribe to", "url", s.url)
	return sbscr
}

func (s *Sender) onNewHeader(client *ethclient.Client, h *types.Header) (err error) {
	b := evmcore.ConvertFromEthHeader(h)
	s.Log.Debug("new block", "number", b.Number, "block", b.Hash)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var txsCount uint
	err = try(func() error {
		txsCount, err = client.TransactionCount(ctx, b.Hash)
		return err
	})
	if err != nil {
		s.Log.Error("block txs", "number", b.Number, "block", b.Hash, "err", err)
		return
	}
	s.Log.Debug("block txs", "number", b.Number, "block", b.Hash, "count", txsCount)

	for index := uint(0); index < txsCount; index++ {
		var tx *types.Transaction
		err = try(func() error {
			tx, err = client.TransactionInBlock(ctx, b.Hash, index)
			return err
		})
		if err != nil {
			s.Log.Error("tx of block", "number", b.Number, "block", b.Hash, "index", index, "err", err)
			return
		}
		txHash := tx.Hash()

		callback := s.callbacks[txHash]
		if callback == nil {
			continue
		}

		var r *types.Receipt
		err = try(func() error {
			r, err = client.TransactionReceipt(ctx, txHash)
			return err
		})
		if err != nil {
			s.Log.Error("new receipt", "number", b.Number, "block", b.Hash, "index", index, "tx", txHash, "err", err)
			return err
		}
		s.Log.Debug("new receipt", "number", b.Number, "block", b.Hash, "index", index, "tx", txHash)

		callback(r, nil)
	}

	return nil
}

func (s *Sender) delay() {
	select {
	case <-time.After(2 * time.Second):
		return
	case <-s.done:
		return
	}
}

func try(f func() error) (err error) {
	defer func() {
		if catch := recover(); catch != nil {
			err = fmt.Errorf("client panic: %v", catch)
		}
	}()

	err = f()
	return
}
