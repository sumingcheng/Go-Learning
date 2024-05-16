package main

import "fmt"

const (
	// 当 iota 为 0 时，1 << 0 等于 1。
	one = 1 << iota
	two
)

func main() {
	fmt.Println(one, two)
}
