package presentation

import (
	"fmt"
	"strings"

	"github.com/Fantom-foundation/go-lachesis/inter"
)

type Neo4jFields map[string]string

func (ff *Neo4jFields) String() string {
	buf := &strings.Builder{}

	buf.WriteString("{")
	first := true
	for k, v := range *ff {
		if !first {
			buf.WriteString(",")
		} else {
			first = false
		}
		buf.WriteString(k)
		buf.WriteString(":")
		buf.WriteString(v)
	}
	buf.WriteString("}")

	return buf.String()
}

func Neo4jMarshal(x interface{}) Neo4jFields {
	switch e := x.(type) {
	case *inter.EventHeaderData:
		return Neo4jFields{
			"hash":    fmt.Sprintf("'%s'", e.Hash().Hex()),
			"creator": fmt.Sprintf("%d", e.Creator),
		}
	default:
		panic("unsupported type")
	}
}
