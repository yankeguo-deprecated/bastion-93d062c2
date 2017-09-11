package models

import (
	"crypto/rand"
	"encoding/hex"
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

// GenerateSecret 创建一个新的 Secret
func (at *AccessToken) GenerateSecret() (err error) {
	buf := make([]byte, 16)
	if _, err = rand.Read(buf); err != nil {
		return
	}
	at.Secret = hex.EncodeToString(buf)
	return
}
