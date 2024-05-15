package main

import (
	"fmt"
	"log"
	"message_push_system/system"
	"os"
	"os/signal"
)

func main() {
	// 创建消息推送系统
	mps, err := system.NewMessagePushSystem()
	if err != nil {
		log.Fatalln(err)
	}
	defer mps.Close()

	// 启动消息队列监控
	go mps.MonitorAndAlert()

	// 发送示例消息
	err = mps.PushMessage(1, "Hello, Kafka!")
	if err != nil {
		log.Fatalln(err)
	}

	// 处理消费者消费消息
	go mps.ConsumeMessages()

	// 阻塞主线程，保持程序运行
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	fmt.Println("Shutting down...")
}
