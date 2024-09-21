package main

import (
	"fmt"
)

// Logger struct 仅负责日志记录的职责
type Logger struct{}

func (l *Logger) Log(message string) {
	fmt.Println("Log:", message)
}

// UserManager struct 负责用户管理，不负责日志记录
type UserManager struct {
	logger *Logger
}

func (u *UserManager) AddUser(name string) {
	u.logger.Log("Adding user: " + name)
	// 添加用户的逻辑
}

func main() {
	logger := &Logger{}
	userManager := &UserManager{logger: logger}
	userManager.AddUser("Alice")
}
