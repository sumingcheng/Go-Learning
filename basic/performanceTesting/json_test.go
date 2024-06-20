package main

import (
	"encoding/json"
	"github.com/bytedance/sonic"
	"testing"
)

type Student struct {
	Name   string
	Age    int
	Gender string
}

type Class struct {
	Id       string
	Students []Student
}

var (
	s = Student{"张三", 18, "女"}
	c = Class{
		Id:       "1年2班",
		Students: []Student{s, s, s},
	}
)

func BenchmarkJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes, _ := json.Marshal(c)
		var c2 Class
		json.Unmarshal(bytes, &c2)
	}
}

func BenchmarkSonic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes, _ := sonic.Marshal(c)
		var c2 Class
		sonic.Unmarshal(bytes, &c2)
	}
}
