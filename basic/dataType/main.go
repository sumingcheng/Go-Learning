package main

import (
	"fmt"
	"math"
)

func main() {
	// 整数类型
	fmt.Println("int8 range:", math.MinInt8, "to", math.MaxInt8)
	fmt.Println("uint8 range:", 0, "to", math.MaxUint8)
	fmt.Println("int16 range:", math.MinInt16, "to", math.MaxInt16)
	fmt.Println("uint16 range:", 0, "to", math.MaxUint16)
	fmt.Println("int32 range:", math.MinInt32, "to", math.MaxInt32)
	fmt.Println("uint32 range:", 0, "to", math.MaxUint32)
	fmt.Println("int64 range:", math.MinInt64, "to", math.MaxInt64)
	fmt.Println("uint64 range:", uint64(0), "to", uint64(math.MaxUint64)) // 显式转换为 uint64 类型

	// 浮点数类型
	fmt.Println("float32 range:", math.SmallestNonzeroFloat32, "to", math.MaxFloat32)
	fmt.Println("float64 range:", math.SmallestNonzeroFloat64, "to", math.MaxFloat64)

	// 复数类型无法直接获取最大最小值，但可以通过 float32 和 float64 的最大最小值来了解其范围

	// 特殊类型
	// 对于 int, uint 和 uintptr 类型，其范围依赖于具体的平台（32位或64位）
}
