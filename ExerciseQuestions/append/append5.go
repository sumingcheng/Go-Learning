package main

import "fmt"

func change(s ...int) {
	s = append(s, 3)
}

/*
知识点：可变参数函数、append()操作。Go语言中的变参数函数...，可以将slice传递给可变参数函数，不会创建新的切片。
第一次调用change()时，append()操作使传入的切片在底层数组上扩展了容量，原slice的底层数组并没有扩容；
第二次调用change()函数时，使用了切片的[i,j]来传递了一个新的切片，依赖于slice1，它的底层数组和slice1共用底层数组是同一个，不过slice1的长度是2，
所以在change()函数中对slice1底层数组的修改会影响到slice1。
*/
func main() {
	slice := make([]int, 5)
	slice[0] = 1
	slice[1] = 2

	change(slice...)
	fmt.Println(slice)

	change(slice[:2]...)
	fmt.Println(slice)
}
