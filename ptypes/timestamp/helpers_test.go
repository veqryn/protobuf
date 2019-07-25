package timestamp

import (
	"testing"

	"github.com/golang/protobuf/jsonpb"
)

func TestJsonPBMarshal(t *testing.T) {
	_ = jsonpb.Marshaler{OrigName: true, EmitDefaults: true}
}
