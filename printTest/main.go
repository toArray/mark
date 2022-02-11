package main

import "fmt"

//打印什么 var init main
//常量、变量、init()、main() 依次进行初始化

var a = func() int {
	fmt.Println("var")
	return 0
}()

func init() {
	fmt.Println("init")

}

func main() {
	fmt.Println("main")
}
