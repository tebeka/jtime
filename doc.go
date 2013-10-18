package jtime

/* Package jtime provides flexible JSON time.Time marshing/unmarshaling

How it works?

Use jtime.Time in your struct, and before calling json.Marshal/json.Unmarashal
set the marshaller using SetMarshaller.

You can either write your own marshaller or use two of the available marshallers:
* FormatMashaller uses time.Time format to marshal/unmarshal time as JSON strings
* UnixMarshaler marshal time.Time to JSON integers (with or without msec)

jtime.Time embeds time.Time so you can use all time.Time methods with it.

Example:

	package main

	import (
		"encoding/json"
		"fmt"
		"log"

		"bitbucket.org/tebeka/jtime"
	)

	type T struct {
		Created jtime.Time `json:"created"`
	}

	func main() {
		jtime.SetMarshaler(&jtime.UnixMarshaler{})
		data := []byte(`{"created":1382135725}`) // Oct 18, 2013
		t := T{}
		if err := json.Unmarshal(data, &t); err != nil {
			log.Fatalf("error umarshaling: %s\n", err)
		}
		fmt.Println(t.Created) // 2013-10-18 15:35:25 -0700 PDT
	}

Caution: Changing marshaller in mid-flight is dangerous :)
*/
