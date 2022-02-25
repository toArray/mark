package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(time.Second * time.Duration(i))
			wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}
