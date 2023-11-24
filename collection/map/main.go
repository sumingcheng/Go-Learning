package main

import "fmt"

func main() {
	// 声明并且初始化Map()
	// [key type]value type
	//myInfo := map[string]string{}
	//
	//myInfo["name"] = "zhangsan"
	//myInfo["age"] = "18"
	//
	//fmt.Println(myInfo) // map[age:18 name:zhangsan]
	//
	//// 直接初始化
	//myInfo2 := map[string]string{
	//	"name": "zhangsan",
	//	"age":  "18",
	//}
	//
	//fmt.Println(myInfo2) // map[age:18 name:zhangsan]
	//
	//// make初始化
	//myInfo3 := make(map[string]string, 3)
	//myInfo3["name"] = "zhangsan"
	//myInfo3["age"] = "999"
	//myInfo3["smc"] = "smc"
	//// 超过容量，会自动扩容
	//myInfo3["smc2"] = "smc2"
	//
	//fmt.Println(myInfo3) // map[age:999 name:zhangsan]
	//fmt.Println(len(myInfo3))
	//other()
	//test()
	//whetherItExists()
	//write()
	courseMap()
}

func other() {
	myInfo := map[string]string{
		"name": "zhangsan",
		"age":  "18",
	}
	// 修改
	myInfo["age"] = "999"

	fmt.Println(myInfo) // map[age:18]

	// 读取
	fmt.Println(myInfo["age"]) // 18

	// 删除
	delete(myInfo, "age")
	fmt.Println(myInfo) // map[name:zhangsan]
}

func test() {
	person := make(map[string]interface{})
	person["name"] = "John Doe"
	person["age"] = 30
	person["isEmployed"] = true

	fmt.Println(person)
}

func whetherItExists() {
	myInfo := map[string]string{
		"name": "zhangsan",
		"age":  "18",
	}
	myInfo["hobby"] = "table tennis"

	// 判断key是否存在
	// 第一个返回值为key对应的value，第二个返回值为key是否存在
	value, ok := myInfo["name"]
	fmt.Println(value, ok) // zhangsan true

	_, ok = myInfo["height"]
	fmt.Println(value, ok) // false
}

func write() {
	myInfo := map[string]string{
		"name": "zhangsan",
		"age":  "18",
	}
	myInfo["hobby"] = "table tennis"

	// for range针对map的枚举是无序的
	for key, value := range myInfo {
		fmt.Println(key, value) // 打印是无序的
	}
}

func courseMap() {
	courseMap := make(map[string]string, 3)
	courseMap["name"] = "golang"
	courseMap["site"] = "imooc"
	courseMap["price"] = "free"

	// 删除
	delete(courseMap, "price")
	fmt.Println(courseMap)
	// 删除不存在的key，不会报错
	delete(courseMap, "price")

	//if value, existed := courseMap["name"]; existed {
	//	fmt.Println(value)
	//} else {
	//	fmt.Println("key does not exist")
	//}

	//fmt.Println(courseMap)

	//value, existed := courseMap["name"]
	//fmt.Println(value, existed) // golang true

	// map 是无序的集合，不要依赖map的顺序
	//for key, value := range courseMap {
	//	fmt.Println(key, value)
	//}

	//	多类型的map
	//courseMap2 := make(map[string]interface{}, 3)
	//courseMap2["name"] = "golang"
	//courseMap2["site"] = true
	//courseMap2["price"] = 0
	//
	//fmt.Println(courseMap2)
}
