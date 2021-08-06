package main

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/logger"
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Reader struct {
	url    string
	output chan *inter.Event

	done chan struct{}
	work sync.WaitGroup

	logger.Instance
}

func NewReader(url string, start idx.Block) *Reader {
	s := &Reader{
		url:      url,
		done:     make(chan struct{}),
		Instance: logger.MakeInstance(),
	}

	s.work.Add(1)
	go s.background(start)

	return s
}

func (s *Reader) Close() {
	if s.done == nil {
		return
	}
	close(s.done)
	s.done = nil

	s.work.Wait()
}

func (s *Reader) Events() <-chan *inter.Event {
	return s.output
}

func (s *Reader) background(start idx.Block) {
	defer s.work.Done()
	s.Log.Info("started")
	defer s.Log.Info("stopped")

	s.output = make(chan *inter.Event, 10)
	defer close(s.output)

	var (
		client   *ethclient.Client
		err      error
		curBlock = big.NewInt(int64(start))
		maxBlock = big.NewInt(0)
		sbscr    ethereum.Subscription
		headers  = make(chan *types.Header, 1)
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
			client, err = s.connect()
			if err != nil {
				disconnect()
				delay()
				continue
			}
			sbscr, err = s.subscribe(client, headers)
			if err != nil {
				disconnect()
				delay()
				continue
			}
		}

		for curBlock.Cmp(maxBlock) <= 0 {
			err = s.readEvents(curBlock, client)
			if err != nil {
				disconnect()
				delay()
				continue
			}
			curBlock.Add(curBlock, big.NewInt(1))
		}

		// wait for nex task
		select {
		case b := <-headers:
			if maxBlock.Cmp(b.Number) < 0 {
				maxBlock.Set(b.Number)
			}
		case <-s.done:
			return
		}
	}
}

func (s *Reader) readEvents(n *big.Int, client *ethclient.Client) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	blk, err := client.BlockByNumber(ctx, n)
	if err != nil {
		s.Log.Error("new block", "block", n, "err", err)
		return
	}
	s.Log.Info("new block", "block", n)

	atropos := hash.Event(blk.Hash())
	s.Log.Info("new atropos", "id", atropos)
	e, err := client.GetEventPayload(ctx, atropos.String(), true)
	if err != nil {
		panic(err)
	}

	data, _ := e.MarshalJSON()
	fmt.Println(string(data))

	// TODO: extract all the events
	//

	return
}

func (s *Reader) connect() (*ethclient.Client, error) {
	client, err := ethclient.Dial(s.url)
	if err != nil {
		s.Log.Error("connect to", "url", s.url, "err", err)
		return nil, err
	}
	s.Log.Info("connect to", "url", s.url)
	return client, nil
}

func (s *Reader) subscribe(client *ethclient.Client, headers chan *types.Header) (sbscr ethereum.Subscription, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	try(func() error {
		sbscr, err = client.SubscribeNewHead(ctx, headers)
		return err
	})
	if err != nil {
		s.Log.Error("subscribe to", "url", s.url, "err", err)
		return
	}
	s.Log.Info("subscribe to", "url", s.url)
	return
}

func delay() {
	<-time.After(2 * time.Second)
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
