package models

import (
	"time"

	"ireul.com/bastion/types"
	"ireul.com/orm"
)

// DB 封装 orm.DB
type DB struct {
	*orm.DB
}

// NewDB 创建一个新的 DB 实例
func NewDB(env, url string) (db *DB, err error) {
	var d *orm.DB
	if d, err = orm.Open("mysql", url); err != nil {
		return nil, err
	}
	// create
	db = &DB{d}
	// enable log if dev
	if env == types.DEV || env == types.TEST {
		db.LogMode(true)
	}
	return
}

// AutoMigrate 自动执行数据库更新
func (db *DB) AutoMigrate() error {
	return db.DB.AutoMigrate(
		AuditLog{},
		Grant{},
		Token{},
		Server{},
		SSHKey{},
		User{},
	).Error
}

// Touch 更新一个模型的 UsedAt 字段
func (db *DB) Touch(m interface{}) {
	db.DB.Model(m).UpdateColumn("UsedAt", time.Now())
}
