package main

import "fmt"

type MyInt int

// 这个泛型函数接受任何底层类型为 int 的类型
func myFunc[T ~int](a T) T {
	return a + 1
}

func main() {
	var a int = 5
	var b MyInt = 6

	fmt.Println(myFunc(a)) // 输出 6
	fmt.Println(myFunc(b)) // 输出 7
}
