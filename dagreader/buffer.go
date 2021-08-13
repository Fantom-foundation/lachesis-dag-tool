package main

import (
	"sync"

	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/logger"
	"github.com/Fantom-foundation/lachesis-base/gossip/dagordering"
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/dag"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/Fantom-foundation/lachesis-base/utils/cachescale"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type EventsBuffer struct {
	db        Db
	currEpoch idx.Epoch
	inmem     map[idx.Epoch]map[hash.Event]dag.Event

	buffer *dagordering.EventsBuffer

	busy sync.WaitGroup
	sync.RWMutex

	logger.Instance
}

func NewEventsBuffer(db Db) *EventsBuffer {
	s := &EventsBuffer{
		db:        db,
		currEpoch: db.GetEpoch(),
		inmem:     make(map[idx.Epoch]map[hash.Event]dag.Event, 3),
	}

	s.inmem[s.currEpoch] = make(map[hash.Event]dag.Event, 5000)

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

			err := s.db.Load(e)
			if err != nil {
				return err
			}

			s.inmem[epoch][id] = e
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

func (s *EventsBuffer) Close() {
	s.Lock()
	defer s.Unlock()

	s.buffer.Clear()
}

func (s *EventsBuffer) GetEpoch() idx.Epoch {
	s.RLock()
	defer s.RUnlock()

	return s.currEpoch
}

type asyncTask struct {
	event  inter.EventI
	onDone func()
}

func (t *asyncTask) Payload() inter.EventI {
	return t.event
}

func (t *asyncTask) Done() {
	if t.onDone != nil {
		t.onDone()
	}
}
