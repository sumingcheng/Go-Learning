package main

import "fmt"

// 定义一个接口，只要实现了Speak方法的类型，都可以被认为是Animal1类型
type Animal1 interface {
	Speak() string
}

// Dog1 结构体
type Dog1 struct{}

// Dog1 类型实现 Animal1 接口
func (d Dog1) Speak() string {
	return "汪汪"
}

// Cat1 结构体
type Cat1 struct{}

// Cat1 类型实现 Animal1 接口
func (c Cat1) Speak() string {
	return "喵喵"
}

func main() {
	// 创建一个Animal1类型的切片，里面可以存储任何实现了Animal1接口的类型的实例
	// 运行时动态地调用不同类型的 Speak 方法，体现了动态类型的行为。
	animals := []Animal1{Dog1{}, Cat1{}}

	// 遍历这个切片，并调用每个元素的Speak方法
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
