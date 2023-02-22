package main

import (
	"sync"
	"testing"
)

const N = 10

var wg = &sync.WaitGroup{}

func TestAaaa(t *testing.T) {
	for i := 0; i < N; i++ {
		go func(i int) {
			wg.Add(1)
			println(i)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}
