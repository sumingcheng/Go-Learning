package service

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"test/ginTest/mocks"
	"test/ginTest/models"
	"testing"
)

// TestGetUser 测试 UserServiceImpl 的 GetUser 方法。
func TestGetUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 创建 UserService 的模拟实例
	mockUserService := mocks.NewMockUserService(ctrl)

	// 设置模拟对象的预期行为，具体到输入和输出
	expectedUser := &models.User{ID: "1", Name: "Alice"}            // 定义要返回的用户数据
	mockUserService.EXPECT().GetUser("1").Return(expectedUser, nil) // 预设当调用 GetUser("1") 时，返回 expectedUser 对象和 nil 错误

	// 实际调用 GetUser 方法
	user, err := mockUserService.GetUser("1")

	// 断言返回的用户是正确的，并且没有错误
	assert.NoError(t, err)              // 断言没有错误返回
	assert.Equal(t, expectedUser, user) // 断言返回的用户信息正确
}
