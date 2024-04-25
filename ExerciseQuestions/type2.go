package main

import "fmt"

func GetValue1() int {
	return 1
}

func GetValue2() interface{} {
	return 1
}

func main() {
	i := GetValue2()

	// 在 Go 中，.type 断言只能用于接口类型的值，这是用来检测和提取接口值的实际类型。
	// 如果你尝试对非接口类型（如 int、string 等具体类型）使用 .type 断言，编译器会报错：
	// 所以 不能判断 GetValue1() 的类型
	switch i.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case interface{}:
		fmt.Println("interface")
	default:
		fmt.Println("unknown")
	}
}
