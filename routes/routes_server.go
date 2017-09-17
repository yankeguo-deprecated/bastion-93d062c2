package routes

import (
	"ireul.com/bastion/models"
	"ireul.com/web"
)

// ServerList list all servers
func ServerList(ctx *web.Context, db *models.DB, r APIRender, a Auth) {
	list := []models.Server{}
	db.Find(&list)
	r.Success("servers", list)
}
