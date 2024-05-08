package main

import (
	"encoding/json"
	"fmt"
)

type People2 struct {
	Name string `json:"name"` // 字段首字母大写，以便可以被 json 包访问
}

func main() {
	js := `{"name": "seekload"}` // 正确定义 JSON 字符串
	var p People2
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Printf("%+v\n", p) // 格式化输出结构体内容
}
