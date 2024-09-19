package main

import (
	"fmt"
	"sort"
)

type Employee struct {
	Department string
	Years      int
}

func main() {
	employees := []Employee{
		{"Engineering", 3},
		{"Human Resources", 5},
		{"Engineering", 1},
		{"Human Resources", 2},
		{"Engineering", 5},
	}

	// 按部门排序
	sort.SliceStable(employees, func(i, j int) bool {
		return employees[i].Department < employees[j].Department
	})

	// 同一部门内按年限排序
	sort.SliceStable(employees, func(i, j int) bool {
		return employees[i].Years < employees[j].Years
	})

	fmt.Println("Sorted Employees:", employees)
}
