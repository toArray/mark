package main

import (
	"fmt"
	"testing"
)

func TestMAN(t *testing.T) {
	var a *int
	a = new(int)
	*a = 1
	fmt.Println(a)
}
