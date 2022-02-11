package gogo_harmony_tests

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"

	gotest "github.com/ry0suke17/gogo-harmony-tests/proto/go"
	gogotest "github.com/ry0suke17/gogo-harmony-tests/proto/gogofaster"
	"github.com/stretchr/testify/require"
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
	err = proto.Unmarshal(ggTestMarshaled, gTest)
	require.NoError(t, err)
	ts, err := ptypes.TimestampProto(ggTest.At)
	require.NoError(t, err)
	innerTs, err := ptypes.TimestampProto(ggTest.Inner.At)
	require.NoError(t, err)

	log.Printf("gTest: %+v", ggTest)

	want := &gotest.Test{
		At:   ts,
		Type: gotest.Test_TYPE_FIRST,
		Inner: &gotest.Inner{
			At:   innerTs,
			Type: gotest.Inner_TYPE_FIRST,
		},
	}
	assert.Equal(t, want, gTest)
}
