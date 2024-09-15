package main

import "fmt"

func main() {
	m := make(map[string]int)
	// 在Go语言中，当你从一个映射（map）中读取一个不存在的键时
	// Go 不会返回错误，而是返回该值类型的零值。对于int类型，零值是0。
	m["foo"]++
	fmt.Println(m["foo"])
}
