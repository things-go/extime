package extime

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestConver(t *testing.T) {
	require.Equal(t, ToDate(testTime), Date(testTime))
	require.Equal(t, ToDateDot(testTime), DateDot(testTime))
	require.Equal(t, ToDateNop(testTime), DateNop(testTime))
	require.Equal(t, ToDateSlash(testTime), DateSlash(testTime))
	require.Equal(t, ToDateMonthDot(testTime), DateMonthDot(testTime))
	require.Equal(t, ToDateMonthNop(testTime), DateMonthNop(testTime))
	require.Equal(t, ToDateMonthSlash(testTime), DateMonthSlash(testTime))
	require.Equal(t, ToTime(testTime), Time(testTime))
	require.Equal(t, ToTimeNop(testTime), TimeNop(testTime))
	require.Equal(t, ToUnixTimestamp(testTime), UnixTimestamp(testTime))
	require.Equal(t, ToUnixNanoTimestamp(testTime), UnixNanoTimestamp(testTime))
}

func TestDateJSON(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var got Date
		var want *Date

		b, err := json.Marshal(want)
		require.NoError(t, err)
		err = got.UnmarshalJSON(b)
		require.NoError(t, err)
		require.True(t, got.StdTime().IsZero())
	})
	t.Run("not nil", func(t *testing.T) {
		var got Date

		want := Date(time.Date(2021, 2, 11, 0, 0, 0, 0, time.Local))
		b, err := json.Marshal(want)
		require.NoError(t, err)
		err = json.Unmarshal(b, &got)
		require.NoError(t, err)
		require.Equal(t, want.StdTime().Unix(), got.StdTime().Unix())

		t.Log(got.String())
	})
}

func TestDateTEXT(t *testing.T) {
	t.Run("not nil", func(t *testing.T) {
		var got Date

		want := Date(time.Date(2021, 2, 11, 0, 0, 0, 0, time.Local))
		b, err := want.MarshalText()
		require.NoError(t, err)

		err = got.UnmarshalText(b)
		require.NoError(t, err)
		require.Equal(t, want.StdTime().Unix(), got.StdTime().Unix())

		t.Log(got.String())
	})
}
