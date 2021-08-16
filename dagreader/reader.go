package main

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/Fantom-foundation/go-opera/ftmclient"
	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/logger"
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Fantom-foundation/lachesis-dag-tool/dagreader/internal"
)

type Reader struct {
	url    string
	output chan internal.ToStore

	done chan struct{}
	work sync.WaitGroup

	logger.Instance
}

func NewReader(url string, start idx.Block) *Reader {
	r := &Reader{
		url:      url,
		output:   make(chan internal.ToStore, 10),
		done:     make(chan struct{}),
		Instance: logger.MakeInstance(),
	}

	r.SetName("reader")

	r.work.Add(1)
	go r.background(start)

	return r
}

func (r *Reader) Close() {
	if r.done == nil {
		return
	}
	close(r.done)
	r.work.Wait()
	r.done = nil
}

func (s *Reader) Events() <-chan internal.ToStore {
	return s.output
}

func (r *Reader) background(start idx.Block) {
	defer r.work.Done()
	defer close(r.output)
	r.Log.Info("started")
	defer r.Log.Info("stopped")

	var (
		client   *ftmclient.Client
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
			r.Log.Error("disonnect from", "url", r.url)
		}
	}
	defer disconnect()

	was := make(map[hash.Event]struct{})

	for {
		// client connect
		for client == nil {
			select {
			case <-r.done:
				return
			default:
			}
			client, err = r.connect()
			if err != nil {
				disconnect()
				delay()
				continue
			}
			sbscr, err = r.subscribe(client, headers)
			if err != nil {
				disconnect()
				delay()
				continue
			}
		}

		for curBlock.Cmp(maxBlock) <= 0 {
			was, err = r.readEvents(curBlock, client, was)
			if err != nil {
				disconnect()
				delay()
				break
			}
			curBlock.Add(curBlock, big.NewInt(1))
		}

		r.Log.Warn("wait for next block")
		select {
		case b := <-headers:
			if maxBlock.Cmp(b.Number) < 0 {
				maxBlock.Set(b.Number)
			}
		case <-r.done:
			return
		}
	}
}

func (s *Reader) readEvents(n *big.Int, client *ftmclient.Client, was0 map[hash.Event]struct{}) (was1 map[hash.Event]struct{}, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	blk, err := client.BlockByNumber(ctx, n)
	cancel()
	if err != nil {
		s.Log.Error("detected block", "n", n, "err", err)
		return
	}
	atropos := hash.Event(blk.Hash())
	s.Log.Info("detected block", "n", n, "atropos", atropos)

	was1 = make(map[hash.Event]struct{})
	queue := make([]hash.Event, 0, 100)
	queue = append(queue, atropos)

	for len(queue) > 0 {
		e := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		var event inter.EventI
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		event, err = client.GetEvent(ctx, e)
		cancel()
		if err != nil {
			return
		}
		s.Log.Info("got event", "block", n, "id", event.ID())
		select {
		case s.output <- &asyncTask{
			block: idx.Block(n.Uint64()),
			event: event,
			role:  "TODO",
		}:
			was1[event.ID()] = struct{}{}
		case <-s.done:
			err = fmt.Errorf("interrupted")
			return
		}

		for _, p := range event.Parents() {
			if _, was := was0[p]; was {
				continue
			}
			if _, was := was1[p]; was {
				continue
			}

			s.Log.Debug("detected event", "id", event.ID(), "parent", p, "block", n)
			queue = append(queue, p)
		}
	}

	return
}

func (s *Reader) connect() (*ftmclient.Client, error) {
	client, err := ftmclient.Dial(s.url)
	if err != nil {
		s.Log.Error("connect to", "url", s.url, "err", err)
		return nil, err
	}
	s.Log.Info("connect to", "url", s.url)
	return client, nil
}

func (s *Reader) subscribe(client *ftmclient.Client, headers chan *types.Header) (sbscr ethereum.Subscription, err error) {
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
