package main

import "fmt"

type People1 interface {
	show()
}

type Student1 struct {
}

func (stu *Student1) show() {

}

/*
第一次nil检查是针对s，它输出"s is nil"，这是因为s确实是一个空指针，即指向Student1类型的nil指针。

第二次nil检查是针对接口变量p，它被赋值为s，然后进行了nil检查。在Go语言中，当一个接口变量的动态值和动态类型都为nil时，
接口变量本身才会被认为是nil。但在这里，尽管s是一个nil指针，但它的类型是*Student1，而不是nil。因此，尽管s指向nil，
但它的动态类型不是nil，因此p也不是nil，所以输出了"p is not nil"。
*/
func main() {
	var s *Student1

	if s == nil {
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is not nil")
	}

	var p People1 = s

	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}
}
