package main

import "fmt"

type TimesMatcher struct {
	base int
}

// NewTimesMatcher 是一个构造函数，接受一个 int 类型的参数 base，
// 并返回一个指向 TimesMatcher 实例的指针。
func NewTimesMatcher(base int) *TimesMatcher {
	return &TimesMatcher{base: base}
}

// main 函数是程序的入口点。
func main() {
	// 调用 NewTimesMatcher 函数创建一个 TimesMatcher 实例，
	// 其 base 字段被初始化为 3。
	p := NewTimesMatcher(3)

	// 输出 p 指针指向的 TimesMatcher 实例的内容。
	// 这里将打印出 TimesMatcher 结构体实例的信息，包括 base 字段的值。
	fmt.Println(p)
}
