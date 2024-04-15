package sample

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
