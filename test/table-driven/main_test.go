package main

import (
	"reflect"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	// 定义参数结构体
	type args struct {
		userID int
	}
	// 定义测试用例
	tests := []struct {
		name    string
		args    args
		want    User
		wantErr bool
	}{
		{"user exists", args{1}, User{ID: 1, Name: "Alice"}, t.Failed()},
		{"user does not exist", args{999}, User{}, t.Failed()},
	}
	// 遍历测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserInfo(tt.args.userID)
			// 检查错误是否符合预期
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// 检查结果是否符合预期
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
