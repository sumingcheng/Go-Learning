package main

import "fmt"

func main() {
	/*
		2: 2 指定索引 2 的值为 2。
		3 紧跟在 2: 2 后面，因此这里默认索引是 3，所以索引 3 的值为 3。
		0: 1 指定索引 0 的值为 1。
	*/
	var x = []int{2: 2, 3, 0: 1}

	fmt.Println(x) // [1 0 2 3]
}
