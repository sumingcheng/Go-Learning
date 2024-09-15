package main

import (
	"fmt"
	"sort"
)

type S3 struct {
	v int
}

func main() {
	s := []S3{{1}, {3}, {5}, {2}}
	fmt.Printf("Original: %v\n", s)

	// 排序 s
	sort.Slice(s, func(i, j int) bool {
		return s[i].v < s[j].v
	})

	fmt.Printf("Sorted: %v\n", s)
}
