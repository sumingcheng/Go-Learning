package calcuator

// Plus 需要公共调用的方法都需要大写
// 类型定义，变量类型在参数后面，后面跟返回值
func Plus(a int, b int) int {
	PrintOver()
	return a + b
}

func Subtract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a int, b int) int {
	return a / b
}
