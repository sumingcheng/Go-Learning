package main

import "fmt"

var a bool = true

func main() {
	// defer 写在前面会先注册，但是不会立即执行
	defer func() {
		fmt.Println("1")
	}()

	if a == true {
		fmt.Println("2")
		return
	}

	defer func() {
		fmt.Println("3")
	}()
}
