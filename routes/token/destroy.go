package token

import (
	"fmt"
	"github.com/pagoda-tech/bastion/models"
	"github.com/pagoda-tech/bastion/routes/middlewares"
	"github.com/pagoda-tech/macaron"
)

// Destroy 删除一个 Token
func Destroy(ctx *macaron.Context, r *middlewares.Render, a middlewares.Auth, db *models.DB) {
	// validate signed in
	if !a.SignedIn() {
		r.Fail(a.Code, a.Message)
		return
	}

	// extract current token if 'current'
	id := ctx.Params(":id")
	if id == "current" || id == fmt.Sprint(a.CurrentToken.ID) {
		db.Delete(a.CurrentToken)
		r.Success()
		return
	}

	// check user
	t := &models.Token{}
	db.First(t, id)
	if db.NewRecord(t) {
		r.Fail("not_found", "没有找到该令牌")
		return
	}

	// check user belongs
	if a.CurrentUser.ID == t.UserID || a.CurrentUser.IsAdmin {
		db.Delete(t)
		r.Success()
		return
	}

	r.Fail("no_permission", "没有权限")
}
