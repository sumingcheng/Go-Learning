package main

import "fmt"

type S struct {
	m string
}

// *S 是指针类型，返回的是指针
func foo() *S {
	// &S{"foo"} 是一个指向 S{"foo"} 的指针
	return &S{"foo"}
}

func main() {
	p := foo()
	fmt.Println(p.m)
}
