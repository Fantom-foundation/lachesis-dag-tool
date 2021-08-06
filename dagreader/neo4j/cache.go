package neo4j

import (
	"sync"
	"time"

	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/log"
)

type CachedDb struct {
	db        *Db
	currEpoch idx.Epoch
	inmem     map[idx.Epoch]map[hash.Event]*inter.Event

	busy sync.WaitGroup
	sync.RWMutex
}

func NewCachedDb(db *Db) *CachedDb {
	s := &CachedDb{
		db:        db,
		currEpoch: db.GetEpoch(),
		inmem:     make(map[idx.Epoch]map[hash.Event]*inter.Event, 3),
	}

	log.Info("Init neo4j cache", "begin", time.Now(), "epoch", s.currEpoch)
	ee := make(map[hash.Event]*inter.Event, 5000)
	for e := range db.getEvents(s.currEpoch) {
		ee[e.OriginHash] = e.Event
		log.Debug("already exists", "event", e.OriginHash, "parents", e.Parents)
	}
	s.inmem[s.currEpoch] = ee
	log.Info("Init neo4j cache", "finish", time.Now(), "events", len(ee))

	return s
}

func (s *CachedDb) Close() error {
	s.Lock()
	defer s.Unlock()

	s.busy.Wait()

	return s.db.Close()
}

func (s *CachedDb) HasEvent(e hash.Event) bool {
	s.RLock()
	defer s.RUnlock()

	ee, exists := s.inmem[e.Epoch()]
	if !exists {
		return false
	}
	_, has := ee[e]
	return has
}

func (s *CachedDb) GetEvent(e hash.Event) *inter.Event {
	s.RLock()
	defer s.RUnlock()

	ee, exists := s.inmem[e.Epoch()]
	if !exists {
		return nil
	}
	event, _ := ee[e]

	return event
}

// Load data from events chain.
func (s *CachedDb) Load(events <-chan ToStore) {
	s.busy.Add(1)
	defer s.busy.Done()

	toDisk := make(chan ToStore, 100)
	defer close(toDisk)

	go s.db.Load(toDisk)

	for task := range events {
		var updateEpoch func()
		event := task.Payload()
		e := event.ID()
		epoch := e.Epoch()

		s.Lock()
		ee, exists := s.inmem[epoch]
		if !exists {
			ee = make(map[hash.Event]*inter.Event, 5000)
			s.inmem[epoch] = ee
			s.currEpoch = epoch
			updateEpoch = func() {
				s.db.setEpoch(epoch)
			}
			delete(s.inmem, epoch-2)
		}
		ee[e] = event
		s.Unlock()
		task.Done()

		later := &asyncTask{event: event, onDone: updateEpoch}
		toDisk <- later
	}
}

func (s *CachedDb) GetEpoch() idx.Epoch {
	s.RLock()
	defer s.RUnlock()

	return s.currEpoch
}

type asyncTask struct {
	event  *inter.Event
	onDone func()
}

func (t *asyncTask) Payload() *inter.Event {
	return t.event
}

func (t *asyncTask) Done() {
	if t.onDone != nil {
		t.onDone()
	}
}
