package luoqiangMark

import (
	"fmt"
	"testing"
)

/*
TestDeferFor
以下代码打印什么
5555543210
*/
func TestDeferFor(t *testing.T) {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}

	//这边先还有最后一次i++ 所有5
	for i := 0; i < 5; i++ {
		defer func() {
			fmt.Printf("%d ", i)
		}()
	}
}
