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
	in  chan *inter.EventHeaderData
	out chan neo4j.ToStore
	wg  sync.WaitGroup
}

func newStore(db *neo4j.Store, synced bool) *store {
	s := &store{
		Store: db,
		in:    make(chan *inter.EventHeaderData, 10),
		out:   make(chan neo4j.ToStore, 10),
	}

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.Store.Load(s.out)
	}()

	if synced {
		go s.syncLoop()
	} else {
		go s.asyncLoop()
	}

	return s
}

func (s *store) asyncLoop() {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		defer close(s.out)

		for e := range s.in {
			t := &task{event: e}
			s.out <- neo4j.ToStore(t)
		}
	}()
}

func (s *store) syncLoop() {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		defer close(s.out)

		var wg sync.WaitGroup
		for e := range s.in {
			wg.Add(1)
			t := &task{event: e, onDone: wg.Done}
			s.out <- neo4j.ToStore(t)
			wg.Wait()
		}
	}()
}

func (s *store) Close() {
	close(s.in)
}

func (s *store) WaitForAll() {
	s.wg.Wait()
}

func (s *store) Save(event *inter.EventHeaderData) {
	s.in <- event
}
