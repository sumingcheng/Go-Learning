package main

import (
	"fmt"
)

func main() {
	var original int = 128 // 定义一个整数变量

	// 显示原始值
	fmt.Println("Original:", original)

	// 转换为不同的整数类型
	fmt.Println("As int8:", int8(original))
	fmt.Println("As int16:", int16(original))
	fmt.Println("As int32:", int32(original))
	fmt.Println("As int64:", int64(original))
	fmt.Println("As uint:", uint(original))
	fmt.Println("As uint8:", uint8(original))
	fmt.Println("As uint16:", uint16(original))
	fmt.Println("As uint32:", uint32(original))
	fmt.Println("As uint64:", uint64(original))
	fmt.Println("As uintptr:", uintptr(original))

	// 转换为浮点类型
	fmt.Println("As float32:", float32(original))
	fmt.Println("As float64:", float64(original))

	// 转换为复数类型
	fmt.Println("As complex64:", complex64(complex(float32(original), 0)))
	fmt.Println("As complex128:", complex128(complex(float64(original), 0)))
}
