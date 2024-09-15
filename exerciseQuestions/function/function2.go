package main

import "fmt"

// 变参函数（Variadic Function）指的是一种可以接受可变数量参数的函数。
func num18(num ...int) {
	// 修改切片的第一个元素
	num[0] = 18
}

func main() {
	i := []int{5, 6, 7}

	// 传递切片的引用
	num18(i...)

	// 这显示了切片的引用传递特性和函数中对切片内容的修改是如何影响到函数外部的原始切片的。
	fmt.Println(i[0])
}
