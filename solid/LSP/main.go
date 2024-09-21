package main

import "fmt"

// Bird struct 提供了基本的鸟类方法
type Bird struct{}

func (b *Bird) Fly() {
	fmt.Println("Bird is flying")
}

// Duck struct 从 Bird 继承
type Duck struct {
	Bird
}

func main() {
	bird := &Bird{}
	duck := &Duck{}

	bird.Fly()
	duck.Fly() // Duck 类型替换 Bird 类型
}
