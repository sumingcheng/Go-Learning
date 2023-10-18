package main

import (
	"basic/calcuator"
	"fmt"
)

func main() {
	fmt.Println("计算器")
	var result int = calcuator.Plus(1, 2)
	fmt.Println(result)

	// println 和 fmt.Println 是有区别的
	var result2 int = calcuator.Multiply(1, 2)
	println(result2)

	// 类型推断会推断出 result3 的类型，所以不用定义
	result3 := calcuator.Divide(1, 2)
	fmt.Println(result3)

	calcuator.PrintOver()
}
