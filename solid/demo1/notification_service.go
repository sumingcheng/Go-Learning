package main

import (
	"errors"
)

// NotificationService 是一个结构体，包含了一个 MessageSender 类型的字段 sender
type NotificationService struct {
	sender MessageSender
}

// NewNotificationService 是一个构造函数，用于创建 NotificationService 实例
func NewNotificationService(sender MessageSender) *NotificationService {
	return &NotificationService{sender: sender}
}

// Notify 是 NotificationService 结构体的一个方法，用于发送通知
func (n *NotificationService) Notify(message string) error {
	if message == "" {
		return errors.New("message cannot be empty")
	}
	return n.sender.Send(message)
}
