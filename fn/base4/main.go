package main

func main() {
	var slice ComputeSlice = []int{1, 2}
	println(slice.Sum())
	println(slice.Sub())
	println(slice.Mul())
	println(slice.Div())
}

type ComputeSlice []int

func (cs ComputeSlice) Sum() int {
	return cs[0] + cs[1]
}

func (cs ComputeSlice) Sub() int {
	return cs[0] - cs[1]
}

func (cs ComputeSlice) Mul() int {
	return cs[0] * cs[1]
}

func (cs ComputeSlice) Div() int {
	return cs[0] / cs[1]
}
