package main

func get2() []byte {
	raw := make([]byte, 10000)
	result := make([]byte, 3)
	copy(result, raw[:3])
	println(len(raw), cap(raw), &raw[0])
	return result // 返回一个新的切片，仅包含需要的数据
}

func main() {
	data := get2()
	println(len(data), cap(data), &data[0])
}
