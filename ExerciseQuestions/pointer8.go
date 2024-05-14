package main

import "fmt"

// 定义P为*int的别名
type P *int

// 定义Q为*int的别名
type Q *int

func main() {
	// 创建一个int类型的指针p，初始值为指向的int的零值（0）
	var p P = new(int)
	// 通过指针p将值增加到8
	*p += 8

	// Go的常见做法是通过解引用指针来操作它所指向的数据的值，而不是直接操作指针本身的地址
	// 将p的值赋给一个*int类型的变量x
	var x *int = p

	// 将x的值赋给一个Q类型的变量q
	var q Q = x

	// 通过指针q将值增加1
	*q++

	// 打印通过p和q指针看到的值，因为它们指向同一个内存地址，所以都会输出9
	fmt.Println(*p, *q) // 输出: 9 9
}
