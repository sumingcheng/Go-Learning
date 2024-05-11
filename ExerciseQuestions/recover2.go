package main

import (
	"fmt"
)

func riskyFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("在 riskyFunction 中恢复:", r)
		}
	}()
	fmt.Println("执行有风险的操作.")
	panic("发生了不好的事情")
	fmt.Println("此行不会执行.")
}

func main() {
	riskyFunction()
	fmt.Println("从 riskyFunction 正常返回。")
}
