package extime

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestUnixTimestampJSON(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var got UnixTimestamp
		var want *UnixTimestamp

		b, err := json.Marshal(want)
		require.NoError(t, err)
		err = got.UnmarshalJSON(b)
		require.NoError(t, err)
		require.True(t, got.StdTime().IsZero())
	})
	t.Run("not nil", func(t *testing.T) {
		var got UnixTimestamp

		want := UnixTimestamp(time.Now())
		b, err := json.Marshal(want)
		require.NoError(t, err)
		err = json.Unmarshal(b, &got)
		require.NoError(t, err)
		require.Equal(t, want.StdTime().Unix(), got.StdTime().Unix())

		t.Log(got.String())
	})
}

func TestUnixNanoTimestampJSON(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var got UnixNanoTimestamp
		var want *UnixNanoTimestamp

		b, err := json.Marshal(want)
		require.NoError(t, err)
		err = got.UnmarshalJSON(b)
		require.NoError(t, err)
		require.True(t, got.StdTime().IsZero())
	})

	t.Run("not nil", func(t *testing.T) {
		var got UnixNanoTimestamp

		want := UnixNanoTimestamp(time.Now())
		b, err := json.Marshal(want)
		require.NoError(t, err)
		err = json.Unmarshal(b, &got)
		require.NoError(t, err)
		require.Equal(t, want.StdTime().UnixNano(), got.StdTime().UnixNano())

		t.Log(got.String())
	})
}
