package main

import "fmt"

// Shape interface 定义了 Draw 方法
type Shape interface {
	Draw()
}

// Circle 实现了 Shape interface
type Circle struct{}

func (c *Circle) Draw() {
	fmt.Println("Drawing circle")
}

// Rectangle 实现了 Shape interface
type Rectangle struct{}

func (r *Rectangle) Draw() {
	fmt.Println("Drawing rectangle")
}

// ShapeDrawer struct 通过 Shape interface 进行操作，易于扩展
type ShapeDrawer struct {
	shapes []Shape
}

func (sd *ShapeDrawer) AddShape(shape Shape) {
	sd.shapes = append(sd.shapes, shape)
}

func (sd *ShapeDrawer) DrawAll() {
	for _, shape := range sd.shapes {
		shape.Draw()
	}
}

func main() {
	circle := &Circle{}
	rectangle := &Rectangle{}

	drawer := ShapeDrawer{}
	drawer.AddShape(circle)
	drawer.AddShape(rectangle)

	drawer.DrawAll()
}
