package neo4j

import (
	"fmt"
	"strings"

	"github.com/Fantom-foundation/go-lachesis/inter"
)

type Fields map[string]string

func (ff *Fields) String() string {
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

func Marshal(x interface{}) Fields {
	switch e := x.(type) {
	case *inter.EventHeaderData:
		return Fields{
			"id":      fmt.Sprintf("'%s'", e.Hash().FullID()),
			"creator": fmt.Sprintf("%d", e.Creator),
		}
	default:
		panic("unsupported type")
	}
}
