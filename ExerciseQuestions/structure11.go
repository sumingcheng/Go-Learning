package main

import "fmt"

type Foo2 struct {
	val int // 定义Foo2结构体，包含一个整数字段val
}

func (f *Foo2) Inc(inc int) {
	f.val += inc
}

//func (f Foo2) Inc(inc int) {
//	f.val += inc
//}

func main() {
	var f Foo2         // 创建一个Foo2结构体的实例f
	f.Inc(100)         // 调用f的Inc方法，试图增加f.val的值100
	fmt.Println(f.val) // 打印f的val字段的值
}
