package main

import "fmt"

func main() {
	m1 := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
		"Q":               64,
	}

	// 注意 map 是无序的
	for s, _ := range m1 {
		fmt.Println(s)
	}
}
