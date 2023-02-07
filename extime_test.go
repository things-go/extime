package extime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestValid(t *testing.T) {
	type args struct {
		year  int
		month int
		day   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"正常 ",
			args{
				2019,
				7,
				31,
			},
			true,
		},
		{
			"2月day不正常 ",
			args{
				2019,
				2,
				31,
			},
			false,
		},
		{
			"0月不正常 ",
			args{
				2019,
				0,
				31,
			},
			false,
		},
		{
			"13月不正常 ",
			args{
				2019,
				13,
				31,
			},
			false,
		},
		{
			"0年不正常 ",
			args{
				0,
				1,
				31,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Valid(tt.args.year, tt.args.month, tt.args.day); got != tt.want {
				t.Errorf("Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDays(t *testing.T) {
	from := time.Date(2011, time.January, 1, 0, 0, 0, 0, time.Local)
	to := time.Date(2012, time.January, 4, 12, 1, 2, 100000, time.Local)
	t.Log(Days(to.Sub(from)))
}

func TestYearDays(t *testing.T) {
	require.Equal(t, 366, YearDays(2020))
	require.Equal(t, 365, YearDays(2019))
	require.Equal(t, 366, YearDays(2000))
}

func TestMonthDays(t *testing.T) {
	type args struct {
		year  int
		month time.Month
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1月", args{2020, time.January}, 31},
		{"闰年2月", args{2020, time.February}, 29},
		{"非闰年2月", args{2019, time.February}, 28},
		{"3月", args{2020, time.March}, 31},
		{"4月", args{2020, time.April}, 30},
		{"5月", args{2020, time.May}, 31},
		{"6月", args{2020, time.June}, 30},
		{"7月", args{2020, time.July}, 31},
		{"8月", args{2020, time.August}, 31},
		{"9月", args{2020, time.September}, 30},
		{"10月", args{2020, time.October}, 31},
		{"11月", args{2020, time.November}, 30},
		{"12月", args{2020, time.December}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MonthDays(tt.args.year, tt.args.month); got != tt.want {
				t.Errorf("MonthDays() = %v, want %v", got, tt.want)
			}
		})
	}
	require.Panics(t, func() {
		MonthDays(2020, 13)
	})
}

func TestMonthDays2(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1月", args{time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local)}, 31},
		{"2月", args{time.Date(2020, time.February, 1, 0, 0, 0, 0, time.Local)}, 29},
		{"3月", args{time.Date(2020, time.March, 1, 0, 0, 0, 0, time.Local)}, 31},
		{"4月", args{time.Date(2020, time.April, 1, 0, 0, 0, 0, time.Local)}, 30},
		{"5月", args{time.Date(2020, time.May, 1, 0, 0, 0, 0, time.Local)}, 31},
		{"6月", args{time.Date(2020, time.June, 1, 0, 0, 0, 0, time.Local)}, 30},
		{"7月", args{time.Date(2020, time.July, 1, 0, 0, 0, 0, time.Local)}, 31},
		{"8月", args{time.Date(2020, time.August, 1, 0, 0, 0, 0, time.Local)}, 31},
		{"9月", args{time.Date(2020, time.September, 1, 0, 0, 0, 0, time.Local)}, 30},
		{"10月", args{time.Date(2020, time.October, 1, 0, 0, 0, 0, time.Local)}, 31},
		{"11月", args{time.Date(2020, time.November, 1, 0, 0, 0, 0, time.Local)}, 30},
		{"12月", args{time.Date(2020, time.December, 1, 0, 0, 0, 0, time.Local)}, 31},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MonthDays2(tt.args.t); got != tt.want {
				t.Errorf("MonthDays2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStartAndEndOfDay(t *testing.T) {
	t.Log(StartOfDay(time.Now()))
	t.Log(PreviousStartOfDay(time.Now()))
	t.Log(NextStartOfDay(time.Now()))

	t.Log(EndOfDay(time.Now()))
	t.Log(PreviousEndOfDay(time.Now()))
	t.Log(NextEndOfDay(time.Now()))
}

func TestStartAndEndOfMonth(t *testing.T) {
	t.Log(StartOfMonth(time.Now()))
	t.Log(PreviousStartOfMonth(time.Now()))
	t.Log(NextStartOfMonth(time.Now()))

	t.Log(EndOfMonth(time.Now()))
	t.Log(PreviousEndOfMonth(time.Now()))
	t.Log(NexEndOfMonth(time.Now()))
}

func TestStartAndEndOfYear(t *testing.T) {
	t.Log(StartOfYear(time.Now()))
	t.Log(PreviousStartOfYear(time.Now()))
	t.Log(NextStartOfYear(time.Now()))

	t.Log(EndOfYear(time.Now()))
	t.Log(PreviousEndOfYear(time.Now()))
	t.Log(NextEndOfYear(time.Now()))
}

func TestStartAndEndOfWeek(t *testing.T) {
	t.Log(StartOfWeek(time.Now()))
	t.Log(PreviousStartOfWeek(time.Now()))
	t.Log(NextStartOfWeek(time.Now()))

	t.Log(EndOfWeek(time.Now()))
	t.Log(PreviousEndOfWeek(time.Now()))
	t.Log(NextEndOfWeek(time.Now()))

	t.Log(StartOfWeek(time.Now(), time.Sunday))
	t.Log(PreviousStartOfWeek(time.Now(), time.Sunday))
	t.Log(NextStartOfWeek(time.Now(), time.Sunday))

	t.Log(EndOfWeek(time.Now(), time.Sunday))
	t.Log(PreviousEndOfWeek(time.Now(), time.Sunday))
	t.Log(NextEndOfWeek(time.Now(), time.Sunday))
}

func TestGetQuarter(t *testing.T) {
	require.Equal(t, GetQuarter(1), Quarter1th)
	require.Equal(t, GetQuarter(5), Quarter2th)
	require.Equal(t, GetQuarter(8), Quarter3th)
	require.Equal(t, GetQuarter(11), Quarter4th)

	date := time.Date(2022, 2, 1, 0, 0, 0, 0, time.Local)
	require.Equal(t, GetQuarter2(date), Quarter1th)
}

func TestValidQuarter(t *testing.T) {
	require.True(t, ValidQuarter(1))
	require.True(t, ValidQuarter(2))
	require.True(t, ValidQuarter(3))
	require.True(t, ValidQuarter(4))

	require.False(t, ValidQuarter(0))
	require.False(t, ValidQuarter(5))
}
func TestStartAndEndOfQuarter(t *testing.T) {
	timeQ1 := time.Date(2022, 2, 11, 11, 0, 0, 0, time.Local)
	require.Equal(t, StartOfQuarter(2022, Quarter1th), StartOfQuarter2(timeQ1))
	require.Equal(t, PreviousStartOfQuarter(2022, Quarter1th), PreviousStartOfQuarter2(timeQ1))
	require.Equal(t, NextStartOfQuarter(2022, Quarter1th), NextStartOfQuarter2(timeQ1))
	require.Equal(t, EndOfQuarter(2022, Quarter1th), EndOfQuarter2(timeQ1))
	require.Equal(t, PreviousEndOfQuarter(2022, Quarter1th), PreviousEndOfQuarter2(timeQ1))
	require.Equal(t, NextEndOfQuarter(2022, Quarter1th), NextEndOfQuarter2(timeQ1))

	timeQ2 := time.Date(2022, 4, 21, 11, 0, 0, 0, time.Local)
	require.Equal(t, StartOfQuarter(2022, Quarter2th), StartOfQuarter2(timeQ2))
	require.Equal(t, PreviousStartOfQuarter(2022, Quarter2th), PreviousStartOfQuarter2(timeQ2))
	require.Equal(t, NextStartOfQuarter(2022, Quarter2th), NextStartOfQuarter2(timeQ2))
	require.Equal(t, EndOfQuarter(2022, Quarter2th), EndOfQuarter2(timeQ2))
	require.Equal(t, PreviousEndOfQuarter(2022, Quarter2th), PreviousEndOfQuarter2(timeQ2))
	require.Equal(t, NextEndOfQuarter(2022, Quarter2th), NextEndOfQuarter2(timeQ2))

	timeQ3 := time.Date(2022, 8, 12, 11, 0, 0, 0, time.Local)
	require.Equal(t, StartOfQuarter(2022, Quarter3th), StartOfQuarter2(timeQ3))
	require.Equal(t, PreviousStartOfQuarter(2022, Quarter3th), PreviousStartOfQuarter2(timeQ3))
	require.Equal(t, NextStartOfQuarter(2022, Quarter3th), NextStartOfQuarter2(timeQ3))
	require.Equal(t, EndOfQuarter(2022, Quarter3th), EndOfQuarter2(timeQ3))
	require.Equal(t, PreviousEndOfQuarter(2022, Quarter3th), PreviousEndOfQuarter2(timeQ3))
	require.Equal(t, NextEndOfQuarter(2022, Quarter3th), NextEndOfQuarter2(timeQ3))

	timeQ4 := time.Date(2022, 11, 22, 11, 0, 0, 0, time.Local)
	require.Equal(t, StartOfQuarter(2022, Quarter4th), StartOfQuarter2(timeQ4))
	require.Equal(t, PreviousStartOfQuarter(2022, Quarter4th), PreviousStartOfQuarter2(timeQ4))
	require.Equal(t, NextStartOfQuarter(2022, Quarter4th), NextStartOfQuarter2(timeQ4))
	require.Equal(t, EndOfQuarter(2022, Quarter4th), EndOfQuarter2(timeQ4))
	require.Equal(t, PreviousEndOfQuarter(2022, Quarter4th), PreviousEndOfQuarter2(timeQ4))
	require.Equal(t, NextEndOfQuarter(2022, Quarter4th), NextEndOfQuarter2(timeQ4))
}
