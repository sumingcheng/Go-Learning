package main

import "fmt"

// 函数返回一个函数切片
func test3() []func() {
	var funs []func() // 定义一个函数类型的切片

	for i := 0; i < 2; i++ { // 循环两次
		// 将闭包函数 append 到切片中
		funs = append(funs, func() {
			fmt.Println(&i, i) // 输出 i 的地址和值
		})
	}
	return funs // 返回函数切片
}

func main() {
	funs := test3()          // 调用 test 函数，获取函数切片
	fmt.Println(funs)        // [0xb32e80 0xb32e80]
	for _, f := range funs { // 遍历切片中的每个函数
		f() // 调用函数
	}
}
