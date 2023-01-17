package extime

import (
	"errors"
	"time"
)

// Time 格式: 2006-01-02 15:04:05
type Time time.Time

// ToTime time.Time to Time
func ToTime(t time.Time) Time { return Time(t) }

// MarshalJSON implemented interface Marshaler
func (t Time) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)

	if y := tt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(TimeLayout)+2)
	b = append(b, '"')
	b = tt.AppendFormat(b, TimeLayout)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implemented interface Unmarshaler
func (t *Time) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.ParseInLocation(`"`+TimeLayout+`"`, string(data), time.Local)
	*t = Time(tt)
	return err
}

// MarshalText implemented interface TextMarshaler
func (t Time) MarshalText() ([]byte, error) {
	tt := time.Time(t)

	if y := tt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Time.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(TimeLayout))
	b = tt.AppendFormat(b, TimeLayout)
	return b, nil
}

// UnmarshalText implemented interface TextUnmarshaler
func (t *Time) UnmarshalText(text []byte) error {
	// Ignore null, like in the main JSON package.
	if string(text) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.ParseInLocation(TimeLayout, string(text), time.Local)
	*t = Time(tt)
	return err
}

// StdTime convert to standard time
func (t Time) StdTime() time.Time { return time.Time(t) }

// String implemented interface Stringer
func (t Time) String() string { return time.Time(t).Format(TimeLayout) }

// ParseTime parse time with layout 2006-01-02 15:04:05
func ParseTime(value string) (Time, error) {
	t, err := time.ParseInLocation(TimeLayout, value, time.Local)
	if err != nil {
		return Time{}, err
	}
	return Time(t), nil
}
