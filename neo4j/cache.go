package neo4j

import (
	"sync"

	"github.com/Fantom-foundation/go-lachesis/hash"
	"github.com/Fantom-foundation/go-lachesis/inter"
	"github.com/Fantom-foundation/go-lachesis/inter/idx"
)

type CachedDb struct {
	db        *Db
	currEpoch idx.Epoch
	inmem     map[idx.Epoch]map[hash.Event]*inter.EventHeaderData

	sync.RWMutex
}

func NewCachedDb(db *Db) *CachedDb {
	s := &CachedDb{
		db:        db,
		currEpoch: db.GetEpoch(),
		inmem:     make(map[idx.Epoch]map[hash.Event]*inter.EventHeaderData, 3),
	}

	return s
}

func (s *CachedDb) Close() error {
	s.Lock()
	defer s.Unlock()

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

func (s *CachedDb) GetEvent(e hash.Event) *inter.EventHeaderData {
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
	toDisk := make(chan ToStore, 100)
	defer close(toDisk)

	go s.db.Load(toDisk)

	for task := range events {
		var updateEpoch func()
		event := task.Payload()
		e := event.Hash()
		epoch := e.Epoch()

		s.Lock()
		ee, exists := s.inmem[epoch]
		if !exists {
			ee = make(map[hash.Event]*inter.EventHeaderData, 5000)
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
	event  *inter.EventHeaderData
	onDone func()
}

func (t *asyncTask) Payload() *inter.EventHeaderData {
	return t.event
}

func (t *asyncTask) Done() {
	if t.onDone != nil {
		t.onDone()
	}
}
