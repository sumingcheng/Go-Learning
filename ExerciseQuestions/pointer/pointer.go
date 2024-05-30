package main

import "fmt"

// 定义一个函数 incr，接收一个指向 int 的指针，返回一个 int 类型的值
func incr(p *int) int {
	// 通过解引用指针 p 来增加 p 指向的值
	// 解引用可以理解为“取消引用操作”，即通过引用（指针）取消对地址的直接使用，转而使用该地址中存储的数据。
	*p++
	// 返回修改后的值
	return *p
}

func main() {
	// 初始化一个 int 类型的变量 p，并设置为 1
	p := 1
	// 调用 incr 函数，传递 p 的地址
	incr(&p)
	// 输出 p 的当前值，此时 p 已被 incr 函数修改为 2
	fmt.Println(p) // 输出 2
}
