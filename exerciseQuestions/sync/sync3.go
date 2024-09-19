package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var value int

func Setup() {
	value = 42
	fmt.Println("Value set")
}

func DoSetup() {
	once.Do(Setup)
}

func main() {
	go DoSetup()
	go DoSetup()
	go DoSetup()

	// WaitGroup 等待 goroutines 完成
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		DoSetup()
	}()

	go func() {
		defer wg.Done()
		DoSetup()
	}()

	go func() {
		defer wg.Done()
		DoSetup()
	}()

	wg.Wait()
	fmt.Println("value:", value)
}
