package main

import "fmt"

func main() {
	checkScore(90)
	checkNumber(2)
}

func checkNumber(num int) {
	switch v := num % 2; v {
	case 0:
		fmt.Println(num, "is even")
	default:
		fmt.Println(num, "is odd")
	}
}

func checkScore(score int) {
	switch {
	case score >= 90:
		fmt.Println("A grade")
	case score >= 80:
		fmt.Println("B grade")
	case score >= 70:
		fmt.Println("C grade")
	default:
		fmt.Println("D grade")
	}
}
