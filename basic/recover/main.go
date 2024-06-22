package main

import "fmt"

func riskyOperation() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("r:", r)
		}
	}()

	panic("出了问题")
}

func main() {
	riskyOperation()
	fmt.Println("程序继续...")
}
