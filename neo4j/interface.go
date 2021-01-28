package neo4j

import (
	"github.com/Fantom-foundation/go-lachesis/inter"
)

type ToStore interface {
	Payload() *inter.EventHeaderData
	Done()
}
