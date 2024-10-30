package models

import "gorm.io/gorm"

type User struct {
	ID     int    `json:"ID"`
	Name   string `json:"Name"`
	Status string `json:"Status"`
}

// 模拟数据库查询，根据 ID 获取用户
func FetchUserByID(id int) (*User, error) {
	// 这里使用简单的模拟数据
	if id == 1 {
		return &User{ID: 1, Name: "张三", Status: "active"}, nil
	} else if id == 2 {
		return &User{ID: 2, Name: "李四", Status: "inactive"}, nil
	} else {
		// 模拟通过 GORM 返回 ErrRecordNotFound
		return nil, gorm.ErrRecordNotFound
	}
}
