package user

import (
	"fmt"
	"github.com/pagoda-tech/bastion/models"
	"github.com/pagoda-tech/bastion/routes/middlewares"
	"github.com/pagoda-tech/macaron"
)

// Show 显示一个用户
func Show(ctx *macaron.Context, r *middlewares.Render, a middlewares.Auth, db *models.DB) {
	// validate signed in
	if !a.SignedIn() {
		r.Fail(a.Code, a.Message)
		return
	}

	// extract current user if 'current'
	id := ctx.Params(":id")
	if id == "current" || id == fmt.Sprint(a.CurrentUser.ID) {
		r.Success("user", a.CurrentUser)
		return
	}

	// check IsAdmin
	if !a.CurrentUser.IsAdmin {
		r.Fail("not_found", "没有找到该用户")
		return
	}

	// find
	u := &models.User{}
	db.First(u, id)

	if !db.NewRecord(u) {
		r.Success("user", u)
	} else {
		r.Fail("not_found", "没有找到该用户")
	}
}
