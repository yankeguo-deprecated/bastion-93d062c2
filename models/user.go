package models

import (
	"github.com/jinzhu/gorm"
)

// User 代表一个用户
type User struct {
	gorm.Model

	// Login 登录名
	Login string `gorm:"unique_index"`
	// Nickname 昵称
	Nickname string
	// PasswordDigest bcrypt 加密后的密码
	PasswordDigest string `gorm:"type:text"`
	// IsBlocked 用户是否被禁用
	IsBlocked bool `gorm:"not null"`
	// IsAdmin 用户是否是管理员
	IsAdmin bool `gorm:"not null"`
}
