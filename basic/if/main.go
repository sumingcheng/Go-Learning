package main

import "fmt"

func main() {
	checkScore(90)
	checkScore(80)
	checkScore(70)
}

//func checkNumber(num int) {
//	switch v := num % 2; v {
//	case 0:
//		fmt.Println(num, "is even")
//	default:
//		fmt.Println(num, "is odd")
//	}
//}

//func checkScore(score int) {
//	switch {
//	case score >= 90:
//		fmt.Println("A grade")
//		fallthrough
//	case score >= 80:
//		fmt.Println("B grade")
//		fallthrough
//	case score >= 70:
//		fmt.Println("C grade")
//		fallthrough
//	default:
//		fmt.Println("D grade")
//	}
//}

func checkScore(score int) {
	switch {
	case score >= 90, score >= 80, score >= 70:
		fmt.Println("A grade")
	default:
		fmt.Println("D grade")
	}
}
