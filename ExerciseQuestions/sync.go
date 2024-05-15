package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var balance int

func Deposit(amount int) {
	defer mutex.Unlock()
	mutex.Lock()
	balance += amount
}

func Balance() int {
	defer mutex.Unlock()
	mutex.Lock()
	return balance
}

func main() {
	Deposit(100)
	fmt.Println("Current Balance:", Balance())
}
