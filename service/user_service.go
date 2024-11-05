package service

import (
	"github.com/utc-six/errorx/dao"
	"github.com/utc-six/errorx/errors"
	"github.com/utc-six/errorx/models"

	"github.com/jmoiron/sqlx" // 新增
	// "gorm.io/gorm" // 如果不再使用 GORM，可移除此行
)

var SQLXDB *sqlx.DB // 新增

func InitDB() {
	var err error
	// 初始化 sqlx.DB，使用与 GORM 相同的数据库
	SQLXDB, err = sqlx.Connect("sqlite3", "test.db")
	if err != nil {
		panic("无法连接到数据库")
	}

	// 自动创建表，如果需要
	schema := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		name TEXT,
		status TEXT
	);`
	SQLXDB.MustExec(schema)
}

// BatchCreateUsers 使用 sqlx 进行批量创建用户
func BatchCreateUsers(tx *sqlx.Tx, users []models.User) error {
	// 调用 DAO 层的批量创建方法
	if err := dao.BatchCreateUsers(tx, users); err != nil {
		return err
	}
	return nil
}

// 获取用户信息
func GetUser(id int) (*models.User, error) {
	// 使用 sqlx 查询用户
	var user models.User
	err := SQLXDB.Get(&user, "SELECT id, name, status FROM users WHERE id = ?", id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" { // 根据实际情况调整
			return nil, errors.ErrNotFound
		}
		return nil, errors.ErrInternal
	}

	if user.Status == "inactive" {
		return nil, errors.New(403, "用户已被禁用")
	}

	return &user, nil
}
