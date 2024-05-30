package main

import "fmt"

var o = fmt.Println

/*
在第一次循环中，由于通道 c 是空的，select 可以立即执行 case c <- 1: 分支，发送值 1 到通道并打印 3。
在第二次循环中，通道 c 中有一个元素，因此可以执行 case <-c: 分支，从通道中接收值并打印 2，随后通道被设置为 nil。
在第三次循环中，由于通道 c 已经被设置为 nil，无法执行发送或接收操作，所以 select 会执行 default 分支，打印 1。
*/
func main() {
	c := make(chan int, 1) // 创建一个容量为1的缓冲通道

	for range [3]struct{}{} { // 循环三次
		select { // select 语句处理多个通道操作
		default:
			o(1) // 如果所有case都不能立即执行，打印1
		case <-c: // 尝试从通道c接收值
			o(2)
			c = nil // 将通道c设为nil，之后的操作将无法再使用c进行通道操作
		case c <- 1: // 尝试向通道c发送值1
			o(3)
		}
	}
}
