package neo4j

import (
	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/lachesis-base/hash"
)

// storedEvent stores origin event Hash,
// because EventHeaderData.Hash() differs as not all the fields are stored.
type storedEvent struct {
	OriginHash hash.Event
	inter.EventI
}
