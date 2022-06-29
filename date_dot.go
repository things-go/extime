package extime

import (
	"errors"
	"time"
)

// DateDot 格式: 2006.01.02
type DateDot time.Time

// ToDateDot time.Time to DateDot
func ToDateDot(t time.Time) DateDot { return DateDot(t) }

// MarshalJSON implemented interface Marshaler
func (t DateDot) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)

	if y := tt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("DateDot.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DateDotLayout)+2)
	b = append(b, '"')
	b = tt.AppendFormat(b, DateDotLayout)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implemented interface Unmarshaler
func (t *DateDot) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.ParseInLocation(`"`+DateDotLayout+`"`, string(data), time.Local)
	*t = DateDot(tt)
	return err
}

// MarshalText implemented interface TextMarshaler
func (t DateDot) MarshalText() ([]byte, error) {
	tt := time.Time(t)

	if y := tt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("DateDot.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DateDotLayout))
	b = tt.AppendFormat(b, DateDotLayout)
	return b, nil
}

// UnmarshalText implemented interface TextUnmarshaler
func (t *DateDot) UnmarshalText(text []byte) error {
	// Ignore null, like in the main JSON package.
	if string(text) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.ParseInLocation(DateDotLayout, string(text), time.Local)
	*t = DateDot(tt)
	return err
}

// StdTime convert to standard time
func (t DateDot) StdTime() time.Time { return time.Time(t) }

// String implemented interface Stringer
func (t DateDot) String() string {
	return time.Time(t).Format(DateDotLayout)
}
