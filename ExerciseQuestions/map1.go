package main

import "fmt"

type person struct {
	name string
}

func main() {
	var m map[person]int

	// person 结构体实例化
	p := person{name: "mike"}
	fmt.Println(p)

	// 没有找到 key 为 person 的值，所以返回 0
	fmt.Println(m[p])
}
