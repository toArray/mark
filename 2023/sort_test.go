package main2023

import (
	"fmt"
	"testing"
)

func BenchmarkLoopStep2(b *testing.B) {
	//制作源数据，长度为10000
	//src := CreateSource(10000)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fmt.Println(11111)
	}
}

func TestMaopao(t *testing.T) {
	arr := []int32{3, 1, 34, 5, 6, 1}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	fmt.Println(arr)
}

func TestQK(t *testing.T) {
	arr := []int32{3, 1, 34, 5, 6, 1}

	fmt.Println(q(arr))
}
func q(arr []int32) []int32 {
	if len(arr) <= 1 {
		return arr
	}
	mid := arr[0]
	left := []int32{}
	right := []int32{}

	for i := 1; i < len(arr); i++ {
		if arr[i] < mid {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}

	fmt.Println(1)

	left = q(left)
	right = q(right)
	res := make([]int32, 0)
	res = append(res, left...)
	res = append(res, mid)
	res = append(res, right...)
	return res
}
