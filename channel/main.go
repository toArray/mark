package main

import (
	"time"
)

var c = make(chan int)
var a int

func f() {
	a = 1
	<-c
	<-c
	time.Sleep(time.Second * 1000)
}
func main() {
	go f()
	//c <- 0
	print(a)
	time.Sleep(time.Second)
}
