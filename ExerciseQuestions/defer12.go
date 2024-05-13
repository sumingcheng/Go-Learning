package main

func main() {
	// 调用 DeferTest1 函数，并打印返回值
	println(DeferTest1(1)) // 输出将是 4
	// 调用 DeferTest2 函数，并打印返回值
	println(DeferTest2(1)) // 输出将是 2
}

// DeferTest1 函数接受一个整型参数 i 并返回一个整型
func DeferTest1(i int) (r int) {
	r = i // 初始化返回值 r 为参数 i 的值
	// defer 语句，将延迟执行到函数返回前
	defer func() {
		r += 3 // 在返回前，r 的值增加 3
	}()

	return r // 返回 r 的值
}

// DeferTest2 函数接受一个整型参数 i 并返回一个整型
func DeferTest2(i int) (r int) {
	// defer 语句，将延迟执行到函数返回前
	defer func() {
		r += i // 在返回前，r 的值增加 i，此处为 1
	}()

	return 2 // 设置返回值为 2，此时 r 被设为 2
}
