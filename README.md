# go-easyhide (WORK IN PROGRESS)

## Example
```
//go:generate easyhide

import (
	"regexp"

	"github.com/denis96z/go-easyhide/pkg/easyhide"
)

//easyhide:json
type T1 struct {
	A1 string `json:"a1" easyhide:"show"`
	A2 string `json:"a2" easyhide:"hide"`
	A3 string `easyhide:"hide:HL"`
	A4 string `easyhide:"hide:HR"`
	A5 string `easyhide:"hide:NE"`
	A6 string `easyhide:"hide:HL,NE"`
	A7 string `easyhide:"hide:HR,NE"`
	A8 string `easyhide:"hide:RE,NE:Rxp8:RxpRpl8"`
}

var (
	Rxp8    = regexp.MustCompile(`^v=(\w{4})\w{4}$`)
	RxpRpl8 = `v=${1}` + easyhide.HiddenMarker
)

func ExamplePrint() {
    v1 := sample.T1{
		A1: "value1",
		A2: "value2",
		A3: "value3",
		A4: "value4",
		A8: "v=12345678",
	}

    b, _ := v.EasyHide()
    fmt.Println(string(b))
}

//EXPECTED OUTPUT:
//{"a1":"value1","a2":"****","A3":"****ue3","A4":"val****","A5":"","A6":"","A7":"","A8":"v=1234****"}
```
