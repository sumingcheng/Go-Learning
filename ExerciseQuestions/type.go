package main

import "fmt"

// 类型定义 定义的类型是全新的类型，和原有的类型不是同一种类型
type myType1 int
type myType2 int

func main() {
	var i int = 10
	// 因为尽管它们都基于 int，但它们被视为独立的类型。
	// 直接赋值会报错，需要强制转换
	var i1 myType1 = myType1(i)
	var i2 myType2 = myType2(i)
	fmt.Println(i1, i2)
}
