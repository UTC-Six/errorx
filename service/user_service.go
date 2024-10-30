package service

import (
	"your_project/errors"
	"your_project/models"

	"gorm.io/gorm"
)

// 获取用户信息
func GetUser(id int) (*models.User, error) {
	user, err := models.FetchUserByID(id)
	if err != nil {
		// 检查是否为 GORM 的记录未找到错误
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		// 其他数据库错误
		return nil, errors.ErrInternal
	}

	// 业务逻辑中可能产生的自定义错误
	if user.Status == "inactive" {
		return nil, errors.New(403, "用户已被禁用")
	}

	return user, nil
}
