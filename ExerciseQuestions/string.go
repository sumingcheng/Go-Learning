package main

import "fmt"

func main() {
	// Go 语言中字符串是不可变的
	//str := "hello"
	//str[0] = "x"
	//fmt.Println(str)

	// 这段代码先将字符串转换为字节切片，修改切片中的元素，然后将切片再转换回字符串
	str := "hello"
	b := []byte(str)
	fmt.Println(b)
	b[0] = 'x'
	newStr := string(b)
	fmt.Println(newStr)

}
