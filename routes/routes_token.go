package routes

import (
	"github.com/pagoda-tech/bastion/models"
	"github.com/pagoda-tech/bastion/utils"
	"github.com/pagoda-tech/macaron"
	"fmt"
)

// TokenCreateForm 创建 Token 表单
type TokenCreateForm struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Desc     string `json:"desc"`
}

// TokenCreate 创建 Token 路由
func TokenCreate(ctx *macaron.Context, db *models.DB, f TokenCreateForm, r APIRender) {
	u := &models.User{}
	// find
	db.Where("login = ?", f.Login).First(u)
	if u.ID == 0 {
		r.Fail(ParamsInvalid, "用户名或密码错误")
		return
	}
	// check password
	if !u.CheckPassword(f.Password) {
		r.Fail(ParamsInvalid, "用户名或密码错误")
		return
	}
	// create token
	t := &models.Token{
		UserID: u.ID,
		Desc:   f.Desc,
	}
	t.GenerateSecret()
	db.Create(t)
	// return data
	r.Success(func(m utils.Map) {
		m.Set("token", utils.NewMap(
			"id",
			t.ID,
			"secret",
			t.Secret,
		))
	})
}

// TokenDestroy 删除一个 Token
func TokenDestroy(ctx *macaron.Context, r APIRender, a Auth, db *models.DB) {
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
		r.Fail(TokenNotFound, "没有找该令牌")
		return
	}

	// check user belongs
	if a.CurrentUser.ID == t.UserID || a.CurrentUser.IsAdmin {
		db.Delete(t)
		r.Success()
		return
	}

	r.Fail(TokenNotFound, "没有找该令牌")
}

// TokenList 列出 Token
func TokenList(ctx *macaron.Context, db *models.DB, r APIRender, a Auth) {
	userID := ctx.Params(":userId")

	if userID == "current" {
		userID = fmt.Sprint(a.CurrentUser.ID)
	} else {
		if userID != fmt.Sprint(a.CurrentUser.ID) && !a.CurrentUser.IsAdmin {
			r.Fail(UserNotFound, "没有找到该用户")
			return
		}
	}

	tokens := []models.Token{}

	db.Where("user_id = ?", userID).Find(&tokens)

	r.Success("tokens", tokens)
}
