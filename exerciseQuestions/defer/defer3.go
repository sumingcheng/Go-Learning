package main

import "fmt"

func increaseA() int {
	var i int // 初始化为0

	defer func() {
		i++
	}()
	// return i 发生后，defer 语句执行，i 变为 1，但这不影响返回值，因为返回值已经被确定（已经从 i 复制到返回位置）。
	return i
}

func increaseB() (r int) {
	defer func() {
		r++
	}()

	// return r 语句指示将命名返回值 r 的当前值（初始为 0）准备返回。
	// 在返回之前执行 defer，将 r 递增到 1。
	// 返回值（r）此时为 1，因此返回给调用者的是修改后的 r。
	return r
}

func main() {
	/*
		命名返回值和非命名返回值在与 defer 一起使用时的重要区别。
		命名返回值在整个函数中都是有效的，并且 defer 可以修改最终的返回值，
		而非命名返回值则在返回时已经确定，defer 的修改不会影响该值
	*/

	fmt.Println(increaseA())
	fmt.Println(increaseB())
}
