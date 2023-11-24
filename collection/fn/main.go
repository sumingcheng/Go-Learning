package main

import "fmt"

func main() {
	test()
}

func test() {
	a := 10
	_test := func() {
		fmt.Println(a)
	}
	_test()
}
