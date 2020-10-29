package neo4j

import (
	"testing"

	"github.com/Fantom-foundation/go-lachesis/inter"
	"github.com/stretchr/testify/require"
)

func TestMarshaling(t *testing.T) {
	_ = require.New(t)

	header := &inter.EventHeaderData{}
	header.Creator = 9

	fields := Marshal(header)
	t.Log(fields.String())
}
