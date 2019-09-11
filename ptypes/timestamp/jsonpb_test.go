package timestamp

import (
	"testing"

	"github.com/golang/protobuf/jsonpb"
	google_tspb "github.com/golang/protobuf/ptypes/timestamp"
)

func TestJsonPBMarshal(t *testing.T) {
	t.Parallel()

	marshaller := jsonpb.Marshaler{OrigName: true, EmitDefaults: true}

	for _, test := range []struct {
		ts *Timestamp
	}{
		{&Timestamp{Seconds: 0, Nanos: 0}},
		{&Timestamp{Seconds: 1506956400, Nanos: 0}},
		{&Timestamp{Seconds: 1570028400, Nanos: 50000000}},
	} {
		veqryn, err := marshaller.MarshalToString(test.ts)
		if err != nil {
			t.Errorf("MarshalToString(%v) error = %s", test.ts, err)
		}

		google, err := marshaller.MarshalToString(&google_tspb.Timestamp{Seconds: test.ts.Seconds, Nanos: test.ts.Nanos})
		if err != nil {
			t.Errorf("MarshalToString(%v) error = %s", test.ts, err)
		}

		if veqryn != google {
			t.Errorf("MarshalToString(%v) = %s, want %s", test.ts, veqryn, google)
		}
	}
}

func TestJsonPBUnmarshal(t *testing.T) {
	t.Parallel()

	marshaller := jsonpb.Marshaler{}

	for _, test := range []struct {
		ts *google_tspb.Timestamp
	}{
		{&google_tspb.Timestamp{Seconds: 0, Nanos: 0}},
		{&google_tspb.Timestamp{Seconds: 1506956400, Nanos: 0}},
		{&google_tspb.Timestamp{Seconds: 1570028400, Nanos: 50000000}},
	} {
		str, err := marshaller.MarshalToString(test.ts)
		if err != nil {
			t.Errorf("MarshalToString(%v) error = %s", test.ts, err)
		}

		got := &Timestamp{}
		err = jsonpb.UnmarshalString(str, got)
		if err != nil {
			t.Errorf("UnmarshalString(%v) error = %s", str, err)
		}

		if test.ts.Seconds != got.Seconds || test.ts.Nanos != got.Nanos {
			t.Errorf("UnmarshalString(%v) = %s, want %s", str, got, test.ts)
		}
	}
}
