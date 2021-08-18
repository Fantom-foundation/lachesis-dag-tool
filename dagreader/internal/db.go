package internal

import (
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/dag"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
)

type Db interface {
	GetEpoch() idx.Epoch
	HasEvent(e hash.Event) bool
	GetEvent(e hash.Event) dag.Event
	Load(events <-chan *EventInfo)
}

type EventInfo struct {
	Block  idx.Block
	Event  dag.Event
	Role   string
	OnDone func()
}

func (e *EventInfo) Done() {
	if e.OnDone != nil {
		e.OnDone()
	}
}
