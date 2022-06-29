package extime

import (
	"errors"
	"time"
)

// Date 格式: 2006-01-02
type Date time.Time

// ToDate time.Time to Date
func ToDate(t time.Time) Date { return Date(t) }

// MarshalJSON implemented interface Marshaler
func (t Date) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)

	if y := tt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Date.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DateLayout)+2)
	b = append(b, '"')
	b = tt.AppendFormat(b, DateLayout)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implemented interface Unmarshaler
func (t *Date) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.ParseInLocation(`"`+DateLayout+`"`, string(data), time.Local)
	*t = Date(tt)
	return err
}

// MarshalText implemented interface TextMarshaler
func (t Date) MarshalText() ([]byte, error) {
	tt := time.Time(t)

	if y := tt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("Date.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DateLayout))
	b = tt.AppendFormat(b, DateLayout)
	return b, nil
}

// UnmarshalText implemented interface TextUnmarshaler
func (t *Date) UnmarshalText(text []byte) error {
	// Ignore null, like in the main JSON package.
	if string(text) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.ParseInLocation(DateLayout, string(text), time.Local)
	*t = Date(tt)
	return err
}

// StdTime convert to standard time
func (t Date) StdTime() time.Time { return time.Time(t) }

// String implemented interface Stringer
func (t Date) String() string {
	return time.Time(t).Format(DateLayout)
}
