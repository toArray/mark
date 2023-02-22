package main

import (
	"fmt"
	"testing"
)

func TestAAA(t *testing.T) {
	var a interface{}
	b := int32(1)
	a = b

	c := a.(int32)
	c, ok := a.(int32)
	fmt.Println(c, ok)
	select {}

	fmt.Println(111111)
}
