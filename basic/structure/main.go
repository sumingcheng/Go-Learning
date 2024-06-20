package main

import (
	"fmt"
	"github.com/bytedance/sonic"
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

func main() {
	s := Student{"张三", 18, "女"}
	c := Class{
		Id:       "1年2班",
		Students: []Student{s, s, s},
	}

	// JSON 序列化
	bytes, err := sonic.Marshal(c)
	if err != nil {
		fmt.Println("JSON 序列化失败", err)
		return
	}

	str := string(bytes) // 当前 bytes 就是一个二进制的字节流
	fmt.Println(str)

	//	反序列化
	var c2 Class
	err = sonic.Unmarshal(bytes, &c2)
	if err != nil {
		fmt.Println("JSON 反序列化失败", err)
		return
	}

	fmt.Printf("%+v\n", c2)
}
