package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// 用户相关的路由分组
	userGroup := router.Group("/user")
	{
		userGroup.POST("/", addUser)         // 添加用户
		userGroup.DELETE("/:id", deleteUser) // 删除用户
		userGroup.PUT("/:id", updateUser)    // 更新用户
	}

	// 监听并在 0.0.0.0:8080 上启动服务
	router.Run(":8080")
}

// 添加用户的处理函数
func addUser(c *gin.Context) {
	c.JSON(200, gin.H{"message": "用户添加成功"})
}

// 删除用户的处理函数
func deleteUser(c *gin.Context) {
	// 从 URL 中获取用户 ID
	userID := c.Param("id")
	c.JSON(200, gin.H{"message": "成功删除用户", "userID": userID})
}

// 更新用户的处理函数
func updateUser(c *gin.Context) {
	userID := c.Param("id")
	c.JSON(200, gin.H{"message": "用户更新成功", "userID": userID})
}
