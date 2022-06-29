package extime

import (
	"errors"
	"time"
)

// DateMonthNop 格式: 200601
type DateMonthNop time.Time

// ToDateMonthNop time.Time to DateMonthNop
func ToDateMonthNop(t time.Time) DateMonthNop { return DateMonthNop(t) }

// MarshalJSON implemented interface Marshaler
func (t DateMonthNop) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)

	if y := tt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("DateMonthNop.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DateMonthNopLayout)+2)
	b = append(b, '"')
	b = tt.AppendFormat(b, DateMonthNopLayout)
	b = append(b, '"')
	return b, nil
}

// UnmarshalJSON implemented interface Unmarshaler
func (t *DateMonthNop) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.ParseInLocation(`"`+DateMonthNopLayout+`"`, string(data), time.Local)
	*t = DateMonthNop(tt)
	return err
}

// MarshalText implemented interface TextMarshaler
func (t DateMonthNop) MarshalText() ([]byte, error) {
	tt := time.Time(t)

	if y := tt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("DateMonthNop.MarshalJSON: year outside of range [0,9999]")
	}

	b := make([]byte, 0, len(DateMonthNopLayout))
	b = tt.AppendFormat(b, DateMonthNopLayout)
	return b, nil
}

// UnmarshalText implemented interface TextUnmarshaler
func (t *DateMonthNop) UnmarshalText(text []byte) error {
	// Ignore null, like in the main JSON package.
	if string(text) == "null" {
		return nil
	}
	// Fractional seconds are handled implicitly by Parse.
	tt, err := time.ParseInLocation(DateMonthNopLayout, string(text), time.Local)
	*t = DateMonthNop(tt)
	return err
}

// StdTime convert to standard time
func (t DateMonthNop) StdTime() time.Time { return time.Time(t) }

// String implemented interface Stringer
func (t DateMonthNop) String() string {
	return time.Time(t).Format(DateMonthNopLayout)
}
