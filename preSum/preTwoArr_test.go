package preSum

import (
	"fmt"
	"testing"
)

//二维数据
var arr [][]int = [][]int{
	{1, 5, 6, 8},
	{9, 6, 7, 3},
	{5, 3, 2, 4},
}

var sum [][]int

//二维数据
func TestTwoArr(t *testing.T) {
	//定义sum[i][j] 是从(0,0)到(i,j)的和
	sum = make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		sum[i] = make([]int, len(arr[i]))
	}

	//前缀和-第一行
	row0 := arr[0]
	for i := 0; i < len(row0); i++ {
		if i == 0 {
			sum[0][i] = row0[i]
		} else {
			sum[0][i] = sum[0][i-1] + row0[i]
		}
	}

	//前缀和-第一列
	for k, v := range arr {
		if k == 0 {
			sum[k][0] = v[0]
		} else {
			sum[k][0] = sum[k-1][0] + v[0]
		}
	}

	//所有前缀和
	for i := 1; i < len(arr); i++ {
		for j := 1; j < len(arr[i]); j++ {
			sum[i][j] = sum[i][j-1] + sum[i-1][j] - sum[i-1][j-1] + arr[i][j] //上面+左边-重复+自己
		}
	}

	fmt.Println(sum)
	fmt.Println(getSum(1, 1, 2, 2))
	fmt.Println(getSum(0, 1, 1, 3))
}

func getSum(x1, y1, x2, y2 int) int {
	if x1 == 0 && y1 == 0 {
		return sum[x2][y2]
	}

	if x1 == 0 {
		return sum[x2][y2] - sum[x2][y1-1]
	}

	if y1 == 0 {
		return sum[x2][y2] - sum[x1-1][y2]
	}

	return sum[x2][y2] - sum[x1-1][y2] - sum[x2][y1-1] + sum[x1-1][y1-1]

}
