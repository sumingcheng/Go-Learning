package main

import (
	"fmt"
	"time"
)

func main() {
	var m = [...]int{1, 2, 3}

	// 当go func()运行时，它实际上使用的是循环结束时i和v的值，而不是goroutine启动时的值。
	// 因此，所有的goroutine最终打印的都是循环结束时的i和v的值，即2和3。

	// 因为 goroutine 是异步的 它们在实际执行时可能会在循环结束后。
	// 因此，当goroutine最终执行时，它们会访问i和v变量的当前值，而在这个时候，i和v的值已经是循环结束时的值了。

	// 这种现象是由闭包引用的特性所导致的。闭包函数在创建时会引用其周围作用域的变量，
	// 但它们实际上不会复制这些变量的值，而是在函数执行时动态地获取它们的值。
	// 因此，当 goroutine 最终执行时，它们实际上访问的是i和v的当前值，而这个值已经在循环结束后被更新为最终的值。
	for i, v := range m {
		go func() {
			fmt.Println(i, v)
		}()
	}

	time.Sleep(3 * time.Second)

	// 为了解决这个问题，可以通过将迭代变量的值复制给局部变量，确保每个goroutine使用的是不同的副本，而不是共享同一个变量。
	// 这样就可以避免goroutine访问到循环结束后的值。

	for i, v := range m {
		go func(i, v int) {
			fmt.Println(i, v)
		}(i, v)
	}

	time.Sleep(3 * time.Second)

}
