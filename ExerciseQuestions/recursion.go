package main

import "fmt"

var f9 = func(i int) {
	fmt.Print("x")
}

func main() {
	//var f9 func(int)

	f9 = func(i int) {
		fmt.Print(i)

		if i > 0 { // 如果 i 大于 0，则递归调用 f 函数
			f9(i - 1) // 说白了就是 main 内部的 f9 还没有被定义完毕，所以调用的是外面的f9
		}
	}

	f9(10) // 调用重新定义的 f 函数，并传入初始值 10
}
