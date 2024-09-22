package service

import (
	"errors"
	"test/ginTest/models"
)

type UserServiceImpl struct{}

func (s *UserServiceImpl) GetUser(id string) (*models.User, error) {
	// 示例：返回一个假用户
	if id == "1" {
		return &models.User{ID: "1", Name: "Alice"}, nil
	}
	return nil, errors.New("user not found")
}
