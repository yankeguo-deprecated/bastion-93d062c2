package routes

import (
	"github.com/pagoda-tech/bastion/models"
	"github.com/pagoda-tech/macaron"
)

// UserUpdateForm 更新用户表单
type UserUpdateForm struct {
	Nickname string `json:"nickname"`
}

// UserShow 显示一个用户
func UserShow(ctx *macaron.Context, r APIRender, a Auth, db *models.DB) {
	// extract current user if 'current'
	id := uint(ctx.ParamsInt(":id"))
	if id == a.CurrentUser.ID {
		r.Success("user", a.CurrentUser)
		return
	}

	// check IsAdmin
	if !a.CanAccessUser(id) {
		r.Fail(UserNotFound, "没有找到该用户")
		return
	}

	// find
	u := &models.User{}
	db.First(u, id)

	if db.NewRecord(u) {
		r.Fail(UserNotFound, "没有找到该用户")
		return
	}

	r.Success("user", u)
}

// UserUpdate 更新一个用户信息
func UserUpdate(ctx *macaron.Context, r APIRender, a Auth, db *models.DB, f UserUpdateForm) {
	if len(f.Nickname) >= 20 {
		r.Fail(ParamsInvalid, "昵称过长")
		return
	}

	id := uint(ctx.ParamsInt(":id"))

	if !a.CanAccessUser(id) {
		r.Fail(UserNotFound, "没有找到该用户")
		return
	}

	u := &models.User{}

	if id == a.CurrentUser.ID {
		*u = *a.CurrentUser
	} else {
		db.First(u, id)
	}

	if db.NewRecord(u) {
		r.Fail(UserNotFound, "没有找到该用户")
		return
	}

	if len(f.Nickname) > 0 {
		db.Model(u).Update("nickname", f.Nickname)
	}

	r.Success("user", u)
}
