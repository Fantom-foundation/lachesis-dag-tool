package neo4j

import (
	"github.com/Fantom-foundation/go-lachesis/inter"
)

type EventData struct {
	Event *inter.Event
	Ready func()
}
