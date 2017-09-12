package user

import (
	"github.com/pagoda-tech/bastion/routes/middlewares"
	"github.com/pagoda-tech/macaron"
)

// Show 显示一个用户
func Show(ctx *macaron.Context, r *middlewares.Render, a middlewares.Auth) {
	if !a.SignedIn() {
		r.Fail(a.Code, a.Message)
		return
	}
	id := ctx.Params(":id")
	if id == "current" {
		r.Success("user", a.CurrentUser)
		return
	}
}
