package models

// User 定义用户结构体
type User struct {
	ID   string
	Name string
}

// UserService 描述与用户相关的服务接口
type UserService interface {
	GetUser(id string) (*User, error)
}
