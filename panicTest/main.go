package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			time.Sleep(time.Millisecond * 2000)
			fmt.Println("test")
		}
	}()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		panic("panic go test")
	}()

	for {
		select {
		default:

		}
	}

}
