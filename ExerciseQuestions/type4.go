package main

import "fmt"

type Animal interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string {
	return "Woof"
}

type Cat struct{}

func (c Cat) Speak() string {
	return "Meow"
}

func identifyAnimal(a Animal) {
	if cat, ok := a.(Cat); ok {
		fmt.Println("This is a Cat, saying:", cat.Speak())
	} else if dog, ok := a.(Dog); ok {
		fmt.Println("This is a Dog, saying:", dog.Speak())
	} else {
		fmt.Println("Unknown animal")
	}
}

func main() {
	identifyAnimal(Dog{})
}
