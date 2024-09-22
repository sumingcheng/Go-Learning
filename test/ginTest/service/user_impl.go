package service

import "errors"

type UserServiceImpl struct{}

func (s *UserServiceImpl) GetUser(id string) (*User, error) {
	// 示例：返回一个假用户
	if id == "1" {
		return &User{ID: "1", Name: "Alice"}, nil
	}
	return nil, errors.New("user not found")
}
