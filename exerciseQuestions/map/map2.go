package main

import "fmt"

// Math 定义一个结构体 Math，包含两个整型字段 x 和 y
type Math struct {
	x, y int
}

// 定义一个全局变量 m，是一个 map，键为字符串类型，值为 Math 结构体
var m = map[string]Math{
	"foo": Math{2, 3}, // 在 map 中添加了一个键值对，键为 "foo"，值为 Math 结构体{2, 3}
}

func main() {
	// 从 map 中取出键为 "foo" 的值，并赋值给 tmp 变量
	tmp := m["foo"]
	// 修改 tmp 变量的 x 字段为 4
	tmp.x = 4
	// 将修改后的 tmp 变量重新存入 map 中，覆盖原来的值
	m["foo"] = tmp
	// 打印 map 中键为 "foo" 的值的 x 字段，预期结果为 4
	fmt.Println(m["foo"].x)
}
