package main

import (
	"sync"

	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/go-opera/logger"
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"

	"github.com/Fantom-foundation/lachesis-dag-tool/dagreader/neo4j"
)

type task struct {
	event  inter.EventI
	onDone func()
}

func (t *task) Payload() inter.EventI {
	return t.event
}

func (t *task) Done() {
	if t.onDone != nil {
		t.onDone()
	}
}

type Db interface {
	GetEpoch() idx.Epoch
	HasEvent(e hash.Event) bool
	GetEvent(e hash.Event) inter.EventI
	Load(<-chan neo4j.ToStore)
}

type store struct {
	Db
	out    chan neo4j.ToStore
	synced bool
	wg     sync.WaitGroup

	logger.Instance
}

func newStore(db Db, synced bool) *store {
	s := &store{
		Db:     db,
		out:    make(chan neo4j.ToStore, 10),
		synced: synced,

		Instance: logger.MakeInstance(),
	}

	s.SetName("store")

	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.Db.Load(s.out)
	}()

	return s
}

func (s *store) Close() {
	close(s.out)
}

func (s *store) WaitForAll() {
	s.wg.Wait()
}

func (s *store) Save(event inter.EventI) {
	var wg sync.WaitGroup

	t := &task{event: event}
	if s.synced {
		wg.Add(1)
		t.onDone = wg.Done
	}

	s.Log.Info("got event", "id", event.ID())
	s.out <- neo4j.ToStore(t)

	if s.synced {
		wg.Wait()
	}
}
