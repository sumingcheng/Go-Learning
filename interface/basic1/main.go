package main

import "fmt"

// Speaker 定义一个接口
type Speaker interface {
	Speak() string
}

// Dog 定义一个实现了 Speaker 接口的类型
type Dog struct{}

func (d Dog) Speak() string {
	return "汪汪"
}

// Cat 定义另一个实现了 Speaker 接口的类型
type Cat struct{}

func (c Cat) Speak() string {
	return "喵喵"
}

func performSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	dog := Dog{}
	cat := Cat{}

	performSpeak(dog) // 输出: 汪汪
	performSpeak(cat) // 输出: 喵喵
}
