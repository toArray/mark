package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		red, err := CreateRedPacket(10, 1000)
		if err != nil {
			return
		}

		for i := 0; i < 10; i++ {
			money, res := red.Open()
			if res == false {
				break
			}

			fmt.Printf("第%v个人拆红包，金额：%v\n", i+1, float64(money)/100.00)
			if money == 0 {
				panic(1)
			}
		}
		fmt.Println("-----------")
		fmt.Println()
	}

	time.Sleep(time.Second * 10000)
}

/*
getRedBagAverageNum
@Desc 	红包算法（自己思路,垃圾代码）
@Param 	redBagCount int64	红包个数
@Param 	redMoney 	int64	红包金额
*/
func getRedBagAverageNum(redBagCount, redMoney int64) (res []int64, err error) {
	minMoney := int64(1)                //每个红包不低于多少钱
	maxMoney := 200                     //每个红包不超过多少钱
	baseMoney := minMoney * redBagCount //保底金额
	if redMoney < baseMoney {
		//红包金额不足
		err = errors.New("total money too little")
		return
	}

	//每个位置先分配最低金额
	res = make([]int64, redBagCount, redBagCount)
	for i := 0; i < int(redBagCount); i++ {
		res[i] = minMoney
	}

	//剩余多少钱
	leftMoney := redMoney - baseMoney

	//随机分配金额
	for i := 0; i < int(redBagCount); i++ {
		//如果是最后一个红包,直接加金额
		if i == int(redBagCount-1) {
			res[i] += leftMoney
			break
		}

		//随机红包金额
		randNum := rand.Intn(int(leftMoney + 1))
		if err != nil {
			err = errors.New("redBag rand failed")
			break
		}

		addNum := randNum
		if randNum > maxMoney {
			addNum -= maxMoney
		}

		res[i] += int64(randNum)    //红包加金额
		leftMoney -= int64(randNum) //总金额减少
		if leftMoney <= 0 {
			break
		}
	}

	return
}
