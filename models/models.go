package models

import (
	"github.com/jinzhu/gorm"
)

// AutoMigrate 自动对所有模型进行 AutoMigrate
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(AccessToken{}, Server{}, SSHKey{}, User{})
}
