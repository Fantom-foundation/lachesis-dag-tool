package neo4j

import (
	"encoding/json"
	"strings"

	"github.com/Fantom-foundation/go-opera/inter"
	"github.com/Fantom-foundation/lachesis-base/hash"
	"github.com/neo4j/neo4j-go-driver/neo4j"
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

// TODO: all the fields
func marshal(x interface{}) fields {
	switch e := x.(type) {
	case *inter.Event:
		return fields{
			"id":      e.ID().Hex(),
			"creator": int64(e.Creator()),
		}
	default:
		panic("unsupported type")
	}
}

// TODO: all the fields
func unmarshal(ff fields, x interface{}) {
	switch x.(type) {
	case *inter.Event:
		// e.Creator = idx.StakerID(ff["creator"].(int64))
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
