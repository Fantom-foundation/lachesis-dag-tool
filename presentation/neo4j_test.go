package presentation

import (
	"testing"

	"github.com/Fantom-foundation/go-lachesis/inter"
	"github.com/stretchr/testify/require"
)

func TestToJSON(t *testing.T) {
	require := require.New(t)

	header := &inter.EventHeaderData{}
	encoder := NewNeo4jEncoding(header, "GasPowerLeft", "Parents")

	header.Creator = 9
	data, err := encoder.Marshal(header)
	require.NoError(err)
	t.Log(string(data))
}
