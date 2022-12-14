package preSum

import (
	"fmt"
	"testing"
)

//差分数组
var d []int

//一维前缀和差分
func TestPreD(t *testing.T) {
	//原始数组
	arr := []int{1, 3, 7, 5, 2}

	//原始数组差分值
	d = make([]int, len(arr)+1) //[0,0,0,0,0]
	add(2, 4, 5)                //[0,0,5,0,0]
	add(1, 3, 2)                //[0,2,5,0,-2]
	add(0, 2, -3)               //[-3,2,5,-3,-2]

	//差分前缀和
	for i := 1; i < len(d); i++ {
		d[i] += d[i-1]
	}

	//累加原始数组
	for i := 0; i < len(arr); i++ {
		arr[i] += d[i]
	}

	fmt.Println(arr)
}

func add(l, r, v int) {
	d[l] += v
	d[r+1] -= v
}
