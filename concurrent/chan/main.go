package main

import (
	"fmt"
	"sync"
	"time"
)

var s []int
var wg sync.WaitGroup
var c = make(chan int)

func addStep() {
	defer wg.Done()
	step := 1

	for {
		c <- step
		step++
		time.Sleep(1 * time.Second)
	}
}

func addEl() {
	defer wg.Done()
	for {
		select {
		case step := <-c:
			s = append(s, step)
			fmt.Println("add", s)
			time.Sleep(2 * time.Second)
		}
	}

}
func main() {
	wg.Add(2)
	go addStep()
	go addEl()
	wg.Wait()
}
