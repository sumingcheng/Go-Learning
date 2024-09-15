package main

import "fmt"

func main() {
	if1()
	if2()
}

func if1() {
	if a := 1; false {
	} else if b := 2; false {
	} else if c := 3; false {
	} else {
		fmt.Println(a, b, c)
	}
}

// if1 函数等价于 if2 函数
func if2() {
	{
		a := 1
		if false {

		} else {
			{
				b := 2
				if false {

				} else {

					{
						c := 3
						if false {

						} else {
							fmt.Println(a, b, c)
						}
					}
				}
			}
		}
	}
}
