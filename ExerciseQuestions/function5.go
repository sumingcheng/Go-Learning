package main

import (
	"fmt"
	"os"
)

// Foo1 函数声明返回一个error接口。
func Foo1() error {
	// 声明一个*os.PathError类型的变量err，并初始化为nil。
	// 这是一个指向os.PathError的指针，os.PathError实现了error接口。
	var err *os.PathError = nil

	// 可能的业务逻辑，比如某种文件操作，这里省略。

	// 函数返回一个指向nil的*os.PathError指针作为error接口。
	return err
}

func main() {
	// 调用Foo1，返回值赋予err变量。err是error类型。
	err := Foo1()

	// 打印err的值，预期输出是<nil>，因为err是nil的*os.PathError类型。
	fmt.Println(err)

	// 检查err是否为nil。这里会输出false，这可能会令人意外。
	// 原因是err虽然是一个指向nil的*os.PathError，但它被返回为一个error类型。
	// 在Go中，一个接口值（这里是error）包含两部分：类型和值。
	// 这里的类型是*os.PathError，值是nil。因为类型信息存在，所以err不等于完全的nil。
	fmt.Println(err == nil)
}
