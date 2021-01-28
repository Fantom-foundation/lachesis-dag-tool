package main

import (
	"sync"

	"github.com/Fantom-foundation/go-lachesis/hash"
	"github.com/Fantom-foundation/go-lachesis/inter"
	"github.com/Fantom-foundation/go-lachesis/inter/idx"

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

type Neo4jDb interface {
	GetEpoch() idx.Epoch
	HasEvent(e hash.Event) bool
	GetEvent(e hash.Event) *inter.EventHeaderData
	Load(<-chan neo4j.ToStore)
}

type store struct {
	Neo4jDb
	out    chan neo4j.ToStore
	synced bool
	wg     sync.WaitGroup
}

func newStore(db Neo4jDb, synced bool) *store {
	s := &store{
		Neo4jDb: db,
		out:     make(chan neo4j.ToStore, 10),
		synced:  synced,
	}

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.Neo4jDb.Load(s.out)
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
