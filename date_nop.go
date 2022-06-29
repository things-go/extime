package extime

import (
	"errors"
	"time"
)

// DateNop 格式: 20060102
type DateNop time.Time

// ToDateNop time.Time to DateNop
func ToDateNop(t time.Time) DateNop { return DateNop(t) }

// MarshalJSON implemented interface Marshaler
func (t DateNop) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)

	if y := tt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("DateNop.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DateNopLayout)+2)
	b = append(b, '"')
	b = tt.AppendFormat(b, DateNopLayout)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implemented interface Unmarshaler
func (t *DateNop) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.ParseInLocation(`"`+DateNopLayout+`"`, string(data), time.Local)
	*t = DateNop(tt)
	return err
}

// MarshalText implemented interface TextMarshaler
func (t DateNop) MarshalText() ([]byte, error) {
	tt := time.Time(t)

	if y := tt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("DateNop.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DateNopLayout))
	b = tt.AppendFormat(b, DateNopLayout)
	return b, nil
}

// UnmarshalText implemented interface TextUnmarshaler
func (t *DateNop) UnmarshalText(text []byte) error {
	// Ignore null, like in the main JSON package.
	if string(text) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.ParseInLocation(DateNopLayout, string(text), time.Local)
	*t = DateNop(tt)
	return err
}

// StdTime convert to standard time
func (t DateNop) StdTime() time.Time { return time.Time(t) }

// String implemented interface Stringer
func (t DateNop) String() string {
	return time.Time(t).Format(DateNopLayout)
}
