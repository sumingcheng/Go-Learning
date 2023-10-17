package main

import "fmt"

func main() {
	//a := 1
	//b := true
	//c := 'c'
	//d := "ddd"
	//
	//f := a + b + c + d
	//fmt.Println(f)

	// 先声明变量 a 赋值的时候给 a 进行类型推断
	var a = 1
	// 先声明再进行类型推断
	b := 2
	fmt.Println(a, b)

	//如果类型先写着前面，就确定了类型，确定后就无法推断了
	//int a = 1 这种情况就没办法做类型推断，因为类型已经确定了
}
