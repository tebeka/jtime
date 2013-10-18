# Flexible JSON time Handling

`time.Time` marshal to JSON only in one format (RFC3339Nano).
`jtime.Time` embeds `time.Time` and lets the user set the format via
Marshallers.

# Example

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

# Contact

[https://bitbucket.org/tebeka/jtime](https://bitbucket.org/tebeka/jtime)

# License
MIT (see LICENSE.txt)
