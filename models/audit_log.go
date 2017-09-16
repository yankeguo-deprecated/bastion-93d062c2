package models

import (
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
