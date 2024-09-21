package main

import "fmt"

type User struct {
	ID   int
	Name string
}

var users = map[int]User{
	1: {ID: 1, Name: "Alice"},
	2: {ID: 2, Name: "Bob"},
}

func GetUserInfo(userID int) (User, error) {
	if user, ok := users[userID]; ok {
		return user, nil
	}
	return User{}, fmt.Errorf("user not found")
}
