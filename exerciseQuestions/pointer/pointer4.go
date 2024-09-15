package main

import (
	"fmt"
)

type N4 int

// 值接收者方法
func (n N4) value() {
	n++
	fmt.Printf("value方法 - 值接收者:\n内存地址: %p, 值: %v\n", &n, n)
}

// 指针接收者方法
func (n *N4) pointer() {
	*n++
	fmt.Printf("pointer方法 - 指针接收者:\n内存地址: %p, 值: %v\n", n, *n)
}

func main() {
	var a N4 = 25

	p := &a // p是指向a的指针

	// 正确调用方式
	a.value()   // 调用值接收者方法    值接收者方法：操作的是调用者的副本，原始值不变。
	p.pointer() // 调用指针接收者方法  指针接收者方法：操作的是调用者的引用，原始值会被修改。

	// 输出a的最终值
	fmt.Printf("main函数 - 变量a:\n内存地址: %p, 值: %v\n", &a, a)
}
