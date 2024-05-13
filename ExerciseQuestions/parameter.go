package main

import (
	"fmt"
	"testing"
)

// 定义可变参数函数 hello1
func hello1(num ...int) {
	num[0] = 18 // 将可变参数的第一个元素修改为 18
}

func Test13(t *testing.T) {
	i := []int{5, 6, 7} // 定义一个整数切片
	hello1(i...)        // 将切片展开传递给 hello1 函数
	fmt.Println(i[0])   // 打印切片的第一个元素
}

func main() {
	t := &testing.T{} // 创建一个 testing.T 的实例
	Test13(t)         // 调用测试函数 Test13
}
