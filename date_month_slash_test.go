package extime

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestDateMonthSlashJSON(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var got DateMonthSlash
		var want *DateMonthSlash

		b, err := json.Marshal(want)
		require.NoError(t, err)
		err = got.UnmarshalJSON(b)
		require.NoError(t, err)
		require.True(t, got.StdTime().IsZero())
	})
	t.Run("not nil", func(t *testing.T) {
		var got DateMonthSlash

		want := DateMonthSlash(time.Date(2021, 2, 11, 0, 0, 0, 0, time.Local))
		b, err := json.Marshal(want)
		require.NoError(t, err)
		err = json.Unmarshal(b, &got)
		require.NoError(t, err)
		require.Equal(t, want.StdTime().Year(), got.StdTime().Year())
		require.Equal(t, want.StdTime().Month(), got.StdTime().Month())

		t.Log(got.String())
	})
}

func TestDateMonthSlashTEXT(t *testing.T) {
	t.Run("not nil", func(t *testing.T) {
		var got DateMonthSlash

		want := DateMonthSlash(time.Date(2021, 2, 11, 0, 0, 0, 0, time.Local))
		b, err := want.MarshalText()
		require.NoError(t, err)

		err = got.UnmarshalText(b)
		require.NoError(t, err)
		require.Equal(t, want.StdTime().Year(), got.StdTime().Year())
		require.Equal(t, want.StdTime().Month(), got.StdTime().Month())

		t.Log(got.String())
	})
}
