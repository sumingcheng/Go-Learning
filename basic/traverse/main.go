package main

import "fmt"

func main() {
	/*	// 遍历数组或者切片
		nums := []int{2, 4, 6, 8}
		for i, num := range nums {
			fmt.Printf("index: %d, value: %d\n", i, num)
		}

		// 不需要的参数可以用 _ 代替
		for _, num := range nums {
			fmt.Println("value:", num)
		}

		//  遍历 map
		m := map[string]int{"a": 1, "b": 2}
		for k, v := range m {
			fmt.Printf("key: %s, value: %d\n", k, v)
		}

		// 遍历字符串
		for i, c := range "go" {
			fmt.Print(c)
			fmt.Printf("  index: %d, char: %c\n", i, c)
		}

		// 遍历通道
		ch := make(chan int)
		go func() {
			ch <- 2
			ch <- 4
			close(ch)
		}()
		for num := range ch {
			fmt.Println("received:", num)
		}*/

	userName := `sumingcheng`
	passWord := `123992s6`

	for _, pw := range passWord {
		for _, us := range userName {
			if us == pw {
				//goto hasRepeatedChar
				fmt.Print("有重复字母")
				return
			}
		}
	}

	//hasRepeatedChar:
	//	fmt.Print("有重复字母")
}
