package main

import "fmt"

func main() {
	var f1 = func() {}
	var f2 = func() {}
	fmt.Println(f1, f2)
	//if f1 != f2 {
	//	fmt.Println("f1 and f2 are equal")
	//}
}
