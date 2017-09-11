package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

// AccessToken 代表一个用户的访问 Token
type AccessToken struct {
	gorm.Model
	// UserID 用户 ID
	UserID uint `gorm:"index"`
	// Token 随机 Token
	Secret string `gorm:"unique_index"`
	// Desc 描述
	Desc string `gorm:"type:text"`
	// UsedAt 最后一次使用时间
	UsedAt time.Time
}
