package syncTest

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestOnce(t *testing.T) {
	once := sync.Once{}

	for i := 1; i < 100; i++ {
		once.Do(func() {
			fmt.Println("只执行一次,", i)
		})
		once.Do(func() {
			fmt.Println("这个不会执行,", i)
		})
	}
	//once.
	time.Sleep(time.Second * 100)
}
