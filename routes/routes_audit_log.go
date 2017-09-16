package routes

import (
	"ireul.com/bastion/models"
	"ireul.com/web"
)

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
