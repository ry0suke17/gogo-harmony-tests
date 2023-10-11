package gogo_harmony_tests

import (
	"bytes"
	"encoding/json"
	"log"
	"strings"
	"testing"
	"time"

	gogojsonpb "github.com/gogo/protobuf/jsonpb"
	"github.com/golang/protobuf/jsonpb"
	goproto "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"

	gotest "github.com/ry0suke17/gogo-harmony-tests/proto/go"
	gogotest "github.com/ry0suke17/gogo-harmony-tests/proto/gogofaster"
)

func Test_GoGoFasterToGo(t *testing.T) {
	ggTest := &gogotest.Test{
		At:   time.Date(2022, 1, 11, 3, 4, 5, 6, time.UTC),
		Type: gogotest.Test_TYPE_FIRST,
		Inner: &gogotest.Inner{
			At:   time.Date(2022, 2, 11, 3, 4, 5, 6, time.UTC),
			Type: gogotest.Inner_TYPE_FIRST,
		},
	}
	ggTestMarshaled, err := ggTest.Marshal()
	require.NoError(t, err)

	gTest := &gotest.Test{}
	err = goproto.Unmarshal(ggTestMarshaled, gTest)
	require.NoError(t, err)
	ts, err := ptypes.TimestampProto(ggTest.At)
	require.NoError(t, err)
	innerTs, err := ptypes.TimestampProto(ggTest.Inner.At)
	require.NoError(t, err)

	want := &gotest.Test{
		At:   ts,
		Type: gotest.Test_TYPE_FIRST,
		Inner: &gotest.Inner{
			At:   innerTs,
			Type: gotest.Inner_TYPE_FIRST,
		},
	}

	// success
	assertGoTestEqual(t, want, gTest)
	// failed to not match atomicMessageInfo
	//assert.Equal(t, want, gTest)
}

func assertGoTestEqual(
	t *testing.T,
	expected *gotest.Test,
	actual *gotest.Test,
) {
	assert.Equal(t, expected.At, actual.At)
	assert.Equal(t, expected.Type, actual.Type)
	assert.Equal(t, expected.Inner.At, actual.Inner.At)
	assert.Equal(t, expected.Inner.Type, actual.Inner.Type)
}

func Test_GoToGoGoFaster(t *testing.T) {
	at, err := ptypes.TimestampProto(time.Date(2022, 1, 11, 3, 4, 5, 6, time.UTC))
	require.NoError(t, err)
	innerAt, err := ptypes.TimestampProto(time.Date(2022, 2, 11, 3, 4, 5, 6, time.UTC))
	require.NoError(t, err)

	gTest := &gotest.Test{
		At:   at,
		Type: gotest.Test_TYPE_FIRST,
		Inner: &gotest.Inner{
			At:   innerAt,
			Type: gotest.Inner_TYPE_FIRST,
		},
	}
	gTestMarshaled, err := goproto.Marshal(gTest)
	require.NoError(t, err)

	ggTest := &gogotest.Test{}
	err = ggTest.Unmarshal(gTestMarshaled)
	require.NoError(t, err)
	ts, err := ptypes.Timestamp(gTest.At)
	require.NoError(t, err)
	innerTs, err := ptypes.Timestamp(gTest.Inner.At)
	require.NoError(t, err)

	want := &gogotest.Test{
		At:   ts,
		Type: gogotest.Test_TYPE_FIRST,
		Inner: &gogotest.Inner{
			At:   innerTs,
			Type: gogotest.Inner_TYPE_FIRST,
		},
	}

	// success
	assertGoGoTestEqual(t, want, ggTest)
	// failed to not match atomicMessageInfo
	//assert.Equal(t, want, gTest)
}

func assertGoGoTestEqual(
	t *testing.T,
	expected *gogotest.Test,
	actual *gogotest.Test,
) {
	assert.Equal(t, expected.At, actual.At)
	assert.Equal(t, expected.Type, actual.Type)
	assert.Equal(t, expected.Inner.At, actual.Inner.At)
	assert.Equal(t, expected.Inner.Type, actual.Inner.Type)
}

var (
	marshallerG = jsonpb.Marshaler{
		EmitDefaults: true,  // Render fields with zero values
		OrigName:     false, // Using camelCase for JSON
		EnumsAsInts:  true,  // Whether to render enum values as integers, as opposed to string values.
	}
	unmarshalerG = jsonpb.Unmarshaler{}

	marshallerGG = gogojsonpb.Marshaler{
		EmitDefaults: true,  // Render fields with zero values
		OrigName:     false, // Using camelCase for JSON
		EnumsAsInts:  true,  // Whether to render enum values as integers, as opposed to string values.
	}
	unmarshalerGG = gogojsonpb.Unmarshaler{}
)

func Test_JSONPB_GoToGo(t *testing.T) {
	marshalG := &gotest.Test{
		At: nil,
	}
	j, err := marshallerG.MarshalToString(marshalG)
	assert.NoError(t, err)
	log.Println(j)

	unmarshalG := &gotest.Test{}
	err = unmarshalerG.Unmarshal(strings.NewReader(j), unmarshalG)
	assert.NoError(t, err)
	log.Println(unmarshalG.At)
	assert.Equal(t, marshalG.At, unmarshalG.At)
}

