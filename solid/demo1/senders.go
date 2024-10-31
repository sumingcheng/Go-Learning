package main

import (
	"log"
)

// EmailSender 发送电子邮件
type EmailSender struct{}

// Send 实现了 MessageSender 接口
func (e *EmailSender) Send(message string) error {
	// 模拟发送电子邮件的逻辑
	log.Printf("Sending email: %s", message)
	return nil // 假设发送成功
}

// SmsSender 发送短信
type SmsSender struct{}

// Send 实现了 MessageSender 接口
func (s *SmsSender) Send(message string) error {
	// 模拟发送短信的逻辑
	log.Printf("Sending SMS: %s", message)
	return nil // 假设发送成功
}
