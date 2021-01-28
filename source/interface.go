package source

import (
	"github.com/Fantom-foundation/go-lachesis/hash"
	"github.com/Fantom-foundation/go-lachesis/inter"
	"github.com/Fantom-foundation/go-lachesis/inter/idx"
)

type Store interface {
	GetEpoch() idx.Epoch
	HasEvent(e hash.Event) bool
	// GetEvent returns event with incorrect .Hash() because not all the fields are stored.
	GetEvent(e hash.Event) *inter.EventHeaderData

	Save(*inter.EventHeaderData)
	Close()
}
