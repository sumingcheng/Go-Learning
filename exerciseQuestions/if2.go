package main

import "fmt"

func main() {
	err := "外部错误" // 这是最外层的err变量

	if true {
		err := "内部错误"    // 新声明的err变量，这将遮蔽外部的err
		fmt.Println(err) // 输出 "内部错误"
	}

	fmt.Println(err) // 输出 "外部错误"
}
