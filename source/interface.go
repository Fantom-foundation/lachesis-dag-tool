package source

import (
	"github.com/Fantom-foundation/go-lachesis/hash"
	"github.com/Fantom-foundation/go-lachesis/inter"
	"github.com/Fantom-foundation/go-lachesis/inter/idx"
)

type Store interface {
	SetEpoch(idx.Epoch)
	GetEpoch() idx.Epoch
	HasEvent(e hash.Event) bool
	GetEvent(e hash.Event) *inter.EventHeaderData

	Save(*inter.EventHeaderData)
	Close()
}
