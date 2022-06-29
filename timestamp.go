package extime

import (
	"encoding/json"
	"errors"
	"time"
)

// UnixTimestamp unix 时间戳
type UnixTimestamp time.Time

// ToUnixTimestamp time.Time to UnixTimestamp
func ToUnixTimestamp(t time.Time) UnixTimestamp { return UnixTimestamp(t) }

// MarshalJSON implemented interface Marshaler
func (t UnixTimestamp) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)

	if y := tt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("UnixTimestamp.MarshalJSON: year outside of range [0,9999]")
	}

	return json.Marshal(tt.Unix())
}

// UnmarshalJSON implemented interface Unmarshaler
func (t *UnixTimestamp) UnmarshalJSON(data []byte) error {
	var sec int64
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	err := json.Unmarshal(data, &sec)
	if err != nil {
		return err
	}
	*t = UnixTimestamp(time.Unix(sec, 0))
	return nil
}

// Time convert to standard time
func (t UnixTimestamp) StdTime() time.Time { return time.Time(t) }

// String implemented interface Stringer
func (t UnixTimestamp) String() string { return time.Time(t).String() }

// UnixNanoTimestamp unix nano 时间戳
type UnixNanoTimestamp time.Time

// ToUnixNanoTimestamp time.Time to UnixNanoTimestamp
func ToUnixNanoTimestamp(t time.Time) UnixNanoTimestamp { return UnixNanoTimestamp(t) }

// MarshalJSON implemented interface Marshaler
func (t UnixNanoTimestamp) MarshalJSON() ([]byte, error) {
	tt := time.Time(t)

	if y := tt.Year(); y < 0 || y >= 10000 {
		// RFC 3339 is clear that years are 4 digits exactly.
		// See golang.org/issue/4556#c15 for more discussion.
		return nil, errors.New("UnixNanoTimestamp.MarshalJSON: year outside of range [0,9999]")
	}

	return json.Marshal(tt.UnixNano())
}

// UnmarshalJSON implemented interface Unmarshaler
func (t *UnixNanoTimestamp) UnmarshalJSON(data []byte) error {
	var nano int64
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}
	err := json.Unmarshal(data, &nano)
	if err != nil {
		return err
	}

	*t = UnixNanoTimestamp(time.Unix(nano/int64(time.Second), nano%int64(time.Second)))
	return nil
}

// StdTime convert to standard time
func (t UnixNanoTimestamp) StdTime() time.Time { return time.Time(t) }

// String implemented interface Stringer
func (t UnixNanoTimestamp) String() string { return time.Time(t).String() }
