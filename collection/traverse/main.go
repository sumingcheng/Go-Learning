package main

import "fmt"

func main() {
	//studentArray := []string{"John", "Paul", "George", "Ringo"}
	//
	//for i := 0; i < len(studentArray); i++ {
	//	println(i, studentArray[i])
	//}
	//
	//for index, value := range studentArray {
	//	println(index, value)
	//}
	//other()
	copyArr()
}

func other() {
	studentArray := []string{"John", "Paul", "George", "Ringo"}
	studentOther := []string{"1", "2", "3", "4"}

	studentArray = append(studentArray, studentOther...)

	for index, value := range studentArray {
		println(index, value)
	}
}

func copyArr() {
	studentArray := []string{"John", "Paul", "George", "Ringo"}
	newArr := make([]string, len(studentArray))
	res := copy(newArr, studentArray)
	fmt.Print(res, newArr)
}
