package main

import (
	"fmt"
	"os"
)

func main() {
	// 使用 fmt.Println 打印到标准输出，并添加换行
	fmt.Println("Hello, World!")

	// 使用 fmt.Printf 带有格式化的输出
	fmt.Printf("Number: %d, String: %s, Float: %.2f\n", 5, "example", 3.14159)

	// 将输出定向到标准错误
	fmt.Fprintf(os.Stderr, "This is an error message.\n")

	// 将格式化的字符串保存到变量
	s := fmt.Sprintf("Name: %s, Age: %d", "John", 30)
	fmt.Println(s)
}
