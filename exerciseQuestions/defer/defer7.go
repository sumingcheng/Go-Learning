package main

import "fmt"

var a1 bool = true

func main() {
	// defer 写在前面会先注册，但是不会立即执行
	defer func() {
		fmt.Println("1")
	}()

	if a1 == true {
		fmt.Println("2")
		return
	}

	defer func() {
		fmt.Println("3")
	}()
}
