package main

import (
	"fmt"
	"sync"
)

func main() {
	var waitGroup sync.WaitGroup

	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go printNum(i, &waitGroup)
	}

	waitGroup.Wait()
}

func printNum(i int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	fmt.Println(i)
}
