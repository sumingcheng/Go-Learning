package main

import "fmt"

type Person struct {
	age int
}

func main() {
	person := &Person{28}

	// 1.
	/*
		这个 defer 直接在调用时计算了 person.age 的值，
		因为它是作为参数传递给 fmt.Println 函数的。
		此时 person.age 是 28，所以这个值被存储起来用于稍后执行的 defer。
	*/
	defer fmt.Println("1", person.age) // 28

	// 2.
	/*
		这个 defer 通过一个匿名函数传递了 person 指针。
		由于它传递的是指针,由于它传递的是指针,由于它传递的是指针，当匿名函数在 defer 执行时，它访问的是 person 的当前 age 值，即最终的 age 值 29。
	*/
	defer func(p *Person) {
		fmt.Println("2", p.age) // 29
	}(person)

	// 3.
	/*
		这个 defer 也是一个匿名函数，但它不接受任何参数。
		它直接使用外围作用域中的 person 变量。因此，它同样会输出 person 在执行时刻的 age 值，也就是 29。
	*/
	defer func() {
		fmt.Println("3", person.age) // 29
	}()

	person.age = 29
}
