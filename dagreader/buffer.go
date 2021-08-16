package main

import (
	"fmt"
	"sync"

	"github.com/Fantom-foundation/go-opera/logger"
	"github.com/Fantom-foundation/lachesis-base/gossip/dagordering"
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/dag"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/utils/cachescale"
	"github.com/syndtr/goleveldb/leveldb/opt"

	"github.com/Fantom-foundation/lachesis-dag-tool/dagreader/internal"
)

type EventsBuffer struct {
	db        internal.Db
	currEpoch idx.Epoch

	cache struct {
		infos  map[hash.Event]internal.ToStore
		events map[idx.Epoch]map[hash.Event]dag.Event
	}

	buffer *dagordering.EventsBuffer

	output chan internal.ToStore
	busy   sync.WaitGroup
	sync.RWMutex

	logger.Instance
}

func NewEventsBuffer(db internal.Db, done <-chan struct{}) *EventsBuffer {
	s := &EventsBuffer{
		db:        db,
		currEpoch: db.GetEpoch(),
		output:    make(chan internal.ToStore, 10),
	}

	s.cache.infos = make(map[hash.Event]internal.ToStore)
	s.cache.events = make(map[idx.Epoch]map[hash.Event]dag.Event, 3)
	s.cache.events[s.currEpoch] = make(map[hash.Event]dag.Event, 5000)

	go db.Load(s.output)

	s.buffer = dagordering.New(dag.Metric{
		Num:  3000,
		Size: cachescale.Identity.U64(10 * opt.MiB),
	}, dagordering.Callback{
		Process: func(e dag.Event) error {
			s.RLock()
			defer s.RUnlock()

			id := e.ID()
			epoch := id.Epoch()
			ee, exists := s.cache.events[epoch]
			if !exists {
				ee = make(map[hash.Event]dag.Event, 5000)
				s.cache.events[epoch] = ee
				s.currEpoch = epoch
				delete(s.cache.events, epoch-2)
			}

			select {
			case s.output <- &asyncTask{
				event: e,
				role:  "TODO",
			}:
				s.cache.events[epoch][id] = e
			case <-done:
				return fmt.Errorf("Interrupted")
			}

			return nil
		},

		Exists: func(e hash.Event) bool {
			s.RLock()
			defer s.RUnlock()

			ee, exists := s.cache.events[e.Epoch()]
			if !exists {
				return false
			}
			_, has := ee[e]
			return has
		},

		Get: func(e hash.Event) dag.Event {
			s.RLock()
			defer s.RUnlock()

			ee, exists := s.cache.events[e.Epoch()]
			if !exists {
				return nil
			}
			event, _ := ee[e]

			return event
		},

		Check: func(e dag.Event, parents dag.Events) error {
			// trust to all
			return nil
		},
	})

	return s
}

func (s *EventsBuffer) Push(e internal.ToStore) {
	s.Lock()
	defer s.Unlock()

	s.buffer.PushEvent(e.Event(), "")
}

func (s *EventsBuffer) Close() {
	s.Lock()
	defer s.Unlock()

	close(s.output)
	s.buffer.Clear()
}

func (s *EventsBuffer) GetEpoch() idx.Epoch {
	s.RLock()
	defer s.RUnlock()

	return s.currEpoch
}

// asyncTask implements ToStore interface
type asyncTask struct {
	block  idx.Block
	event  dag.Event
	role   string
	onDone func()
}

func (t *asyncTask) Block() idx.Block {
	return t.block
}

func (t *asyncTask) Event() dag.Event {
	return t.event
}

func (t *asyncTask) Role() string {
	return t.role
}

func (t *asyncTask) Done() {
	if t.onDone != nil {
		t.onDone()
	}
}
