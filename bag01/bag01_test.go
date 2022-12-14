package main

import (
	"fmt"
	"testing"
)

func TestBag01(t *testing.T) {
	weight := []int{5, 3, 4, 2}     //物品重量
	value2 := []int{15, 20, 30, 40} //物品价值
	bag := 4                        //背包容量

	//dp,前i个物品放进容量为j的最大价值
	dp := make([][]int, len(value2))
	for i := 0; i < len(value2); i++ {
		dp[i] = []int{0, 0, 0, 0, 0}
	}

	//循环物品
	for i := 0; i < len(weight); i++ {
		for j := 1; j <= bag; j++ {
			if j >= weight[i] {
				remainWeight := j - weight[i]
				if i-1 >= 0 {
					maxValue := max(value2[i]+dp[i-1][remainWeight], dp[i-1][j]) //(当前物品价值+剩余容量最大价值，上一个空间最大价值)
					dp[i][j] = maxValue
				} else {
					dp[i][j] = value2[i]
				}
			} else {
				if i-1 >= 0 {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = 0
				}
			}
		}
	}

	fmt.Println(dp, weight, bag, dp[len(weight)-1][bag])
}
