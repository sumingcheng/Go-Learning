package main

import (
	"fmt"
	"sync"
)

// 定义UserAges结构体，包含一个map来存储用户名和年龄，以及一个互斥锁
type UserAges struct {
	ages map[string]int
	sync.Mutex
}

// Add方法用于添加或更新用户的年龄
func (ua *UserAges) Add(
	name string,
	age int,
) {
	ua.Lock()           // 在修改map之前加锁
	defer ua.Unlock()   // 方法结束时解锁
	ua.ages[name] = age // 在map中设置用户的年龄
}

// GetName方法用于获取用户的年龄，如果用户不存在，则返回-1
func (ua *UserAges) GetName(name string) int {
	if age, ok := ua.ages[name]; ok {
		return age // 如果找到用户，返回年龄
	}
	return -1 // 如果找不到用户，返回-1
}

// 主函数
func main() {
	count := 1000          // 设置goroutine的数量
	gw := sync.WaitGroup{} // 创建一个WaitGroup来等待所有goroutine完成
	gw.Add(count * 3)      // 因为将要启动3000个goroutine

	u := UserAges{ages: map[string]int{}} // 创建UserAges实例，初始化ages map
	add := func(i int) {                  // 定义一个匿名函数，用于添加用户
		u.Add(fmt.Sprintf("user_%d", i), i) // 添加用户
		gw.Done()                           // 每完成一个goroutine，减少WaitGroup的计数器
	}

	for i := 0; i < count; i++ {
		go add(i) // 启动三次add goroutine，每个i值启动三个
		go add(i)
		go add(i)
	}

	for i := 0; i < count; i++ {
		go func(i int) { // 启动一个匿名goroutine，用于获取用户年龄
			defer gw.Done()                      // 函数结束时减少WaitGroup的计数器
			u.GetName(fmt.Sprintf("user_%d", i)) // 获取用户年龄
		}(i)
	}

	gw.Wait()           // 等待所有goroutine完成
	fmt.Println("Done") // 打印完成消息
}
