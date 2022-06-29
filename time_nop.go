package extime

import (
	"errors"
	"time"
)

// TimeNop 格式: 20060102150405
type TimeNop time.Time

// ToTimeNop time.Time to TimeNop
func ToTimeNop(t time.Time) TimeNop { return TimeNop(t) }

// MarshalJSON implemented interface Marshaler
func (t TimeNop) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)

	if y := tt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("TimeNop.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(TimeNopLayout)+2)
	b = append(b, '"')
	b = tt.AppendFormat(b, TimeNopLayout)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implemented interface Unmarshaler
func (t *TimeNop) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.ParseInLocation(`"`+TimeNopLayout+`"`, string(data), time.Local)
	*t = TimeNop(tt)
	return err
}

// MarshalText implemented interface TextMarshaler
func (t TimeNop) MarshalText() ([]byte, error) {
	tt := time.Time(t)

	if y := tt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(TimeNopLayout))
	b = tt.AppendFormat(b, TimeNopLayout)
	return b, nil
}

// UnmarshalText implemented interface TextUnmarshaler
func (t *TimeNop) UnmarshalText(text []byte) error {
	// Ignore null, like in the main JSON package.
	if string(text) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.ParseInLocation(TimeNopLayout, string(text), time.Local)
	*t = TimeNop(tt)
	return err
}

// StdTime convert to standard time
func (t TimeNop) StdTime() time.Time { return time.Time(t) }

// String implemented interface Stringer
func (t TimeNop) String() string { return time.Time(t).Format(TimeNopLayout) }
