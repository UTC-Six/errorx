package main

import (
	"github.com/utc-six/errorx/handler"
	"github.com/utc-six/errorx/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	service.InitDB()

	router := gin.Default()

	// 定义路由
	router.GET("/user/:id", handler.GetUserHandler)
	router.POST("/users/batch", handler.BatchCreateUsersHandler) // 新增批量创建用户的路由

	// 运行服务器
	router.Run(":8080")
}
