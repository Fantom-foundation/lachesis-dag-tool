package neo4j

import (
	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/lachesis-base/hash"
)

type ToStore interface {
	Payload() inter.EventI
	Done()
}

// storedEvent stores origin event Hash,
// because EventHeaderData.Hash() differs as not all the fields are stored.
type storedEvent struct {
	OriginHash hash.Event
	inter.EventI
}
