package gogo_harmony_tests

import (
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

func Test_JSONPB_GoToGo(t *testing.T) {
	marshaller := jsonpb.Marshaler{
		EmitDefaults: true,  // Render fields with zero values
		OrigName:     false, // Using camelCase for JSON
		EnumsAsInts:  true,  // Whether to render enum values as integers, as opposed to string values.
	}
	unmarshaler := jsonpb.Unmarshaler{}

	g := &gotest.Test{
		At: nil,
	}
	j, err := marshaller.MarshalToString(g)
	assert.NoError(t, err)
	log.Println(j)

	gTest := gotest.Test{}
	err = unmarshaler.Unmarshal(strings.NewReader(j), &gTest)
	assert.NoError(t, err)
	log.Println(gTest.At)
}

func Test_JSONPB_GoGoToGoGo(t *testing.T) {
	marshaller := gogojsonpb.Marshaler{
		EmitDefaults: true,  // Render fields with zero values
		OrigName:     false, // Using camelCase for JSON
		EnumsAsInts:  true,  // Whether to render enum values as integers, as opposed to string values.
	}
	unmarshaler := gogojsonpb.Unmarshaler{}

	g := &gogotest.Test{
		At: time.Time{},
	}
	j, err := marshaller.MarshalToString(g)
	assert.NoError(t, err)
	log.Println(j)

	gTest := gogotest.Test{}
	err = unmarshaler.Unmarshal(strings.NewReader(j), &gTest)
	assert.NoError(t, err)
	log.Println(gTest.At)
}

func Test_JSONPB_GoGoToGo(t *testing.T) {
	marshaller := gogojsonpb.Marshaler{
		EmitDefaults: true,  // Render fields with zero values
		OrigName:     false, // Using camelCase for JSON
		EnumsAsInts:  true,  // Whether to render enum values as integers, as opposed to string values.
	}
	unmarshaler := jsonpb.Unmarshaler{}

	g := &gogotest.Test{
		At: time.Time{},
	}
	j, err := marshaller.MarshalToString(g)
	assert.NoError(t, err)
	log.Println(j)

	gTest := gotest.Test{}
	err = unmarshaler.Unmarshal(strings.NewReader(j), &gTest)
	assert.NoError(t, err)
	log.Println(gTest.At.AsTime())
}

func Test_JSONPB_GoToGoGo(t *testing.T) {
	marshaller := jsonpb.Marshaler{
		EmitDefaults: true,  // Render fields with zero values
		OrigName:     false, // Using camelCase for JSON
		EnumsAsInts:  true,  // Whether to render enum values as integers, as opposed to string values.
	}
	unmarshaler := gogojsonpb.Unmarshaler{}

	g := &gotest.Test{
		At: nil,
	}

	j, err := marshaller.MarshalToString(g)
	assert.NoError(t, err)
	log.Println(j)
	ggTest2 := gogotest.Test{}
	err = unmarshaler.Unmarshal(strings.NewReader(j), &ggTest2)
	assert.NoError(t, err)
	log.Println(ggTest2.At)
}
