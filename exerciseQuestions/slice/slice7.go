package main

import "fmt"

// 定义一个名为 Slice1 的类型，它是一个整数切片
type Slice1 []int

// NewSlice1 函数创建一个新的 Slice1 实例
func NewSlice1() Slice1 {
	// 返回一个长度为0的 Slice1
	return make(Slice1, 0)
}

// Add 方法为 Slice1 类型添加一个元素
// 参数 elem 是要添加的整数
func (s *Slice1) Add(elem int) *Slice1 {
	// 将元素添加到切片中
	*s = append(*s, elem)
	// 输出添加的元素
	fmt.Print(elem)
	// 返回修改后的切片
	return s
}

// main 函数是程序的入口
func main() {
	// 创建一个新的 Slice1 实例
	s := NewSlice1()
	// 使用 defer 延迟执行匿名函数
	defer func() {
		// 在匿名函数中链式调用 Add 方法
		s.Add(1).Add(2)
	}()
	// 立即调用 Add 方法添加元素3
	s.Add(3)
}
