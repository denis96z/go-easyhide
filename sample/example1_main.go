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
		A9: "https://example.com?k1=v1&password=12345678&k2=v2&token=1234567890",
	}

	b, _ := v1.EasyHide()
	fmt.Println(string(b))
}
