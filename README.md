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
	A8 string `easyhide:"hide:RE,NE:RxpRpl8"`
	A9 string `easyhide:"hide:RES,NE:RxpRpls9"`
}

var (
	RxpRpl8 = easyhide.RegexpReplacement{
		Regexp:      regexp.MustCompile(`^v=(\w{4})\w{4}$`),
		Replacement: `v=${1}` + easyhide.HiddenMarker,
	}

	RxpRpls9 = []easyhide.RegexpReplacement{
		{
			Regexp:      regexp.MustCompile(`password=[^& ]+`),
			Replacement: `password=` + easyhide.HiddenMarker,
		},
		{
			Regexp:      regexp.MustCompile(`token=[^&]+([^&]{4})`),
			Replacement: `token=` + easyhide.HiddenMarker + `${1}`,
		},
	}
)

func ExamplePrint() {
    v1 := sample.T1{
		A1: "value1",
		A2: "value2",
		A3: "value3",
		A4: "value4",
		A8: "v=12345678",
		A9: "https://example.com?k1=v1&password=12345678&k2=v2&token=1234567890",
	}

	b, _ := v1.EasyHide()
	fmt.Println(string(b))
}

//EXPECTED OUTPUT:
//{"a1":"value1","a2":"****","A3":"****ue3","A4":"val****","A5":"","A6":"","A7":"","A8":"v=1234****","A9":"https://example.com?k1=v1\u0026password=****\u0026k2=v2\u0026token=****7890"}
```