func Test_JSONPB_GoGoToGoGo(t *testing.T) {
	marshalGG := &gogotest.Test{
		At: time.Now().UTC(),
	}
	j, err := marshallerGG.MarshalToString(marshalGG)
	assert.NoError(t, err)
	log.Println(j)

	unmarshalGG := &gogotest.Test{}
	err = unmarshalerGG.Unmarshal(strings.NewReader(j), unmarshalGG)
	assert.NoError(t, err)
	log.Println(unmarshalGG.At)
	assert.Equal(t, marshalGG.At, unmarshalGG.At)
}

func Test_JSONPB_GoGoToGo(t *testing.T) {
	marshalGG := &gogotest.Test{
		At: time.Time{},
	}
	j, err := marshallerGG.MarshalToString(marshalGG)
	assert.NoError(t, err)
	log.Println(j)

	unmarshalG := &gotest.Test{}
	err = unmarshalerG.Unmarshal(strings.NewReader(j), unmarshalG)
	assert.NoError(t, err)
	log.Println(unmarshalG.At.AsTime())
	assert.Equal(t, marshalGG.At, unmarshalG.At.AsTime())
}

func Test_JSONPB_GoToGoGo(t *testing.T) {
	marshalG := &gotest.Test{
		At: nil,
	}

	j, err := marshallerG.MarshalToString(marshalG)
	assert.NoError(t, err)
	log.Println(j)

	unmarshalGG := &gogotest.Test{}
	err = unmarshalerGG.Unmarshal(strings.NewReader(j), unmarshalGG)
	assert.Error(t, err) // bad Timestamp: parsing time "" as "2006-01-02T15:04:05.999999999Z07:00": cannot parse "" as "2006"
}

func Test_JSON_GoToGo(t *testing.T) {
	marshalG := &gotest.Test{
		At: timestamppb.Now(),
	}
	j, err := json.Marshal(marshalG)
	assert.NoError(t, err)
	log.Println(string(j))

	unmarshalG := &gotest.Test{}
	err = json.Unmarshal(j, unmarshalG)
	assert.NoError(t, err)
	log.Println(unmarshalG.At)
	assert.Equal(t, marshalG.At.AsTime(), unmarshalG.At.AsTime())
}

func Test_JSON_GoGoToGoGo(t *testing.T) {
	marshalGG := &gogotest.Test{
		At: time.Now().UTC(),
	}
	j, err := json.Marshal(marshalGG)
	assert.NoError(t, err)
	log.Println(string(j))

	unmarshalGG := &gogotest.Test{}
	err = json.Unmarshal(j, unmarshalGG)
	assert.NoError(t, err)
	log.Println(unmarshalGG.At)
	assert.Equal(t, marshalGG.At, unmarshalGG.At)
}

func Test_JSON_GoGoToGo(t *testing.T) {
	marshalGG := &gogotest.Test{
		At: time.Now().UTC(),
	}
	j, err := json.Marshal(marshalGG)
	assert.NoError(t, err)
	log.Println(string(j))

	unmarshalG := &gotest.Test{}
	err = json.Unmarshal(j, unmarshalG)
	assert.Error(t, err) // json: cannot unmarshal string into Go struct field Test.at of type timestamppb.Timestamp
}

func Test_JSON_GoToGoGo(t *testing.T) {
	marshalG := &gotest.Test{
		At: timestamppb.Now(),
	}

	j, err := json.Marshal(marshalG)
	assert.NoError(t, err)
	log.Println(string(j))

	unmarshalGG := &gogotest.Test{}
	err = json.Unmarshal(j, unmarshalGG)
	assert.Error(t, err) // parsing time "{\"seconds\":1696413080,\"nanos\":983043000}" as "\"2006-01-02T15:04:05Z07:00\"": cannot parse "{\"seconds\":1696413080,\"nanos\":983043000}" as "\""
}

func Test_JSONPBToJSONPB(t *testing.T) {
	at := timestamppb.Now()
	marshalG := &gotest.Test{
		At:        nil,
		CreatedAt: at,
	}

	j, err := marshallerG.MarshalToString(marshalG)
	assert.NoError(t, err)
	log.Printf("json: %s", j)

	unmarshalG := &gotest.Test{}
	err = unmarshalerG.Unmarshal(strings.NewReader(j), unmarshalG)
	assert.NoError(t, err)
	log.Printf("at: %s", unmarshalG.At)
	log.Printf("at as time: %s", unmarshalG.At.AsTime())
	log.Printf("cratedAt: %s", unmarshalG.CreatedAt)
}

func Test_JSONToJSONPB(t *testing.T) {
	at := time.Now()
	marshalGG := &gogotest.Test{
		At:        time.Time{},
		CreatedAt: at,
	}

	j, err := json.Marshal(marshalGG)
	assert.NoError(t, err)
	log.Printf("json: %s", j)

	unmarshalG := &gotest.Test{}
	err = unmarshalerG.Unmarshal(bytes.NewReader(j), unmarshalG)
	assert.NoError(t, err)
	log.Printf("at: %s", unmarshalG.At)
	log.Printf("at as time: %s", unmarshalG.At.AsTime())
	log.Printf("cratedAt: %s", unmarshalG.CreatedAt)
}

func Test_JSONPBToJSON(t *testing.T) {
	marshalG := &gotest.Test{
		At:        nil,
		CreatedAt: timestamppb.Now(),
	}

	j, err := marshallerG.MarshalToString(marshalG)
	assert.NoError(t, err)
	log.Printf("json: %s", j)

	unmarshalGG := &gogotest.Test{}
	json.Unmarshal([]byte(j), unmarshalGG)
	assert.NoError(t, err)
	log.Printf("at: %s", unmarshalGG.At)
	log.Printf("cratedAt: %s", unmarshalGG.CreatedAt)
}
