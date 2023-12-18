package main

import "fmt"

func main() {
	// any 本质上是 interface{} 的别名
	student := map[string]any{
		"name": "张三",
		"age":  18,
	}

	fmt.Println(student)

}

type PlusT interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32
}

func plus[T PlusT](a, b T) T {
	return a + b
}
