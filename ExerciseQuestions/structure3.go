package main

import "fmt"

type People3 struct{}

func (p *People3) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

func (p *People3) ShowB() {
	fmt.Println("showB")
}

type Teacher3 struct {
	People3
}

func (t *Teacher3) ShowB() {
	fmt.Println("teacher showB")
}

func main() {
	t := Teacher3{}
	// t 没有自己的 A 方法，所以调用People3 的A 方法
	t.ShowA()
}
