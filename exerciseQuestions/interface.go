package main

import "fmt"

type A1 interface {
	ShowA() int
}

type B1 interface {
	ShowB() int
}

type Work1 struct {
	i int
}

func (w Work1) ShowA() int {
	return w.i + 10
}

func (w Work1) ShowB() int {
	return w.i + 20
}

func main() {
	var a A1 = Work1{3} // 创建一个 Work1 类型的实例，i 初始化为 3，赋给接口变量 a

	s := a.(Work1) // 类型断言，将接口变量 a 转换回 Work1 类型，存储到变量 s

	fmt.Println(s.ShowA()) // 输出 13
	fmt.Println(s.ShowB()) // 输出 23
}
