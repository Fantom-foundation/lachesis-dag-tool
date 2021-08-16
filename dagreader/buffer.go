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
	inmem     map[idx.Epoch]map[hash.Event]dag.Event

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
		inmem:     make(map[idx.Epoch]map[hash.Event]dag.Event, 3),
		output:    make(chan internal.ToStore, 10),
	}

	s.inmem[s.currEpoch] = make(map[hash.Event]dag.Event, 5000)

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
			ee, exists := s.inmem[epoch]
			if !exists {
				ee = make(map[hash.Event]dag.Event, 5000)
				s.inmem[epoch] = ee
				s.currEpoch = epoch
				delete(s.inmem, epoch-2)
			}

			select {
			case s.output <- &asyncTask{
				event: e,
				role:  "TODO",
			}:
				s.inmem[epoch][id] = e
			case <-done:
				return fmt.Errorf("Interrupted")
			}

			return nil
		},

		Exists: func(e hash.Event) bool {
			s.RLock()
			defer s.RUnlock()

			ee, exists := s.inmem[e.Epoch()]
			if !exists {
				return false
			}
			_, has := ee[e]
			return has
		},

		Get: func(e hash.Event) dag.Event {
			s.RLock()
			defer s.RUnlock()

			ee, exists := s.inmem[e.Epoch()]
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

func (s *EventsBuffer) Push(e dag.Event) {
	s.buffer.PushEvent(e, "")
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
	event  dag.Event
	role   string
	onDone func()
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
