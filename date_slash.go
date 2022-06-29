package extime

import (
	"errors"
	"time"
)

// DateSlash 格式: 2006/01/02
type DateSlash time.Time

// ToDateSlash time.Time to DateSlash
func ToDateSlash(t time.Time) DateSlash { return DateSlash(t) }

// MarshalJSON implemented interface Marshaler
func (t DateSlash) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)

	if y := tt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("DateSlash.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DateSlashLayout)+2)
	b = append(b, '"')
	b = tt.AppendFormat(b, DateSlashLayout)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implemented interface Unmarshaler
func (t *DateSlash) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.ParseInLocation(`"`+DateSlashLayout+`"`, string(data), time.Local)
	*t = DateSlash(tt)
	return err
}

// MarshalText implemented interface TextMarshaler
func (t DateSlash) MarshalText() ([]byte, error) {
	tt := time.Time(t)

	if y := tt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("DateSlash.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DateSlashLayout))
	b = tt.AppendFormat(b, DateSlashLayout)
	return b, nil
}

// UnmarshalText implemented interface TextUnmarshaler
func (t *DateSlash) UnmarshalText(text []byte) error {
	// Ignore null, like in the main JSON package.
	if string(text) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.ParseInLocation(DateSlashLayout, string(text), time.Local)
	*t = DateSlash(tt)
	return err
}

// StdTime convert to standard time
func (t DateSlash) StdTime() time.Time { return time.Time(t) }

// String implemented interface Stringer
func (t DateSlash) String() string {
	return time.Time(t).Format(DateSlashLayout)
}
