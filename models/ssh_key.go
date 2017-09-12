package models

import (
	"github.com/pagoda-tech/gorm"
	"time"
)

// SSHKey 代表一个用户的 SSH 公钥
type SSHKey struct {
	gorm.Model

	// UserID 用户ID
	UserID uint `gorm:"index"`
	// Fingerprint 公钥 SHA256 指纹
	Fingerprint string `gorm:"unique_index"`
	// PublicKey SSH 公钥
	PublicKey string `gorm:"type:text"`
	// UsedAt 最后一次使用时间
	UsedAt time.Time
}
