package neo4j

import (
	"encoding/json"
	"strings"

	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/neo4j/neo4j-go-driver/neo4j"

	"github.com/Fantom-foundation/lachesis-dag-tool/dagreader/internal"
)

type fields map[string]interface{}

func readFields(r neo4j.Record) fields {
	ff := make(fields)
	vals := r.Values()
	for i, key := range r.Keys() {
		ff[key] = vals[i]
	}
	return ff
}

func (ff fields) String() string {

	buf := &strings.Builder{}

	buf.WriteString("{")
	first := true
	for k, v := range ff {
		if !first {
			buf.WriteString(",")
		} else {
			first = false
		}
		buf.WriteString(k)
		buf.WriteString(":")
		buf.WriteString(valToString(v))
	}
	buf.WriteString("}")
	return buf.String()
}

func valToString(v interface{}) string {
	bb, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}

	return string(bb)
}

func marshal(x interface{}) fields {
	switch v := x.(type) {
	case *internal.EventInfo:
		return fields{
			"block":   uint64(v.Block),
			"role":    v.Role,
			"id":      eventID(v.Event.ID()),
			"creator": int64(v.Event.Creator()),
		}
	default:
		panic("unsupported type")
	}
}

func unmarshal(ff fields, x interface{}) {
	switch v := x.(type) {
	case *internal.EventInfo:
		v.Block = idx.Block(ff["block"].(uint64))
		v.Role = ff["role"].(string)

		event := &inter.MutableEventPayload{}
		id := eventHash(ff["id"].(string))
		event.SetEpoch(id.Epoch())
		event.SetLamport(id.Lamport())
		event.SetID(eventHashID(id))

		event.SetCreator(idx.ValidatorID(ff["creator"].(int64)))

		v.Event = &event.Build().Event
		return
	default:
		panic("unsupported type")
	}
}

// eventID is a FullID() without aliases.
func eventID(e hash.Event) string {
	return e.Hex()
}

func eventHash(id string) hash.Event {
	return hash.HexToEventHash(id)
}

func eventHashID(e hash.Event) (r [24]byte) {
	copy(r[:], e.Bytes()[8:])
	return
}
