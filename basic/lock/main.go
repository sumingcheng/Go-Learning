package main

import (
	"fmt"
	"sync"
)

var (
	a    = 0
	lock = sync.Mutex{}
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			lock.Lock()
			a++
			lock.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			lock.Lock()
			a++
			lock.Unlock()
		}
	}()

	wg.Wait()
	fmt.Println("a", a)
}
