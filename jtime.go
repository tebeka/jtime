package jtime

import (
	"fmt"
	"strconv"
	"time"
)

type Time struct {
	time.Time
}

type Marshaler interface {
	Marshal(t Time) ([]byte, error)
	Unmarshal(data []byte) (Time, error)
}

var marshaler Marshaler

func SetMarshaler(m Marshaler) {
	marshaler = m
}

type FormatMashaler struct {
	Format string
}

func (fm *FormatMashaler) Marshal(t Time) ([]byte, error) {
	return []byte(`"` + t.Format(fm.Format) + `"`), nil
}

func (fm *FormatMashaler) Unmarshal(data []byte) (Time, error) {
	if len(data) < 2 {
		return Time{}, fmt.Errorf("data too short - %v", data)
	}
	data = data[1 : len(data)-1]
	t, err := time.Parse(fm.Format, string(data))
	if err != nil {
		return Time{}, err
	}
	return Time{t}, err
}

type UnixMarshaler struct {
	MSec bool
}

func (um *UnixMarshaler) Marshal(t Time) ([]byte, error) {
	data := fmt.Sprintf("%d", t.Unix())
	if um.MSec {
		data = fmt.Sprintf("%s%03d", data, t.Nanosecond()/1000)
	}

	return []byte(data), nil
}

func (um *UnixMarshaler) Unmarshal(data []byte) (Time, error) {
	sec, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return Time{}, err
	}
	nsec := int64(0)
	if um.MSec {
		tmp := sec
		sec = sec / 1000
		nsec = (tmp - sec) * 1000
	}

	return Time{time.Unix(sec, nsec)}, nil
}

func validJSONTime(t Time) bool {
	if y := t.Year(); y < 0 || y >= 10000 {
		return false
	}
	return true
}

// MarshalJSON implements the json.Marshaler interface.
func (t Time) MarshalJSON() ([]byte, error) {
	if !validJSONTime(t) {
		return nil, fmt.Errorf("vtime.Time.MarshalJson: year outside of range [0,9999]")
	}
	return marshaler.Marshal(t)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *Time) UnmarshalJSON(data []byte) (err error) {
	*t, err = marshaler.Unmarshal(data)
	return
}

func init() {
	// Default behaviour
	SetMarshaler(&FormatMashaler{time.RFC3339Nano})
}
