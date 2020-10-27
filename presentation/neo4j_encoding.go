package presentation

import (
	"encoding/json"
	"reflect"
	"strings"
)

type Neo4jEncoding struct {
	fields map[int]string
}

func NewNeo4jEncoding(x interface{}, exclude ...string) *Neo4jEncoding {
	e := &Neo4jEncoding{
		fields: make(map[int]string),
	}

	t := reflect.TypeOf(x).Elem()
	v := reflect.ValueOf(x).Elem()
	for i := 0; i < t.NumField(); i++ {
		if !v.Field(i).CanInterface() {
			continue
		}
		if has(exclude, t.Field(i).Name) {
			continue
		}
		e.fields[i] = t.Field(i).Name
	}

	return e
}

func (e *Neo4jEncoding) Marshal(x interface{}) ([]byte, error) {
	buf := &strings.Builder{}

	buf.WriteString("{")
	v := reflect.ValueOf(x).Elem()
	first := true
	for i, name := range e.fields {
		if !first {
			buf.WriteString(",")
		} else {
			first = false
		}
		buf.WriteString(name)
		buf.WriteString(":")
		vv, err := json.Marshal(v.Field(i).Interface())
		if err != nil {
			return nil, err
		}
		_, err = buf.Write(vv)
		if err != nil {
			return nil, err
		}
	}
	buf.WriteString("}")

	return []byte(buf.String()), nil
}

func has(arr []string, v string) bool {
	for _, exist := range arr {
		if v == exist {
			return true
		}
	}
	return false
}
