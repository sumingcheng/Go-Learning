package main

import "fmt"

type S2 struct{}

// 为 S2 定义指针接收者方法 g
func (s *S2) g() {
	fmt.Println("S2.g() - 指针绑定")
}

func main() {
	s := S2{} // 创建 S2 的值类型实例
	sp := &s  // 获取 S2 的指针

	// s.g() // 值类型调用指针绑定方法，会编译错误

	sp.g() // 通过指针调用指针绑定的方法
}
