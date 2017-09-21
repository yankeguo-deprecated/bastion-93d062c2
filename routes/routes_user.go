package routes

import (
	"ireul.com/bastion/models"
	"ireul.com/web"
)

// UserCreateForm for for new user
type UserCreateForm struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// UserCreate create a user
func UserCreate(ctx *web.Context, r APIRender, f UserCreateForm, db *models.DB) {
	if !models.UserLoginRegexp.MatchString(f.Login) {
		r.Fail(ParamsInvalid, "登录名格式不正确")
		return
	}

	if len(f.Password) < models.UserPasswordMinLen {
		r.Fail(ParamsInvalid, "密码过短")
		return
	}

	u := &models.User{}

	db.Where("login = ?", f.Login).First(u)

	if !db.NewRecord(u) {
		r.Fail(ParamsInvalid, "登录名已存在")
		return
	}

	u = &models.User{
		Login:    f.Login,
		Nickname: f.Login,
	}
	if err := u.SetPassword(f.Password); err != nil {
		r.Fail(InternalError, err.Error())
		return
	}
	if err := u.GenerateSSHKey(); err != nil {
		r.Fail(InternalError, err.Error())
		return
	}

	if err := db.Create(u).Error; err != nil {
		r.Fail(InternalError, err.Error())
		return
	}

	r.Success("user", u)
}

// UserList list all users
func UserList(ctx *web.Context, r APIRender, db *models.DB) {
	us := []models.User{}
	db.Find(&us)
	r.Success("users", us)
}

// UserShow 显示一个用户
func UserShow(ctx *web.Context, r APIRender, a Auth, db *models.DB) {
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

// UserUpdatePasswordForm 更新用户密码表单
type UserUpdatePasswordForm struct {
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}

// UserUpdatePassword 修改密码
func UserUpdatePassword(ctx *web.Context, r APIRender, a Auth, db *models.DB, f UserUpdatePasswordForm) {
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

// UserUpdateForm 更新用户表单
type UserUpdateForm struct {
	Nickname string `json:"nickname"`
}

// UserUpdate 更新一个用户信息
func UserUpdate(ctx *web.Context, r APIRender, a Auth, db *models.DB, f UserUpdateForm) {
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

// UserUpdateAuthorityForm update user authority (is_admin, is_blocked)
type UserUpdateAuthorityForm struct {
	IsAdmin   bool `json:"isAdmin"`
	IsBlocked bool `json:"isBlocked"`
}

// UserUpdateAuthority blocks a user
func UserUpdateAuthority(ctx *web.Context, r APIRender, db *models.DB, f UserUpdateAuthorityForm, a Auth) {
	id := ctx.Params(":id")
	u := &models.User{}
	db.First(u, id)
	if db.NewRecord(u) {
		r.Fail(ParamsInvalid, "没有找到该用户")
		return
	}
	if u.IsAdmin != f.IsAdmin {
		db.Model(u).Update("is_admin", f.IsAdmin)
		if f.IsAdmin {
			db.Audit(a.CurrentUser, "users.admin_set", u)
		} else {
			db.Audit(a.CurrentUser, "users.admin_revoke", u)
		}
	}
	if u.IsBlocked != f.IsBlocked {
		db.Model(u).Update("is_blocked", f.IsBlocked)
		if f.IsBlocked {
			db.Audit(a.CurrentUser, "users.block", u)
		} else {
			db.Audit(a.CurrentUser, "users.unblock", u)
		}
	}
	r.Success()
}
