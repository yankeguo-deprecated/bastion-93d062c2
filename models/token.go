package models

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/pagoda-tech/gorm"
	"time"
)

// Token 代表一个用户的访问 Token
type Token struct {
	gorm.Model
	// UserID 用户 ID
	UserID uint `gorm:"index" json:"userId"`
	// Token 随机 Token
	Secret string `gorm:"unique_index" json:"-"`
	// Desc 描述
	Desc string `gorm:"type:text" json:"desc"`
	// UsedAt 最后一次使用时间
	UsedAt *time.Time `json:"usedAt"`
}

// GenerateSecret 创建一个新的 Secret
func (at *Token) GenerateSecret() (err error) {
	buf := make([]byte, 32)
	if _, err = rand.Read(buf); err != nil {
		return
	}
	at.Secret = hex.EncodeToString(buf)
	return
}
