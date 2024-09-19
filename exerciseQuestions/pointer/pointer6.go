package main

import "fmt"

type T2 struct {
	n int
}

func getT2() *T2 {
	return &T2{} // 返回一个 T 类型结构体的指针
}
func main() {

	t := getT2()
	t.n = 2
	fmt.Println(t.n)
}
