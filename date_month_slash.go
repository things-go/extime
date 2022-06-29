package extime

import (
	"errors"
	"time"
)

// DateMonthSlash 格式: 2006/01
type DateMonthSlash time.Time

// ToDateMonthSlash time.Time to DateMonthSlash
func ToDateMonthSlash(t time.Time) DateMonthSlash { return DateMonthSlash(t) }

// MarshalJSON implemented interface Marshaler
func (t DateMonthSlash) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)

	if y := tt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("DateMonthSlash.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DateMonthSlashLayout)+2)
	b = append(b, '"')
	b = tt.AppendFormat(b, DateMonthSlashLayout)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implemented interface Unmarshaler
func (t *DateMonthSlash) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.ParseInLocation(`"`+DateMonthSlashLayout+`"`, string(data), time.Local)
	*t = DateMonthSlash(tt)
	return err
}

// MarshalText implemented interface TextMarshaler
func (t DateMonthSlash) MarshalText() ([]byte, error) {
	tt := time.Time(t)

	if y := tt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("DateMonthSlash.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DateMonthSlashLayout))
	b = tt.AppendFormat(b, DateMonthSlashLayout)
	return b, nil
}

// UnmarshalText implemented interface TextUnmarshaler
func (t *DateMonthSlash) UnmarshalText(text []byte) error {
	// Ignore null, like in the main JSON package.
	if string(text) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.ParseInLocation(DateMonthSlashLayout, string(text), time.Local)
	*t = DateMonthSlash(tt)
	return err
}

// StdTime convert to standard time
func (t DateMonthSlash) StdTime() time.Time { return time.Time(t) }

// String implemented interface Stringer
func (t DateMonthSlash) String() string {
	return time.Time(t).Format(DateMonthSlashLayout)
}
