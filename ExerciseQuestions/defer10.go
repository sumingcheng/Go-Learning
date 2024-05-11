package main

func F(n int) func() int {
	return func() int {
		n++
		return n
	}
}

/*
先打印i的值（6）。
接着执行defer fmt.Println(f())，此时f()自增n至7，打印7。
最后执行defer func() { println(f()) }()，此时f()自增n至8，打印8。
*/

//重点
/*
调试中的输出：如果你的调试依赖于即时和准确的输出顺序，考虑使用 println 或 log.Println，后者尤其适合复杂的应用程序或多线程环境。
格式化和灵活性：如果需要格式化输出或输出复杂的数据结构，fmt.Println 是一个好选择。为确保缓冲问题不影响调试，可以在关键位置手动刷新缓冲区。
*/
func main() {
	f := F(5)

	defer func() {
		println(f())
	}()

	defer println(f())

	i := f()
	println(i)
}
