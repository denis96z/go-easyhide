//go:build example1

package main

import (
	"fmt"

	"github.com/denis96z/go-easyhide/sample"
)

func main() {
	v1 := sample.T1{
		A1: "value1",
		A2: "value2",
		A3: "value3",
		A4: "value4",
		A8: "v=12345678",
	}

	b, _ := v1.EasyHide()
	fmt.Println(string(b))

	v3 := sample.T3{
		C1: sample.T4{
			D1: "value5",
		},
	}

	b, _ = v3.EasyHide()
	fmt.Println(string(b))
}
