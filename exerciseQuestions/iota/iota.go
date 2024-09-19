package main

import "fmt"

const (
	x = iota
	_
	y
	z = "zz"
	k
	p = iota
)

func main() {
	// iota 在每个新的 const 声明块中被重置为 0。在一个 const 声明块内，每新增一行常量声明，iota 的值会自动增加 1。
	// iota 的值是从 0 开始的，所以 x = 0, y = 1, z = "zz", k = "zz", p = 2
	// iota 可以被显式地修改其模式。例如，在本例中，z 被赋予了一个字符串值 "zz"。此时，iota 的增长不会停止，
	// 但是因为 z 的值是显式指定的，所以 k（紧随 z 声明的常量）也会采用同样的值 "zz"，而不是 iota 的当前值。
	// 尽管 z 和 k 的值被显式设置为 "zz"，iota 本身的值继续按顺序增加。当声明 p 时，
	// iota 的值已经增加到 5（从 0 开始，跳过 1，然后是 y 为 2，z 和 k 分别是 3 和 4）。因此 p 的值是 5。
	fmt.Println(x, y, z, k, p)
}
