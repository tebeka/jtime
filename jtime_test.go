package jtime

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

const (
	testYear = 2013
)

var testYearStr = fmt.Sprintf(`"%d"`, testYear)

type TestStruct struct {
	Created Time `json:"created"`
}

func testTime() Time {
	tt, _ := time.ParseInLocation(`"2006"`, testYearStr, time.UTC)
	return Time{tt}
}

func TestFormatMarshaller(t *testing.T) {
	tm := testTime()
	fm := &FormatMashaler{"2006"}
	data, err := fm.Marshal(tm)
	if err != nil {
		t.Fatalf("can't marshal: %s", err)
	}

	if string(data) != testYearStr {
		t.Fatalf("bad marshal: %s", string(data))
	}

	pt, err := fm.Unmarshal(data)
	if err != nil {
		t.Fatalf("can't unmarshal: %s", err)
	}

	if pt.Year() != testYear {
		t.Fatalf("bad year: %d", pt.Year())
	}
}

func checkUnmarshal(data []byte, t *testing.T) {
	ts1 := TestStruct{}
	err := json.Unmarshal(data, &ts1)
	if err != nil {
		t.Fatalf("can't unmarshal - %s", err)
	}

	if ts1.Created.In(time.UTC).Year() != testYear {
		t.Fatalf("bad unmarshaled year - %d", ts1.Created.Year())
	}
}

func TestJSON(t *testing.T) {
	SetMarshaler(&FormatMashaler{"2006"})

	ts := TestStruct{testTime()}
	data, err := json.Marshal(ts)
	if err != nil {
		t.Fatalf("can't json.Marshal - %s", err)
	}
	if string(data) != `{"created":"2013"}` {
		t.Fatalf("bad json encoding: %s", string(data))
	}

	checkUnmarshal(data, t)
}

func TestJSONUnix(t *testing.T) {
	SetMarshaler(&UnixMarshaler{})
	ts := TestStruct{testTime()}
	data, err := json.Marshal(ts)
	if err != nil {
		t.Fatalf("can't json.Marshal - %s", err)
	}
	if string(data) != `{"created":1356998400}` {
		t.Fatalf("bad JSONUnix marshal - %s", string(data))
	}

	checkUnmarshal(data, t)
}

func TestJSONUnixMSec(t *testing.T) {
	SetMarshaler(&UnixMarshaler{true})
	ts := TestStruct{testTime()}
	data, err := json.Marshal(ts)
	if err != nil {
		t.Fatalf("can't json.Marshal - %s", err)
	}
	if string(data) != `{"created":1356998400000}` {
		t.Fatalf("bad JSONUnix marshal - %s", string(data))
	}

	checkUnmarshal(data, t)
}
