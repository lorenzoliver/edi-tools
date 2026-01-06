package parser

import (
	"fmt"
	"reflect"

	"github.com/lorenzoliver/edi-tools/edifact/directories/d01b"
)

var registry = make(map[string]reflect.Type)

func init() {
	registry["APERAK:D:01B:UN"] = reflect.TypeFor[d01b.APERAK]()
}

func GetSpec(messageType string) (any, error) {
	spec, ok := registry[messageType]
	if !ok {
		return nil, fmt.Errorf("no spec found for message type %s", messageType)
	}
	return reflect.New(spec).Interface(), nil
}
