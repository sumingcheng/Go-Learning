package main

import "fmt"

type Person3 struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*Person3)

	// 添加键值对，值为 Person 结构体的指针
	m["john"] = &Person3{"John", 30}

	// 获取 john 对应的 Person 结构体指针
	john := m["john"]

	// 修改 john 指向的结构体的 age 字段
	john.age = 40

	// 此时输出 map 中 "john" 对应的 age 字段值
	fmt.Println(m["john"].age) // 输出 40
}
