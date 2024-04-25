package main

import "fmt"

func main() {
	// 代码中定义了两种匿名结构体。
	// 匿名结构体没有显式的名称，而是直接在声明时定义其结构。
	// 这种方式通常用于快速创建简单的对象，不需要在多个地方复用结构体类型。
	sn1 := struct {
		age  int
		name string
	}{
		age:  11,
		name: "qq",
	}

	sn2 := struct {
		age  int
		name string
	}{
		age:  11,
		name: "qq",
	}

	// Go 语言允许直接比较两个结构体变量，前提是结构体中的所有字段都是可比较的类型。在的代码中，sn1 和 sn2 被比较：
	// 因为结构体中的所有字段（int 和 string）都是可比较的，
	// 所以这两个结构体实例可以直接使用 == 进行比较。如果所有对应字段的值都相等，则两个结构体相等。
	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{
		age: 11,
		m:   map[string]string{"a": "1"},
	}

	sm2 := struct {
		age int
		m   map[string]string
	}{
		age: 11,
		m:   map[string]string{"a": "1"},
	}

	/*
	  包含不可比较字段的结构体不能使用 == 进行比较，否则会引发编译错误。
	  切片 (Slice)：虽然可以遍历切片中的元素或者检查它们的内容，但切片本身在 Go 中是不可比较的。你不能使用 == 或 != 运算符直接比较两个切片。
	  映射 (Map)：与切片类似，映射也不能直接被比较。映射的比较操作会引发编译时错误，因为映射的内部实现是动态的，且可以在运行时改变其大小和内容。
	  函数 (Function)：函数在 Go 中也是不可比较的。你不能比较两个函数是否相等，不论它们的签名或行为是否相同。
	  通道 (Channel)：虽然通道可以在某些情况下比较（如是否指向同一个通道对象），但在结构体比较中通常不建议包含通道类型，因为它们的运行时状态可能影响比较结果。
	*/
	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}

}
