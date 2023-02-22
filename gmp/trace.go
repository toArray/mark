package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {

	//创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	//main
	fmt.Println("Hello World")

	/*
		游戏服获得锁，获得0
		大厅op1 等待锁
		游戏服进入游戏后 sync,redis=0
		游戏服释放锁
		大厅拿到op成功，redis=1
		大厅再次拿数据，判断玩家在游戏内（直接拿游戏内存数据0，不会拿redis）（这中间如果游戏服sync了,+1后的redis又被清空了）



	*/
}
