package main

import "fmt"

type S1 struct{}

// 为 S1 定义值接收者方法 f
func (s S1) f() {
	fmt.Println("S1.f() - 值绑定")
}

func main() {
	s := S1{} // 创建 S1 的值类型实例
	s.f()     // 调用值绑定的方法

	sp := &s // 获取 S1 的指针
	sp.f()   // 通过指针调用值绑定的方法
}
