package main

import "fmt"

func main() {
	//str := "admin"
	//
	//length := len(str)
	//fmt.Println(length)

	//znStr := "中文"
	//length := len(znStr)
	//fmt.Println(length) // 6

	//str := "i am go"
	//aBytes := []byte(str)
	//fmt.Println(aBytes)
	//fmt.Println(len(aBytes))

	//str := "好好学习"
	//aBytes := []byte(str)
	//aRunes := []rune(str)
	//
	//fmt.Println(len(aBytes)) // 12
	//fmt.Println(len(aRunes)) // 4

	//fmt.Println("他说:\"穷是原罪\"") // 他说:"穷是原罪"
	//
	//fmt.Println(`他说:"穷是原罪"`) // 他说:"穷是原罪"

	//fmt.Println("Hello\aWorld")

}

func Str() {
	s := "Z中"
	brr := []byte(s)
	fmt.Println(len(s), len(brr)) // 4 4
	fmt.Println(brr)              // [90 228 184 173] 说明一个中文字符占3个字节 228 184 173 = 中
	fmt.Println(brr[2], s[2])     // 184 184  默认情况下把字符串当做 byte 切片来使用
}
