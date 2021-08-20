package neo4j

import (
	"testing"

	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/stretchr/testify/require"

	"github.com/Fantom-foundation/lachesis-dag-tool/dagreader/internal"
)

func TestNeo4jMarshaling(t *testing.T) {
	require := require.New(t)

	event := &inter.MutableEventPayload{}
	event.SetCreator(3)

	info0 := &internal.EventInfo{
		Block: 10,
		Role:  "root",
		Event: &event.Build().Event,
	}
	ff := marshal(info0)

	info1 := &internal.EventInfo{}
	unmarshal(ff, info1)

	require.Equal(info0, info1)
}

func TestEventIdParsing(t *testing.T) {
	require := require.New(t)
	for i, e0 := range []hash.Event{
		hash.Event{},
		hash.HexToEventHash("0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"),
		hash.FakeEvent(),
		hash.FakeEvent(),
		hash.FakeEvent(),
	} {
		s := eventId2str(e0)
		e1 := str2eventId(s)

		require.Equal(e0, e1, i, s)
	}
}
