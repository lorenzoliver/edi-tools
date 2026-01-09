package parser

import (
	"fmt"
	"reflect"
)

var registry = make(map[string]reflect.Type)

func init() {
}

func GetSpec(messageType string) (any, error) {
	spec, ok := registry[messageType]
	if !ok {
		return nil, fmt.Errorf("no spec found for message type %s", messageType)
	}
	return reflect.New(spec).Interface(), nil
}

func Register(ident string, t reflect.Type) {
	registry[ident] = t
}
