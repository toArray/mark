package main

import (
	"fmt"
	"go.uber.org/zap"
	"time"
)

func main() {
	logger, _ := zap.NewDevelopment()
	logger.Debug("err", zap.String("test", "test"), zap.String("test2", "test2"))
	logger.Error("err", zap.String("test", "test"), zap.String("test2", "test2"))
	logger.Panic("err", zap.String("test", "test"), zap.String("test2", "test2"))
	return

	go func() {
		for {
			time.Sleep(time.Millisecond * 2000)
			fmt.Println("test")
		}
	}()

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		panic("panic go test")
	}()

	for {
		select {
		default:

		}
	}

}
