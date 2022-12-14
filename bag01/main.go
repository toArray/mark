package main

import "fmt"

func main() {
	weight := []int{1, 3, 4}   //物品重量
	value := []int{15, 20, 30} //物品价值
	bag := 4                   //背包容量

	//定义dp数组
	//dp[i][j] 前i个物品放进容量为j的背包的最大价值
	dp := make([][]int, len(weight))
	for i := 0; i < len(weight); i++ {
		dp[i] = make([]int, bag+1)
	}

	fmt.Println(dp)

	//初始化dp数组
	for j := 0; j <= bag; j++ {
		if weight[0] <= j {
			dp[0][j] = value[0]
		}
	}

	//循环物品，第一个物品已经初始化过了,所以这个从第二个物品开始
	for i := 1; i < len(weight); i++ {
		//循环背包容量
		for j := 1; j <= bag; j++ {
			if weight[i] <= j {
				//当前物品放的下
				remainWeight := j - weight[i]                               //剩余容量
				maxValue := max(value[i]+dp[i-1][remainWeight], dp[i-1][j]) //(当前物品价值+剩余容量最大价值，上一个空间最大价值)
				dp[i][j] = maxValue

			} else {
				//当前物品放不下，取上一个空间最大值
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	fmt.Println(dp[len(weight)-1][bag])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
