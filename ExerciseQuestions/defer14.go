package main

import "fmt"

func main() {
	// 调用 doubleScore 函数并打印结果
	fmt.Println(doubleScore(0))    //  0
	fmt.Println(doubleScore(20.0)) //  40.0
	fmt.Println(doubleScore(50.0)) //  50.0
}

// doubleScore 函数接受一个浮点数 source 并返回一个浮点数。
// 如果输入的 score 小于 1 或者大于等于 100，score 会被设置为 source 的原始值
func doubleScore(source float32) (score float32) {
	// defer 延迟调用，函数结束前执行
	defer func() {
		fmt.Println(score)
		if score < 1 || score >= 100 {
			score = source // 将 score 设置为输入值 source
		}
	}()

	return source * 2 // 返回 source 的两倍
}

/*
在 Go 语言中使用命名返回值和 defer 语句时
defer 中的代码会在函数实际返回前执行
但是在返回语句（如果有明确的返回值表达式，如 return source * 2）计算之后执行。
这意味着 defer 代码块可以访问并修改命名的返回值。
*/
