package dao

import (
	"github.com/jmoiron/sqlx"
	"github.com/utc-six/errorx/errors"
	"github.com/utc-six/errorx/models"
)

// BatchCreateUsers 使用 sqlx 批量创建用户
func BatchCreateUsers(tx *sqlx.Tx, users []models.User) error {
	if tx == nil {
		return errors.New(400, "无效的数据库事务")
	}
	if len(users) == 0 {
		return errors.New(400, "用户列表不能为空")
	}

	// 构建批量插入的 SQL 语句
	query := `INSERT INTO users (id, name, status) VALUES (:id, :name, :status)`

	// 执行批量插入
	_, err := tx.NamedExec(query, users)
	if err != nil {
		return errors.New(500, "批量创建用户失败")
	}

	return nil
}
