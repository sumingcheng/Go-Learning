package main

func main() {
	doSth(&Bird{
		legs: 2,
	})
}

type Duck interface {
	walk()
	shout()
}

type Bird struct {
	legs int
}

func (b *Bird) walk() {
	println("walk")
}

func (b *Bird) shout() {
	println("shout")
}

func doSth(animal Duck) {
	animal.walk()
	animal.shout()
}
