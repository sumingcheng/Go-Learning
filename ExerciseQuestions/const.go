package main

import "fmt"

/*
常量在Go中是不可变的值，这意味着它们一旦在编译时被定义，它们的值就不能更改。
不允许获取常量的地址可以避免产生修改常量值的可能，这符合常量“不变”的本质。
这样的设计简化了常量的使用和理解，开发者无需担心常量值会在程序运行中被改变。
*/
const i = 100

var j = 123

func main() {
	//fmt.Println(&i, i)
	fmt.Println(&j, j)
}
