package main

import "fmt"

// Orange 结构体
type Orange struct {
	Quantity int // 定义一个整数字段 Quantity 表示橙子的数量
}

// Increase 方法用于增加橙子的数量
// 参数 n 是要增加的数量
func (o *Orange) Increase(n int) {
	o.Quantity += n // 将 Quantity 增加 n
}

// Decrease 方法用于减少橙子的数量
// 参数 n 是要减少的数量
func (o *Orange) Decrease(n int) {
	o.Quantity -= n // 将 Quantity 减少 n
}

// String 方法返回橙子的数量字符串表示
func (o *Orange) String() string {
	// 使用 fmt.Sprintf 格式化字符串输出橙子的数量
	return fmt.Sprintf("%d", o.Quantity)
}

// main 函数是程序的入口
func main() {
	var orange Orange    // 声明一个 Orange 类型的变量 orange
	orange.Increase(10)  // 调用 Increase 方法增加 10 个橙子
	orange.Decrease(5)   // 调用 Decrease 方法减少 5 个橙子
	fmt.Println(&orange) // string 方法定义在了 *Orange 指针上，没有定义在值上，想使用的话要么通过指针访问，要么定义在值上
}
