package main

import "fmt"

func test2(i int) (ret int) {
	ret = i * 2
	if ret > 10 {
		ret = 10
		return // 提前返回
	}
	return // 默认返回
}

func main() {
	result := test2(10)
	fmt.Println(result)
}
