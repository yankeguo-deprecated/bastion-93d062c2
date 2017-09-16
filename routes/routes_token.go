package routes

import (
	"ireul.com/bastion/models"
	"ireul.com/bastion/utils"
	"ireul.com/web"
)

// TokenCreateForm 创建 Token 表单
type TokenCreateForm struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Desc     string `json:"desc"`
}

// TokenCreate 创建 Token 路由
func TokenCreate(ctx *web.Context, db *models.DB, f TokenCreateForm, r APIRender) {
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
	// audit
	db.Audit(u, "tokens.create", t)
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
func TokenDestroy(ctx *web.Context, r APIRender, a Auth, db *models.DB) {
	id := uint(ctx.ParamsInt(":id"))

	t := &models.Token{}

	if id == a.CurrentToken.ID {
		t = a.CurrentToken
	} else {
		db.First(t, id)
	}

	// check user belongs
	if db.NewRecord(t) || !a.CanAccessUser(t.UserID) {
		r.Fail(TokenNotFound, "没有找该令牌")
		return
	}

	// Delete
	db.Delete(t)
	// audit
	db.Audit(a.CurrentUser, "tokens.destroy", t)
	r.Success()
}

// TokenList 列出 Token
func TokenList(ctx *web.Context, db *models.DB, r APIRender, a Auth) {
	userID := uint(ctx.ParamsInt(":userid"))

	if !a.CanAccessUser(userID) {
		r.Fail(UserNotFound, "没有找到该用户")
		return
	}

	tokens := []models.Token{}

	db.Where("user_id = ?", userID).Find(&tokens)

	r.Success("tokens", tokens)
}
