package main

import (
	"fmt"
)

type ConfigOne struct {
	Daemon string
}

// String 方法返回 ConfigOne 实例的字符串表示形式
func (c *ConfigOne) String() string {
	return fmt.Sprintf("Config Daemon: %s", c.Daemon)
}

func main() {
	// 创建 ConfigOne 的实例并初始化 Daemon 字段
	c := &ConfigOne{Daemon: "ExampleDaemon"}
	fmt.Println(c.String())
}
