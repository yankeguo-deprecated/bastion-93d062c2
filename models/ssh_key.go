package models

import (
	"github.com/pagoda-tech/gorm"
	"time"
)

// SSHKey 代表一个用户的 SSH 公钥
type SSHKey struct {
	gorm.Model

	// Name SSHKey 的名字
	Name string `json:"name"`
	// UserID 用户ID
	UserID uint `gorm:"index" json:"userId"`
	// Fingerprint 公钥 SHA256 指纹
	Fingerprint string `gorm:"unique_index" json:"fingerprint"`
	// PublicKey SSH 公钥
	PublicKey string `gorm:"type:text" json:"publicKey"`
	// UsedAt 最后一次使用时间
	UsedAt *time.Time `json:"usedAt"`
}
