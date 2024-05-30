package main

type Rectangle struct {
	Width, Height int
}

// 使用指针接收者来修改矩形的尺寸
func (r *Rectangle) Resize(newWidth, newHeight int) {
	r.Width = newWidth
	r.Height = newHeight
}

// 使用值接收者来计算矩形的面积，不修改原始数据
func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func main() {
	r := Rectangle{Width: 10, Height: 5}
	r.Resize(100, 50)

	area := r.Area()
	println(area)
}
