package models

import (
	"fmt"
	"time"

	"ireul.com/orm"
)

// SSHKey 代表一个用户的 SSH 公钥
type SSHKey struct {
	orm.Model

	// Name SSHKey 的名字
	Name string `json:"name"`
	// UserID 用户ID
	UserID uint `orm:"index" json:"userId"`
	// Fingerprint 公钥 SHA256 指纹
	Fingerprint string `orm:"unique_index" json:"fingerprint"`
	// PublicKey SSH 公钥
	PublicKey string `orm:"type:text" json:"publicKey"`
	// UsedAt 最后一次使用时间
	UsedAt *time.Time `json:"usedAt"`
}

// AuditableName implements Auditable
func (t SSHKey) AuditableName() string {
	return fmt.Sprintf("SSHKey(%d)", t.ID)
}

// AuditableDetail implements Auditable
func (t SSHKey) AuditableDetail() string {
	return t.Fingerprint
}

// AuditableUserID implements UserAuditable
func (t SSHKey) AuditableUserID() uint {
	return t.UserID
}
