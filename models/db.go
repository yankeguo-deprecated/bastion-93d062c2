package models

import (
	"time"

	"ireul.com/bastion/types"
	"ireul.com/bastion/utils"
	"ireul.com/orm"
	"ireul.com/web"
)

// DB 封装 orm.DB
type DB struct {
	*orm.DB
}

// NewDB 创建一个新的 DB 实例
func NewDB(cfg *utils.Config) (db *DB, err error) {
	var db0 *orm.DB
	if db0, err = orm.Open("mysql", cfg.Database.URL); err != nil {
		return nil, err
	}
	// create
	db = &DB{db0}
	// enable log if dev
	if cfg.Bastion.Env == web.DEV {
		db.LogMode(true)
	}
	return
}

// AutoMigrate 自动执行数据库更新
func (db *DB) AutoMigrate() {
	db.DB.AutoMigrate(AuditLog{}, Token{}, Server{}, SSHKey{}, User{})
}

// Touch 更新一个模型的 UsedAt 字段
func (db *DB) Touch(m interface{}) {
	db.DB.Model(m).UpdateColumn("UsedAt", time.Now())
}

// Audit create a new AuditLog
func (db *DB) Audit(source types.UserAuditable, action string, target types.Auditable) error {
	al := AuditLog{
		UserID: source.AuditableUserID(),
		Source: source.AuditableName(),
		Action: action,
		Target: target.AuditableName(),
	}
	return db.Create(&al).Error
}
