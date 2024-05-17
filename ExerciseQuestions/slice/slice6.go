package main

import "fmt"

// 定义一个名为 Slice 的类型，它是一个整数切片
type Slice []int

// NewSlice 函数创建一个新的 Slice 实例
func NewSlice() Slice {
	// 返回一个长度为0的 Slice
	return make(Slice, 0)
}

// Add 方法为 Slice 类型添加一个元素
// 参数 elem 是要添加的整数
func (s *Slice) Add(elem int) *Slice {
	// 将元素添加到切片中
	*s = append(*s, elem)
	// 输出添加的元素
	fmt.Println(elem)
	// 返回修改后的切片
	return s
}

// main 函数是程序的入口
func main() {
	// 创建一个新的 Slice 实例
	s := NewSlice()
	// 使用 defer 调用 Add 方法，延迟执行
	defer s.Add(1).Add(2)
	// 立即调用 Add 方法添加元素3
	s.Add(3)
}
