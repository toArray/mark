package main

import (
	"fmt"
	"sync"
)

// 全局变量
var num int64
var wg sync.WaitGroup

func add() {
	for i := 0; i < 10000000; i++ {
		num = num + 1
	}
	// 协程退出， 记录 -1
	wg.Done()
}
func main() {
	// 启动2个协程，记录 2
	wg.Add(2)

	go add()
	go add()

	// 等待子协程退出
	wg.Wait()
	fmt.Println(num)
}
