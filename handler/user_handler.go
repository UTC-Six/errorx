package handler

import (
	"fmt"
	"net/http"
	"your_project/errors"
	"your_project/service"

	"github.com/gin-gonic/gin"
)

func GetUserHandler(c *gin.Context) {
	// 从请求中获取用户ID，例如通过 URL 参数
	idParam := c.Param("id")
	// 转换为整数
	var id int
	_, err := fmt.Sscanf(idParam, "%d", &id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的用户ID"})
		return
	}

	user, err := service.GetUser(id)
	if err != nil {
		// 使用 errors.As 解析自定义错误
		var customErr *errors.CustomError
		if errors.As(err, &customErr) {
			c.JSON(customErr.Code, gin.H{"error": customErr.Message})
			return
		}

		// 使用类型断言解析自定义错误
		if ce, ok := err.(*errors.CustomError); ok {
			c.JSON(ce.Code, gin.H{"error": ce.Message})
			return
		}

		// 处理未预料到的错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": "内部服务器错误"})
		return
	}

	c.JSON(http.StatusOK, user)
}
