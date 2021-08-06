package neo4j

import (
	"testing"

	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/stretchr/testify/require"
)

func TestMarshaling(t *testing.T) {
	_ = require.New(t)

	header := new(inter.Event)
	// header.Creator = 9

	ff := marshal(header)
	t.Log(ff)

	unmarshal(ff, header)
}
