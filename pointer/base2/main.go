package main

import "fmt"

func main() {
	x := 1
	y := 2
	x, y = y, x
	//exchange(&x, &y)
	fmt.Println(x, y)
	//fmt.Println(&x, &y)
}

//func exchange(a, b *int) {
//	*a, *b = *b, *a
//}
