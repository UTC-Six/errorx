package main

import (
	"your_project/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 定义路由
	router.GET("/user/:id", handler.GetUserHandler)

	// 运行服务器
	router.Run(":8080")
}
