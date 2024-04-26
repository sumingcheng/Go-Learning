package main

import "fmt"

func main() {
	var a []int // 不分配内存

	b := []int{} // 分配内存

	fmt.Println(a, b)
}
