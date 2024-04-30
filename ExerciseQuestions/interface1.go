package main

import "fmt"

func Foo(x interface{}) {
	fmt.Println(x)
	// 接口类型的值由两部分组成：类型信息和具体的值。即使接口包含的具体值是nil，只要类型信息存在，接口自身就不被视为nil。
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

func main() {
	var x *interface{} = nil
	Foo(x)
}
