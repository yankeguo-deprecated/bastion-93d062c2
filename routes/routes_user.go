package routes

import (
	"github.com/pagoda-tech/bastion/models"
	"github.com/pagoda-tech/macaron"
)

// UserUpdateForm 更新用户表单
type UserUpdateForm struct {
	Nickname string `json:"nickname"`
}

// UserUpdatePasswordForm 更新用户密码表单
type UserUpdatePasswordForm struct {
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
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

// UserUpdatePassword 修改密码
func UserUpdatePassword(ctx *macaron.Context, r APIRender, a Auth, db *models.DB, f UserUpdatePasswordForm) {
	if len(f.NewPassword) < models.UserPasswordMinLen {
		r.Fail(ParamsInvalid, "新密码过短")
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

	if !a.CurrentUser.IsAdmin || len(f.Password) > 0 {
		if !u.CheckPassword(f.Password) {
			r.Fail(ParamsInvalid, "旧密码不正确")
			return
		}
	}

	u.SetPassword(f.NewPassword)
	db.Model(u).Update("password_digest", u.PasswordDigest)

	r.Success()
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
