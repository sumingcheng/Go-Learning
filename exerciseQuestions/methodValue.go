package main

import "fmt"

// N2 定义一个新的类型 ，底层类型为 int
type N2 int

// 为 N2 类型定义一个方法 test
func (n N2) test() {
	// 打印当前接收器 n 的值
	fmt.Println(n)
}

func main() {
	var n N2 = 10 // 定义 N2 类型的变量 n, 初始值为 10
	p := &n       // p 是指向 n 的指针

	n++          // n 变为 11
	f1 := n.test // f1 是一个方法值，捕获了此时 n 的值 11

	n++          // n 变为 12
	f2 := p.test // f2 是一个方法值，捕获了此时 n (通过 p 指针) 的值 12

	n++            // n 变为 13
	fmt.Println(n) // 打印 13

	f1() // 调用 f1，打印 11
	f2() // 调用 f2，打印 12
}
