package main

import "fmt"

func modifyWithNamedReturn() (result int) {
	defer func() {
		result += 5 // 修改命名返回值
	}()

	return 10 // 初始设置返回值为 10，实际上是将 result 设置为 10
}

func main() {
	fmt.Println(modifyWithNamedReturn()) // 输出 15
}
