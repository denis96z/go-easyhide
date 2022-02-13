//go:build example1

package main

import (
	"fmt"

	"go-easyhide/sample"
)

func main() {
	v1 := sample.T1{
		A1: "value1", //will be shown
		A2: "value2", //will be hidden
		A3: "value3", //half chars will be hidden
		A4: "value4", //half chars will be hidden
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
