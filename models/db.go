package models

import (
	"ireul.com/bastion/utils"
	"ireul.com/orm"
	"ireul.com/web"
	"time"
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
	db.DB.AutoMigrate(Token{}, Server{}, SSHKey{}, User{})
}

// Touch 更新一个模型的 UsedAt 字段
func (db *DB) Touch(m interface{}) {
	db.DB.Model(m).UpdateColumn("UsedAt", time.Now())
}
