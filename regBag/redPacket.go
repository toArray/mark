package main

import (
	"errors"
	"math/rand"
)

//RED_PACKET_MIN_MONEY 红包最小金额(单位:分)
const RED_PACKET_MIN_MONEY = 1

//RedPacketModel 红包信息实体
type RedPacketModel struct {
	count          int64   //红包个数
	money          int64   //红包金额(单位:分)
	remainCount    int64   //剩余红包个数
	remainMoney    int64   //剩余红包金额(单位:分)
	bestLuckMoney  int64   //手气最佳金额(单位:分)
	bestLuckIndex  int64   //手气最佳索引位置
	historyRewards []int64 //历史红包记录
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
		count:          count,
		money:          money,
		remainCount:    count,
		remainMoney:    money,
		bestLuckMoney:  0,
		bestLuckIndex:  0,
		historyRewards: nil,
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
	//检测红包是否被抢光
	if r.Check() {
		return
	}

	//最后一个红包
	if r.remainCount == 1 {
		money = r.remainMoney
	} else {
		//最大可用金额 = 剩余红包金额 - 后续多少个没拆的包所需要的保底金额
		//目的是为了保证后续的包至少都能分到最低保底金额,避免后续未拆的红包出现金额0
		maxCanUseMoney := r.remainMoney - RED_PACKET_MIN_MONEY*r.remainCount

		//2倍均值基础金额
		maxAvg := maxCanUseMoney / r.remainCount

		//2倍均值范围数额
		maxMoney := maxAvg*2 + RED_PACKET_MIN_MONEY

		//随机红包数额
		money = rand.Int63n(maxMoney) + RED_PACKET_MIN_MONEY
	}

	//手气最佳
	if money > r.bestLuckMoney {
		r.bestLuckMoney = money
		r.bestLuckIndex = r.count - r.remainCount
	}

	res = true
	r.remainMoney -= money
	r.remainCount--
	r.historyRewards = append(r.historyRewards, money)
	return
}

/*
Check
@Desc 校验红包是否被抢完
*/
func (r *RedPacketModel) Check() bool {
	return r.remainCount == 0
}
