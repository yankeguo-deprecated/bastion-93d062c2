package models

import (
	"ireul.com/bastion/types"
	"ireul.com/orm"
)

// AuditLog a auditable log, records important operations
type AuditLog struct {
	orm.Model
	User         User   `json:"user"`
	UserID       uint   `json:"userId" orm:"index"`
	Source       string `json:"source"`
	SourceDetail string `json:"sourceDetail"`
	Action       string `json:"action"`
	Target       string `json:"target"`
	TargetDetail string `json:"targetDetail"`
}

// Audit create a new AuditLog
func (db *DB) Audit(source types.UserAuditable, action string, target types.Auditable) (AuditLog, error) {
	al := AuditLog{
		UserID:       source.AuditableUserID(),
		Source:       source.AuditableName(),
		SourceDetail: source.AuditableDetail(),
		Action:       action,
		Target:       target.AuditableName(),
		TargetDetail: target.AuditableDetail(),
	}
	return al, db.Create(&al).Error
}
