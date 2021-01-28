package main

import (
	"sync"

	"github.com/Fantom-foundation/go-lachesis/inter"

	"github.com/Fantom-foundation/lachesis-dag-tool/neo4j"
)

type task struct {
	event  *inter.EventHeaderData
	onDone func()
}

func (t *task) Payload() *inter.EventHeaderData {
	return t.event
}

func (t *task) Done() {
	if t.onDone != nil {
		t.onDone()
	}
}

type store struct {
	*neo4j.Store
	out    chan neo4j.ToStore
	synced bool
	wg     sync.WaitGroup
}

func newStore(db *neo4j.Store, synced bool) *store {
	s := &store{
		Store:  db,
		out:    make(chan neo4j.ToStore, 10),
		synced: synced,
	}

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.Store.Load(s.out)
	}()

	return s
}

func (s *store) Close() {
	close(s.out)
}

func (s *store) WaitForAll() {
	s.wg.Wait()
}

func (s *store) Save(event *inter.EventHeaderData) {
	var wg sync.WaitGroup

	t := &task{event: event}
	if s.synced {
		wg.Add(1)
		t.onDone = wg.Done
	}

	s.out <- neo4j.ToStore(t)

	if s.synced {
		wg.Wait()
	}
}
