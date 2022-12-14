package test_kafka

import (
	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"
	"github.com/pkg/errors"
	"luoqiangMark/test_kafka/kafka"
	"time"
)

var kfkConsumer *kafka.Consumer

func mian() {
	opt := &Option{
		Brokers: nil,
		Group:   "",
		Topics:  []string{},
	}
	// 初始化kafka consumer
	err := initKafkaTracerConsumer(opt)
	if err != nil {
		panic(err)
	}
}

func initKafkaTracerConsumer(cfg configs.EntityKafka) error {

	if cfg.Group == "" || len(cfg.Brokers) == 0 {
		return errors.New("kafka 配置错误")
	}

	kfkOpt := &kafka.Option{}
	kfkOpt.Group = cfg.Group
	kfkOpt.Brokers = append(kfkOpt.Brokers, cfg.Brokers...)
	kfkOpt.Topics = []string{
		TopicRegister, //捕鱼注册
		TopicLogin,    // 捕鱼登陆记录
		TopicPay,      // 捕鱼成功下单
		TopicPayOrder, // 捕鱼订单
		// TopicDiamondExchange,  // 捕鱼钻石兑换商品
		// TopicLuckyDraw,        // 捕鱼商城渔券抽奖日志
		// TopicDraw,             // 捕鱼抽奖记录
		// TopicShare,            // 捕鱼分享统计
		// TopicShareAPIUserSync, // 捕鱼shareapi用户同步
		// ToipcBackReward,       // 老用户领取礼包的统计数据
		// TopicUserEvent,        // 用户行为统计数据
		// TopicRedPacket,        // 用户领取红包
		// TopicBrankrupt,        // 破产
		// TopicRoom,             // 捕鱼进出房间
		// TopicCatch,            // 捕鱼击杀鱼日志
		// TopicProp,             // 捕鱼道具变更
		// TopicCost,             // 击杀鱼消耗日志
	}

	var err error
	if kfkConsumer, err = kafka.NewConsumer(kfkOpt); err != nil {
		return err
	}

	go kfkConsumer.Start(internal.RouteMsg)

	return nil
}
