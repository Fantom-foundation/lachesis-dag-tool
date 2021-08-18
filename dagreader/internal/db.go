package internal

import (
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/dag"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
)

type Storage interface {
	GetLastBlock() idx.Block
	HasEvent(hash.Event) bool
	GetEvent(hash.Event) dag.Event
}

type Db interface {
	Storage
	Load(events <-chan *EventInfo)
}

type EventInfo struct {
	Block   idx.Block
	Event   dag.Event
	Role    string
	Dispose func()
}

func (e *EventInfo) Done() {
	if e.Dispose != nil {
		e.Dispose()
	}
}
