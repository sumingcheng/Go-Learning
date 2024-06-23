package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

type User struct {
	Id      int
	Keyword string `gorm:"column:keywords"`
	City    string
}

func read(
	client *gorm.DB,
	city string,
) *User {
	var user User
	err := client.Select("id,city").Where("city=?", city).Limit(1).Take(&user).Error
	checkError(err)
	return &user
}

func initializeDB(client *gorm.DB) {
	// 自动迁移，创建表
	if err := client.AutoMigrate(&User{}); err != nil {
		fmt.Println("Failed to migrate database: ", err)
		os.Exit(1)
	}

	// 模拟数据
	users := []User{
		{Keyword: "developer", City: "北京"},
		{Keyword: "engineer", City: "上海"},
		{Keyword: "designer", City: "广州"},
	}

	// 插入数据到数据库
	for _, user := range users {
		result := client.Create(&user)
		if err := result.Error; err != nil {
			fmt.Println("Failed to insert data: ", err)
			continue
		}
		fmt.Printf("Inserted user %v\n", user)
	}
}

func main() {
	//	数据库连接
	dataSourceName := "root:123114@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True"
	client, err := gorm.Open(mysql.Open(dataSourceName), nil)
	checkError(err)

	// 初始化数据库并插入数据
	//initializeDB(client)
	// 查找
	user := read(client, "北京")
	if user != nil {
		fmt.Printf("%+v\n", *user)
	} else {
		fmt.Println("没有找到")
	}

	// 插入单条数据
	//client.Create(User{
	//	Id:      100,
	//	Keyword: "golang",
	//	City:    "杭州",
	//})

	client.Model(User{}).Where("id=?", 100).Update("city", "朝鲜")
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
