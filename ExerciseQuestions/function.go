package main

import "fmt"

func hello() []string {
	return nil
}

func main() {
	// 函数可以像其他值一样被赋予变量。这意味着你可以把一个函数赋给一个变量，这个变量就成了该函数的一个引用或“函数指针”
	h := hello

	// 函数也可以作为一个值来传递
	hCall := hello()

	if h == nil {
		fmt.Println("h nil")
	} else {
		fmt.Println("h not nil")
	}

	if hCall == nil {
		fmt.Println("hCall nil")
	} else {
		fmt.Println("hCall not nil")
	}
}
