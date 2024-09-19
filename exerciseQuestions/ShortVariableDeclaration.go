package main

import "fmt"

// Go 强制开发者在全局作用域中更明确地声明变量，而不是使用简短变量声明
// Go 语言的目标：保持语言的简单和清晰，同时提高程序的可读性和可维护性。
var (
	size    = 1024
	maxSize = size * 2
)

func main() {
	fmt.Println(size, maxSize)
	shortStatement()
}

// Short Variable Declaration
// 简短变量声明，只能在函数内部使用
func shortStatement() {
	size := 1024
	maxSize := size * 2

	fmt.Println(size, maxSize)
}
