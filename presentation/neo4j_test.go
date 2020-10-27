package presentation

import (
	"testing"

	"github.com/Fantom-foundation/go-lachesis/inter"
	"github.com/stretchr/testify/require"
)

func TestToJSON(t *testing.T) {
	_ = require.New(t)

	header := &inter.EventHeaderData{}
	header.Creator = 9

	fields := Neo4jMarshal(header)
	t.Log(fields.String())
}
