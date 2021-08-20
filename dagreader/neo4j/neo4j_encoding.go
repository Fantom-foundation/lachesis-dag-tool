package neo4j

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/Fantom-foundation/lachesis-base/inter/idx"
	"github.com/ethereum/go-ethereum/common"
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
			"block":   int64(v.Block),
			"role":    v.Role,
			"id":      eventId2str(v.Event.ID()),
			"creator": int64(v.Event.Creator()),
			"parents": v.Event.Parents(),
		}
	default:
		panic("unsupported type")
	}
}

func unmarshal(ff fields, x interface{}) {
	switch v := x.(type) {
	case *internal.EventInfo:
		v.Block = idx.Block(ff["block"].(int64))
		v.Role = ff["role"].(string)

		event := &inter.MutableEventPayload{}
		id := str2eventId(ff["id"].(string))
		event.SetEpoch(id.Epoch())
		event.SetLamport(id.Lamport())
		event.SetID(eventIdTail(id))

		event.SetCreator(idx.ValidatorID(ff["creator"].(int64)))

		event.SetParents(ff["parents"].(hash.Events))

		v.Event = &event.Build().Event
		return
	default:
		panic("unsupported type")
	}
}

func eventId2str(e hash.Event) string {
	return e.FullID()
}

// TODO: mv to the "github.com/Fantom-foundation/lachesis-base/hash"
func str2eventId(s string) (id hash.Event) {
	parts := strings.SplitN(s, ":", 3)

	n, err := strconv.ParseUint(parts[0], 10, 32)
	if err != nil {
		panic(err)
	}
	copy(id[0:], idx.Epoch(n).Bytes())

	n, err = strconv.ParseUint(parts[1], 10, 32)
	if err != nil {
		panic(err)
	}
	copy(id[4:], idx.Lamport(n).Bytes())

	copy(id[8:], common.Hex2BytesFixed(parts[2], 32-4-4))

	return
}

func eventIdTail(e hash.Event) (r [24]byte) {
	copy(r[:], e.Bytes()[8:])
	return
}
