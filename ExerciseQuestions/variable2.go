package main

func main() {
	var val int

	println(&val)
	//fmt.Println(&val)

	f5(10000)

	println(&val)
	//fmt.Println(&val)
}

/*
动态栈：Go 的 goroutine 栈是动态的，可以根据需要增长和缩减。当你进行递归调用 f5 时，如果递归的深度足够深，可能导致栈的重新分配。

栈拷贝和移动：当栈空间不足以支持更深层次的函数调用时，Go 运行时会自动增加栈的大小。
这个过程可能涉及到将现有的栈内容复制到一个更大的空间。这种栈的重新分配可能会导致即使是 main 中定义的局部变量的地址发生变化。
*/

func f5(i int) {
	if i--; i == 0 {
		return
	}
	f5(i)
}
