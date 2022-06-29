package extime

import (
	"errors"
	"time"
)

// DateMonthSlash 格式: 2006.01
type DateMonthDot time.Time

// ToDateMonthDot time.Time to DateMonthDot
func ToDateMonthDot(t time.Time) DateMonthDot { return DateMonthDot(t) }

// MarshalJSON implemented interface Marshaler
func (t DateMonthDot) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)

	if y := tt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("DateMonthDot.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DateMonthDotLayout)+2)
	b = append(b, '"')
	b = tt.AppendFormat(b, DateMonthDotLayout)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implemented interface Unmarshaler
func (t *DateMonthDot) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.ParseInLocation(`"`+DateMonthDotLayout+`"`, string(data), time.Local)
	*t = DateMonthDot(tt)
	return err
}

// MarshalText implemented interface TextMarshaler
func (t DateMonthDot) MarshalText() ([]byte, error) {
	tt := time.Time(t)

	if y := tt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("DateMonthDot.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DateMonthDotLayout))
	b = tt.AppendFormat(b, DateMonthDotLayout)
	return b, nil
}

// UnmarshalText implemented interface TextUnmarshaler
func (t *DateMonthDot) UnmarshalText(text []byte) error {
	// Ignore null, like in the main JSON package.
	if string(text) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.ParseInLocation(DateMonthDotLayout, string(text), time.Local)
	*t = DateMonthDot(tt)
	return err
}

// StdTime convert to standard time
func (t DateMonthDot) StdTime() time.Time { return time.Time(t) }

// String implemented interface Stringer
func (t DateMonthDot) String() string {
	return time.Time(t).Format(DateMonthDotLayout)
}
