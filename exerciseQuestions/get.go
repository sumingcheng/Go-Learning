package main

import "fmt"

func get() []byte {
	raw := make([]byte, 10000)               // 创建一个长度和容量均为10000的字节切片
	fmt.Println(len(raw), cap(raw), &raw[0]) // 打印切片的长度、容量和第一个元素的地址
	return raw[:3]                           // 原始切片的容量 (cap(raw)) 减去起始位置的偏移量，即 10000 - 0 = 10000。
}

func main() {
	data := get()                               // 调用get函数，获取返回的切片
	fmt.Println(len(data), cap(data), &data[0]) // 打印返回切片的长度、容量和第一个元素的地址
}

// 10000 10000 XXX
// 3 10000 XXX
