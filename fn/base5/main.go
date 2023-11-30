package main

import "fmt"

func main() {
	res1 := plus[int](1, 2)
	fmt.Println(res1)
	res2 := plus("hello", "world")
	fmt.Println(res2)
	res3 := plus(1.1, 2.1)
	fmt.Println(res3)
}

/*
Go 语言中的泛型必须有类型约束
Go 语言的泛型需要在函数名后面加上[T]，表示这是一个泛型函数
*/
func plus[T int | string | float64](a T, b T) T {
	return a + b
}
