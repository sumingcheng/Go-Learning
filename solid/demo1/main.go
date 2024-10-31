package main

import (
	"log"
)

func main() {
	// 配置管理
	config := Config{NotificationType: "email"} // 可以改为 "sms"

	// 根据配置创建消息发送者
	sender, err := CreateMessageSender(config)
	if err != nil {
		log.Fatalf("Error creating message sender: %v", err)
	}

	// 创建通知服务实例
	notificationService := NewNotificationService(sender)

	// 发送通知
	if err := notificationService.Notify("Hello via Notification Service"); err != nil {
		log.Fatalf("Error sending notification: %v", err)
	}

	// 发送空消息，触发错误
	if err := notificationService.Notify(""); err != nil {
		log.Printf("Expected error: %v", err)
	}
}
