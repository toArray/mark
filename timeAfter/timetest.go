package main

import (
	"fmt"
	"time"
)

func main() {
	//closeChannel()
	c := make(chan int, 1)
	timeout := time.After(time.Second * 2) //
	t1 := time.NewTimer(time.Second * 6)   // 效果相同 只执行一次
	var i int
	//go func() {
	for {
		select {
		case <-c:
			fmt.Println("channel sign")
			return
		case <-t1.C: // 代码段2
			fmt.Println("6s定时任务")
			return
		case <-timeout: // 代码段1
			i++
			fmt.Println(i, "2s定时输出")
			return
		}
	}
	//}()
	time.Sleep(time.Second * 10)
	close(c)
	time.Sleep(time.Second * 2)
	fmt.Println("main退出")
}
