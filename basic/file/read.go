package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//	读取文件
	defer file.Close()
	content := make([]byte, 100)
	n, err := file.Read(content)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(content[0:n]))
	fmt.Println(string(content))
}
