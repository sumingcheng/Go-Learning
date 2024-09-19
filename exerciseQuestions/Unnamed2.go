package main

import "fmt"

type MyInt1 int

func printInt(i int) {
	fmt.Println(i)
}

func main() {
	var x MyInt1 = 7
	// 直接传递将引起编译错误，需要显式转换：
	printInt(int(x))
}
