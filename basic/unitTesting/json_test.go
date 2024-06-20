package main

import (
	"encoding/json"
	"fmt"
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

func TestJson(t *testing.T) {
	bytes, err := json.Marshal(c)
	if err != nil {
		t.Fail()
	}

	var c2 Class
	err = json.Unmarshal(bytes, &c2)
	if err != nil {
		t.Fail()
	}
	if !(c.Id == c2.Id && len(c.Students) == len(c2.Students)) {
		t.Fail()
	}
	fmt.Println("JSON 测试通过")
}

func TestSonic(t *testing.T) {
	bytes, err := sonic.Marshal(c)
	if err != nil {
		t.Fail()
	}

	var c2 Class
	err = sonic.Unmarshal(bytes, &c2)
	if err != nil {
		t.Fail()
	}
	if !(c.Id == c2.Id && len(c.Students) == len(c2.Students)) {
		t.Fail()
	}
	fmt.Println("sonic 测试通过")
}
