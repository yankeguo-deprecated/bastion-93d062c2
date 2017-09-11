package accesstoken

import (
	"github.com/pagoda-tech/bastion/models"
	"github.com/pagoda-tech/macaron"
)

// CreateForm 创建 AccessToken 表单
type CreateForm struct {
	login    string `form:"login" binding:"required"`
	password string `form:"login" binding:"required"`
}

// CreateAction 创建 AccessToken 路由
func CreateAction(ctx *macaron.Context, db *models.DB, form *CreateForm) {
}
