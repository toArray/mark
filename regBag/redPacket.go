package main

import (
	"errors"
	"math/rand"
)

//RED_PACKET_MIN_MONEY 红包最小金额(单位:分)
const RED_PACKET_MIN_MONEY = 1

//RedPacketModel 红包信息实体
type RedPacketModel struct {
	Count          int64   //红包个数
	Money          int64   //红包金额(单位:分)
	RemainCount    int64   //剩余红包个数
	RemainMoney    int64   //剩余红包金额(单位:分)
	BestLuckMoney  int64   //手气最佳金额(单位:分)
	BestLuckIndex  int64   //手气最佳索引位置
	HistoryRewards []int64 //历史红包记录
}

/*
CreateRedPacket
@Desc 	生成红包
@Param 	count int64	红包个数
@Param 	money int64	红包金额
*/
func CreateRedPacket(count, money int64) (res *RedPacketModel, err error) {
	if count <= 0 || money <= 0 || money < count*RED_PACKET_MIN_MONEY {
		err = errors.New("parameter error")
		return
	}

	res = &RedPacketModel{
		Count:          count,
		Money:          money,
		RemainCount:    count,
		RemainMoney:    money,
		BestLuckMoney:  0,
		BestLuckIndex:  0,
		HistoryRewards: nil,
	}

	return
}

/*
Open
@Desc 	拆红包（二倍均值法）
@Return money 	int64	本地拆红包获得金额
@Return res 	bool	拆红包失败,红包已经被抢光
*/
func (r *RedPacketModel) Open() (money int64, res bool) {
	if r.Check() == false {
		return
	}

	//最后一个红包
	if r.RemainCount == 1 {
		money = r.RemainMoney
	} else {
		//最大可用金额 = 剩余红包金额 - 后续多少个没拆的包所需要的保底金额
		//目的是为了保证后续的包至少都能分到最低保底金额,避免后续未拆的红包出现金额0
		maxCanUseMoney := r.RemainMoney - RED_PACKET_MIN_MONEY*r.RemainCount

		//2倍均值基础金额
		maxAvg := maxCanUseMoney / r.RemainCount

		//2倍均值范围数额
		maxMoney := maxAvg*2 + RED_PACKET_MIN_MONEY

		//随机红包数额
		money = rand.Int63n(maxMoney) + RED_PACKET_MIN_MONEY
	}

	//手气最佳
	if money > r.BestLuckMoney {
		r.BestLuckMoney = money
		r.BestLuckIndex = r.Count - r.RemainCount
	}

	res = true
	r.RemainMoney -= money
	r.RemainCount--
	r.HistoryRewards = append(r.HistoryRewards, money)
	return
}

/*
Check
@Desc 校验红包是否被抢完
*/
func (r *RedPacketModel) Check() bool {
	return r.RemainCount == 0
}
