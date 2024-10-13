package main

import (
	"fmt"
)

// 定义结构体
type Person struct {
	Name string
	Age  int
}

// 函数类型
func greet(name string) string {
	return "Hello, " + name
}

// 接口类型
type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

// 自定义数据结构：队列（Queue）
type Queue struct {
	elements []int
}

// 入队操作
func (q *Queue) Enqueue(value int) {
	q.elements = append(q.elements, value)
}

// 出队操作
func (q *Queue) Dequeue() int {
	if len(q.elements) == 0 {
		fmt.Println("Queue is empty")
		return -1 // 或者返回错误
	}
	value := q.elements[0]
	q.elements = q.elements[1:]
	return value
}

// 获取队列的长度
func (q *Queue) Size() int {
	return len(q.elements)
}

func main() {
	// ===========================
	// 基本数据类型
	// ===========================

	// 1. 整型（Integers）
	var a int = 10
	var b int64 = 1000000000
	c := uint8(255) // 类型推断
	fmt.Println("整型示例:")
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
	fmt.Println()

	// 2. 浮点型（Floats）
	var pi float64 = 3.14159
	e := 2.71828 // 类型推断为 float64
	fmt.Println("浮点型示例:")
	fmt.Println("Pi:", pi)
	fmt.Println("e:", e)
	fmt.Println()

	// 3. 布尔型（Booleans）
	var isActive bool = true
	var isEnabled = false // 类型推断
	fmt.Println("布尔型示例:")
	fmt.Println("Is Active:", isActive)
	fmt.Println("Is Enabled:", isEnabled)
	fmt.Println()

	// 4. 字符串（Strings）
	var greeting string = "Hello, Go!"
	name := "Alice"
	message := greeting + " " + name
	fmt.Println("字符串示例:")
	fmt.Println(message)
	fmt.Println()

	// ===========================
	// 复合数据类型
	// ===========================

	// 1. 数组（Arrays）
	var numbers [3]int = [3]int{1, 2, 3}
	fruits := [3]string{"Apple", "Banana", "Cherry"}
	fmt.Println("数组示例:")
	fmt.Println("Numbers:", numbers)
	fmt.Println("Fruits:", fruits)
	fmt.Println()

	// 2. 切片（Slices）
	var primes []int = []int{2, 3, 5, 7}
	primes = append(primes, 11, 13)
	fmt.Println("切片示例:")
	fmt.Println("Primes:", primes)
	fmt.Println()

	// 3. 映射（Maps）
	capitals := map[string]string{
		"France": "Paris",
		"Japan":  "Tokyo",
		"India":  "New Delhi",
	}
	capitals["Germany"] = "Berlin"
	fmt.Println("映射示例:")
	fmt.Println("Capitals:", capitals)
	fmt.Println()

	// 4. 结构体（Structs）
	person := Person{Name: "Bob", Age: 25}
	fmt.Println("结构体示例:")
	fmt.Println("Person:", person)
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println()

	// 5. 指针（Pointers）
	var d int = 42
	var ptr *int = &d
	fmt.Println("指针示例:")
	fmt.Println("Value of d:", d)
	fmt.Println("Address of d:", ptr)
	fmt.Println("Value at ptr:", *ptr)
	fmt.Println()

	// 6. 函数（Functions）
	var sayHello func(string) string = greet
	msg := sayHello("Charlie")
	fmt.Println("函数示例:")
	fmt.Println(msg)
	fmt.Println()

	// 7. 接口（Interfaces）
	var animal Speaker
	animal = Dog{}
	fmt.Println("接口示例 - Dog:")
	fmt.Println("Dog says:", animal.Speak())

	animal = Cat{}
	fmt.Println("接口示例 - Cat:")
	fmt.Println("Cat says:", animal.Speak())
	fmt.Println()

	// ===========================
	// 自定义数据结构
	// ===========================

	// 队列（Queue）示例
	queue := &Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println("自定义数据结构 - 队列示例:")
	fmt.Println("Queue size:", queue.Size())
	fmt.Println("Dequeue:", queue.Dequeue())
	fmt.Println("Queue size:", queue.Size())
}
