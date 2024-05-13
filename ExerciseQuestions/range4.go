package main

import "fmt"

// 定义一个结构体T1，其中包含一个整型字段n
type T1 struct {
	n int
}

func main() {
	// 初始化一个长度为2的数组ts，数组中的元素类型为T1
	ts := [2]T1{}

	// 使用for循环遍历数组ts，i为索引，t为数组元素的副本（注意：这里t是T1的副本）
	for i, t := range ts {
		switch i {
		case 0:
			// 当索引i为0时，设置当前元素的副本t的n字段为3
			t.n = 3
			// 直接在数组ts中设置索引为1的元素的n字段为9
			//ts[1].n = 9
			(&ts)[1].n = 9 // 通过数组的地址（指针），访问第二个元素，并修改其n字段为9
		case 1:
			// 当索引i为1时，打印当前元素的副本t的n字段和一个空格
			// 注意：此处打印的是副本t的n字段值，由于t是副本，其值未被更改，仍为0
			fmt.Println(t.n, " ")
		}
	}
	// 打印整个数组ts，此时数组的状态为 [T1{n=0} T1{n=9}]
	// 由于t为副本修改，ts[0].n的值没有被修改为3
	fmt.Println(ts)
}
