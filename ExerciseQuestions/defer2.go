package main

import "fmt"

func num(i int) {
	fmt.Println(i)

}

func main() {
	i := 5
	/*
	  变量初始化：变量 i 被初始化为 5。
	  defer 语句：defer num(i) 被调用。
	  关键点是此时 i 的值为 5，因此 num 函数的参数 i 被固定为 5。在 Go 中，defer 语句会立即对其参数进行求值，
	  所以即使后面 i 的值改变了，被 defer 的函数调用所用的参数值仍然是在 defer 语句执行时确定的 5。
	  变量修改：代码继续执行，将 i 修改为 15。
	  结束 main 函数：main 函数执行完毕。由于存在 defer 语句，此时会调用 num(5)。
	  因为之前在 defer 语句中参数已经固定为 5，所以即使 i 现在是 15，输出依然是 5。
	*/
	defer num(i)
	i = i + 10
}
