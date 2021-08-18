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

	events struct {
		info      map[hash.Event]*internal.EventInfo
		processed map[idx.Epoch]map[hash.Event]dag.Event
	}

	ordering *dagordering.EventsBuffer

	output chan *internal.EventInfo
	busy   sync.WaitGroup
	sync.RWMutex

	logger.Instance
}

func NewEventsBuffer(db internal.Db, done <-chan struct{}) *EventsBuffer {
	const count = 3000

	s := &EventsBuffer{
		db:        db,
		currEpoch: db.GetEpoch(),
		output:    make(chan *internal.EventInfo, 10),
	}

	s.events.processed = make(map[idx.Epoch]map[hash.Event]dag.Event, 3)
	s.events.processed[s.currEpoch] = make(map[hash.Event]dag.Event, count)
	s.events.info = make(map[hash.Event]*internal.EventInfo, count)

	go db.Load(s.output)

	s.ordering = dagordering.New(dag.Metric{
		Num:  count,
		Size: cachescale.Identity.U64(10 * opt.MiB),
	}, dagordering.Callback{
		Process: func(e dag.Event) error {
			s.RLock()
			defer s.RUnlock()

			id := e.ID()
			epoch := id.Epoch()
			info := s.events.info[id]
			if info == nil {
				panic("event info not found")
			}
			if _, exists := s.events.processed[epoch]; !exists {
				s.events.processed[epoch] = make(map[hash.Event]dag.Event, count)
				delete(s.events.processed, epoch-2)
				s.currEpoch = epoch
			}

			select {
			case s.output <- info:
				s.events.processed[epoch][id] = e
				delete(s.events.info, id)
			case <-done:
				return fmt.Errorf("Interrupted")
			}

			return nil
		},

		Exists: func(e hash.Event) bool {
			s.RLock()
			defer s.RUnlock()

			ee, exists := s.events.processed[e.Epoch()]
			if !exists {
				return false
			}
			_, has := ee[e]
			return has
		},

		Get: func(e hash.Event) dag.Event {
			s.RLock()
			defer s.RUnlock()

			ee, exists := s.events.processed[e.Epoch()]
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

func (s *EventsBuffer) Push(e *internal.EventInfo) {
	s.Lock()
	defer s.Unlock()

	s.events.info[e.Event.ID()] = e
	s.ordering.PushEvent(e.Event, "")
}

func (s *EventsBuffer) Close() {
	s.Lock()
	defer s.Unlock()

	close(s.output)
	s.ordering.Clear()
}

func (s *EventsBuffer) GetEpoch() idx.Epoch {
	s.RLock()
	defer s.RUnlock()

	return s.currEpoch
}
