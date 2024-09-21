package main

import "fmt"

// ButtonPresser interface 表示按按钮的行为
type ButtonPresser interface {
	PressButton()
}

// Lamp struct 为一个灯具
type Lamp struct{}

func (l *Lamp) PressButton() {
	fmt.Println("Lamp is turned on")
}

// Switch struct 需要一个 ButtonPresser 来控制
type Switch struct {
	device ButtonPresser
}

func (s *Switch) PressButton() {
	s.device.PressButton()
}

func main() {
	lamp := &Lamp{}
	switchDevice := &Switch{device: lamp}
	switchDevice.PressButton()
}
