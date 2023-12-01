package main

import "fmt"

func main() {
	type Student struct {
		name   string
		age    int
		course string
	}

	// 策略模式 通过map的key来区分不同的策略
	students := map[int]Student{
		1: Student{"John", 18, "Computer Science"},
		2: Student{"Peter", 20, "Computer Science"},
		3: Student{"Mary", 22, "Computer Science"},
	}

	fmt.Println(students)

	for _, student := range students {
		println(student.name, student.age, student.course)
	}
}
