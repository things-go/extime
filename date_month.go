package extime

import (
	"errors"
	"time"
)

// DateMonth 格式: 2006-01
type DateMonth time.Time

// ToDateMonth time.Time to DateMonth
func ToDateMonth(t time.Time) DateMonth { return DateMonth(t) }

// MarshalJSON implemented interface Marshaler
func (t DateMonth) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)

	if y := tt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("DateMonth.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DateMonthLayout)+2)
	b = append(b, '"')
	b = tt.AppendFormat(b, DateMonthLayout)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implemented interface Unmarshaler
func (t *DateMonth) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.ParseInLocation(`"`+DateMonthLayout+`"`, string(data), time.Local)
	*t = DateMonth(tt)
	return err
}

// MarshalText implemented interface TextMarshaler
func (t DateMonth) MarshalText() ([]byte, error) {
	tt := time.Time(t)

	if y := tt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("DateMonth.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DateMonthLayout))
	b = tt.AppendFormat(b, DateMonthLayout)
	return b, nil
}

// UnmarshalText implemented interface TextUnmarshaler
func (t *DateMonth) UnmarshalText(text []byte) error {
	// Ignore null, like in the main JSON package.
	if string(text) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.ParseInLocation(DateMonthLayout, string(text), time.Local)
	*t = DateMonth(tt)
	return err
}

// StdTime convert to standard time
func (t DateMonth) StdTime() time.Time { return time.Time(t) }

// String implemented interface Stringer
func (t DateMonth) String() string {
	return time.Time(t).Format(DateMonthLayout)
}
