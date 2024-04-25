package main

import "fmt"

type People struct{}

func (p *People) ShowA() {
	println("showA")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher{}
	// 创建了 Teacher 的实例 t 并调用 t.ShowB() 时，程序将调用 Teacher 中定义的 ShowB() 方法，而不是 People 中的。
	// 这是因为 Teacher 的 ShowB() 方法覆盖（屏蔽）了嵌入的 People 结构体中的同名方法。
	t.ShowB()        // 输出 "teacher showB"
	t.People.ShowB() // showB
}
