package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := s1[1:]
	s2[1] = 4
	fmt.Println(s1) // [1 2 4]
	fmt.Println(s2) // [2 4]

	s2 = append(s2, 5, 6, 7)
	fmt.Println(s2) // [2 4 5 6 7]
	//  S1 和 S2 共享底层数组，所以对 S2 的修改会影响 S1
	fmt.Println(s1) // [1 2 4]
}
