package main

import "fmt"

// 定义接口 A，要求实现者有一个 showA 方法，该方法返回一个 int 类型的值
type A interface {
	showA() int
}

// 定义接口 B，要求实现者有一个 showB 方法，该方法同样返回一个 int 类型的值
type B interface {
	showB() int
}

// 定义结构体 Work，它有一个 int 类型的字段 i
type Work struct {
	i int
}

// 为 Work 类型实现 showA 方法，满足接口 A 的要求
func (w Work) showA() int {
	return w.i + 10 // 返回字段 i 的值加 10
}

// 为 Work 类型实现 showB 方法，满足接口 B 的要求
func (w Work) showB() int {
	return w.i + 20 // 返回字段 i 的值加 20
}

func main() {
	c := Work{3} // 创建 Work 结构体的实例，其字段 i 被初始化为 3
	var a A = c  // 将 c 赋值给接口类型 A 的变量 a
	var b B = c  // 将 c 赋值给接口类型 B 的变量 b

	// 通过接口 A 调用 c 的 showA 方法
	fmt.Println(a.showA()) // 输出 13，即 3 + 10

	// 通过接口 B 调用 c 的 showB 方法
	fmt.Println(b.showB()) // 输出 23，即 3 + 20
}
