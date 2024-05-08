package main

import "fmt"

func main() {
	x := 1

	fmt.Println(x)
	{
		fmt.Println(x)
		// 在 Go 语言中，大括号 {} 创建了一个新的作用域
		// 变量在这个作用域内被声明时就属于这个作用域的局部变量。
		// 这意味着，这些变量仅在这个特定的作用域内可见和可用。
		i, x := 2, 2
		fmt.Println(i, x)
	}
	fmt.Println(x)
}
