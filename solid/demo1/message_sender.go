package main

import "errors"

// MessageSender 是一个接口，定义了发送消息的方法
type MessageSender interface {
	Send(message string) error
}

// CreateMessageSender 根据配置创建对应的 MessageSender 实现
func CreateMessageSender(config Config) (MessageSender, error) {
	switch config.NotificationType {
	case "email":
		return &EmailSender{}, nil
	case "sms":
		return &SmsSender{}, nil
	default:
		return nil, errors.New("unsupported notification type")
	}
}
