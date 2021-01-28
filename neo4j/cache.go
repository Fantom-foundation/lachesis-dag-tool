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
		event := task.Payload()
		e := event.Hash()

		s.Lock()
		ee, exists := s.inmem[e.Epoch()]
		if !exists {
			ee = make(map[hash.Event]*inter.EventHeaderData, 5000)
			s.inmem[e.Epoch()] = ee
			delete(s.inmem, e.Epoch()-2)
		}
		ee[e] = event
		s.Unlock()
		task.Done()

		toDisk <- &asyncTask{event}
	}
}

func (s *CachedDb) SetEpoch(num idx.Epoch) {
	s.Lock()
	defer s.Unlock()

	s.currEpoch = num
	s.db.SetEpoch(num) // TODO: set later, when flush events to disk
}

func (s *CachedDb) GetEpoch() idx.Epoch {
	s.RLock()
	defer s.RUnlock()

	return s.currEpoch
}

type asyncTask struct {
	event *inter.EventHeaderData
}

func (t *asyncTask) Payload() *inter.EventHeaderData {
	return t.event
}

func (t *asyncTask) Done() {
}
