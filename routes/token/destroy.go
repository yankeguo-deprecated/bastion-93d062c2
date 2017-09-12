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

	//TODO: add deletion of other tokens
	r.Fail("not_supported", "当前请求暂时不被支持")
}
