package main

// test 函数返回两个闭包函数
func test(x int) (func(), func()) {
	// 第一个闭包函数：打印 x 的当前值，然后对 x 加 10
	return func() {
			println(x) // 打印 x 的初始值
			x += 10    // 修改 x 的值
		},
		func() { // 第二个闭包函数：打印 x 的当前值
			println(x) // 因为闭包捕获的是 x 的引用，所以这里打印的是 x 被第一个函数修改后的值
		}
}

// main 函数是程序的入口
func main() {
	a, b := test(100) // 调用 test 函数，传入初始值 100
	a()               // 调用第一个闭包，预期打印 100，然后将 x 增加到 110
	b()               // 调用第二个闭包，预期打印 110，显示第一个闭包对 x 的修改效果
}
