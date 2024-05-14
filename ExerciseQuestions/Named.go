package main

import "fmt"

type MyInt int

func (m MyInt) Add(other MyInt) MyInt {
	return m + other
}

func main() {
	var a MyInt = 5
	var b MyInt = 6
	fmt.Println("Sum:", a.Add(b)) // 输出: Sum: 11
}
