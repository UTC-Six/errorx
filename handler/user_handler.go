package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	cusErr "github.com/utc-six/errorx/errors"
	"github.com/utc-six/errorx/models"
	"github.com/utc-six/errorx/service"
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
		var customErr *cusErr.CustomError

		if !errors.As(err, &customErr) {
			c.JSON(500, gin.H{"error": "内部错误啦！！！！！"})
			return
		}
		c.JSON(customErr.Code, gin.H{"error": customErr.Message})

		// 使用类型断言解析自定义错误
		if ce, ok := err.(*cusErr.CustomError); ok {
			c.JSON(ce.Code, gin.H{"error": ce.Message})
			return
		}

		// 处理未预料到的错误
		c.JSON(http.StatusInternalServerError, gin.H{"error": "内部服务器错误"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func BatchCreateUsersHandler(c *gin.Context) {
	var users []models.User

	// 绑定 JSON 请求体到 users 切片
	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "无效的请求数据"})
		return
	}

	// 开始 sqlx 事务
	tx, err := service.SQLXDB.Beginx()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法开始数据库事务"})
		return
	}

	// 调用服务层的批量创建方法
	if err := service.BatchCreateUsers(tx, users); err != nil {
		tx.Rollback()
		var customErr *cusErr.CustomError
		if errors.As(err, &customErr) {
			c.JSON(customErr.Code, gin.H{"error": customErr.Message})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "内部服务器错误"})
		return
	}

	// 提交事务
	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法提交数据库事务"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "用户批量创建成功"})
}
