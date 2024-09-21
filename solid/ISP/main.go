package main

import "fmt"

// Worker interface 定义了工作相关的方法
type Worker interface {
	Work()
}

// Eater interface 定义了吃饭相关的方法
type Eater interface {
	Eat()
}

// Human 实现了 Worker 和 Eater interface
type Human struct{}

func (h *Human) Work() {
	fmt.Println("Human is working")
}

func (h *Human) Eat() {
	fmt.Println("Human is eating")
}

func main() {
	human := &Human{}
	human.Work()
	human.Eat()
}
