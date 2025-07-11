package gorilla

import (
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
)

func TestNewOption(t *testing.T) {
	marshalOptions := protojson.MarshalOptions{
		Multiline:         true,
		Indent:            "	",
		AllowPartial:      true,
		UseProtoNames:     true,
		UseEnumNumbers:    true,
		EmitUnpopulated:   true,
		EmitDefaultValues: true,
	}
	o := NewOptions(WithMarshalOptions(marshalOptions))
	t.Log(marshalOptions == o.MarshalOptions())
}
