package main

import (
	"fmt"
)

// SliceWrapper 定义一个泛型结构体，用于封装对切片的操作
type SliceWrapper[T any] struct {
	Slice []T
}

func main() {
	// 创建一个整数切片的封装
	intSlice := SliceWrapper[int]{Slice: []int{1, 3, 5}}
	intSlice.Add(2)
	fmt.Println("After adding:", intSlice.Slice)

	intSlice.Remove(1)
	fmt.Println("After removing:", intSlice.Slice)

	intSlice.Push(4)
	fmt.Println("After pushing:", intSlice.Slice)

	// 使用 Map 方法
	newSlice := intSlice.Map(func(x int) int {
		return x * x
	})
	fmt.Println("After mapping:", newSlice.Slice)

}

// Add 添加元素到切片
func (s *SliceWrapper[T]) Add(element T) {
	s.Slice = append(s.Slice, element)
}

// Remove 删除切片中的元素
func (s *SliceWrapper[T]) Remove(index int) error {
	if index < 0 || index >= len(s.Slice) {
		return fmt.Errorf("index out of range")
	}
	s.Slice = append(s.Slice[:index], s.Slice[index+1:]...)
	return nil
}

// Push 添加元素到切片
func (s *SliceWrapper[T]) Push(element T) {
	s.Slice = append(s.Slice, element)
}

// Map 方法，应用一个函数到切片的每个元素
func (s *SliceWrapper[T]) Map(f func(T) T) *SliceWrapper[T] {
	newSlice := make([]T, len(s.Slice))
	for i, v := range s.Slice {
		newSlice[i] = f(v)
	}
	return &SliceWrapper[T]{Slice: newSlice}
}
