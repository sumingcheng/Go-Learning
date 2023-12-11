package main

import (
	"fmt"
	"reflect"
)

func main() {
	user := User{}

	user.SetName("张三")
	user.SetAge(18)
	user.SetHobby("篮球")

	fmt.Printf("%+v\n", user)

	t := reflect.TypeOf(&user)
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		fmt.Println(method.Name)
	}
}

type User struct {
	name  string
	age   int
	hobby string
}

func (u *User) SetName(name string) string {
	u.name = name
	return u.name
}

func (u *User) SetAge(age int) int {
	u.age = age
	return u.age
}

func (u *User) SetHobby(hobby string) string {
	u.hobby = hobby
	return u.hobby
}
