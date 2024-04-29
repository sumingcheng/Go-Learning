package main

import "fmt"

func GetValue(
	m map[int]string,
	id int,
) (string, bool) {
	if _, exist := m[id]; exist {
		return "exist", true
	} else {
		return "", false
	}
}

func main() {
	intMap := map[int]string{
		1: "A",
		2: "B",
		3: "C",
	}

	v, err := GetValue(intMap, 3)
	fmt.Println(v, err)
}
