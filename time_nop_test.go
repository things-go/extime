package extime

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTimeNopJSON(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var got TimeNop
		var want *TimeNop

		b, err := json.Marshal(want)
		require.NoError(t, err)
		err = got.UnmarshalJSON(b)
		require.NoError(t, err)
		require.True(t, got.StdTime().IsZero())
	})
	t.Run("not nil", func(t *testing.T) {
		var got TimeNop

		want := TimeNop(time.Now())
		b, err := json.Marshal(want)
		require.NoError(t, err)
		err = got.UnmarshalJSON(b)
		require.NoError(t, err)
		require.False(t, got.StdTime().IsZero())
		require.Equal(t, want.StdTime().Unix(), got.StdTime().Unix())

		t.Log(got.String())
	})
}

func TestTimeNopTEXT(t *testing.T) {
	t.Run("not nil", func(t *testing.T) {
		var got TimeNop

		want := TimeNop(time.Now())
		b, err := want.MarshalText()
		require.NoError(t, err)

		err = got.UnmarshalText(b)
		require.NoError(t, err)
		require.Equal(t, want.StdTime().Unix(), got.StdTime().Unix())

		t.Log(got)
	})
}
