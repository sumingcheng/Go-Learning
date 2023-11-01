package main

import "fmt"

func main() {
	//var arr = [3]string{"a", "b", "c"}
	//arr[3] = "d"
	//fmt.Println(arr)

	//var arr1 []string
	//var arr2 []string
	//
	//arr2 = append(arr1, "1", "2", "3")
	//
	//fmt.Println(arr1, arr2) // [] [1 2 3]

	//intercept6()

}

// 截取 [start, end) 之间的元素
func intercept() {
	arr := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	newArr := arr[1:6]
	fmt.Println(newArr)
}

func intercept2() {
	arr := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	newArr := arr[1:len(arr)]
	fmt.Println(newArr) // [b c d e f g h]
}

// 取到最后
func intercept3() {
	arr := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	newArr := arr[1:]
	fmt.Println(newArr) // [b c d e f g h]
}

func intercept4() {
	arr := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	newArr := arr[:4]
	fmt.Println(newArr) // [a b c d]
}

func intercept5() {
	arr := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	newArr := arr[:]
	fmt.Println(newArr) // [a b c d e f g h]
}

func intercept6() {
	arr := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	newArr := arr[:5]
	arr[0] = "1"
	fmt.Println(newArr) // [1 b c d e]
	fmt.Println(arr)    // [1 b c d e f g h]
}
