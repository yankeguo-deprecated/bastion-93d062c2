package routes

import (
	"ireul.com/bastion/models"
	"ireul.com/web"
)

// AuditLogList list all AuditLog
func AuditLogList(ctx *web.Context, a Auth, r APIRender, db *models.DB) {
	total := 0
	offset := ctx.QueryInt("offset")
	db.Model(&models.AuditLog{}).Count(&total)
	data := []models.AuditLog{}
	db.Order("id DESC").Offset(offset).Limit(50).Find(&data)
	r.Success("auditLogs", data, "offset", offset, "total", total, "limit", 50)
}

// AuditLogsListByUser "/api/users/:id/audit_logs"
func AuditLogsListByUser(ctx *web.Context, a Auth, r APIRender, db *models.DB) {
	userID := uint(ctx.ParamsInt(":id"))

	if !a.CanAccessUser(userID) {
		r.Fail(UserNotFound, "没有找到该用户")
		return
	}

	auditLogs := []models.AuditLog{}
	db.Order("id DESC").Limit(30).Where("user_id = ?", userID).Find(&auditLogs)

	r.Success("auditLogs", auditLogs)
}
