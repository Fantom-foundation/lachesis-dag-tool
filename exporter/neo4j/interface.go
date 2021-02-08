package neo4j

import (
	"github.com/Fantom-foundation/go-lachesis/hash"
	"github.com/Fantom-foundation/go-lachesis/inter"
)

type ToStore interface {
	Payload() *inter.EventHeaderData
	Done()
}

// storedEvent stores origin event Hash,
// because EventHeaderData.Hash() differs as not all the fields are stored.
type storedEvent struct {
	OriginHash hash.Event
	*inter.EventHeaderData
}
