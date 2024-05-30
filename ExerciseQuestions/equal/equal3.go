package main

import (
	"fmt"
)

type Person2 struct {
	Name string
	Age  int
}

func main() {
	// 整数比较
	a := 10
	b := 10
	fmt.Println(a == b) // 输出: true

	// 结构体比较
	p1 := Person2{Name: "John", Age: 30}
	p2 := Person2{Name: "John", Age: 30}
	fmt.Println(p1 == p2) // 输出: true

	// 接口比较
	var i1 interface{} = 100
	var i2 interface{} = 100
	fmt.Println(i1 == i2) // 输出: true

	// 指针比较
	x := 5
	ptr1 := &x
	ptr2 := &x
	fmt.Println(ptr1 == ptr2) // 输出: true

	// 切片和映射比较（只能与nil比较）
	s1 := []int{1, 2, 3}
	var s2 []int = nil
	fmt.Println(s1 == nil, s2 == nil) // 输出: false true
}
