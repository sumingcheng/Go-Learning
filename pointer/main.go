package main

import (
	"fmt"
)

func main() {
	a := 1
	test(&a) // int 数组的指针类型
	fmt.Println(a)
}

func test(a *int) {
	b := 100
	fmt.Printf("%p\n", a)
	a = &b

	fmt.Printf("%p\n", a)
}
