package main

import "fmt"

var step1Failed = false
var step2Failed = true

func complexFunction() error {
	// 假设有多个步骤，每个步骤都可能出错
	if step1Failed {
		goto handleError
	}
	if step2Failed {
		goto handleError
	}
	return nil

handleError:
	// 处理错误
	fmt.Println("An error occurred")
	return fmt.Errorf("an error")
}
