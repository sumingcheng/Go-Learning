package main

import (
	"reflect"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		args    args
		want    User
		wantErr bool
	}{
		{"user exists", args{1}, User{ID: 1, Name: "Alice"}, t.Failed()},
		{"user does not exist", args{999}, User{}, t.Failed()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetUserInfo(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserInfo() got = %v, want %v", got, tt.want)
			}
		})
	}
}
