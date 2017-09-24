package routes

import (
	"time"

	"ireul.com/bastion/models"
	"ireul.com/web"
)

// GrantList list all grants
func GrantList(ctx *web.Context, r APIRender, db *models.DB) {
	gs := []models.Grant{}
	db.Where("tag = ?", ctx.Params(":tag")).Find(&gs)
	gv := models.ConvertGrantResolved(gs)
	for _, g := range gv {
		u := &models.User{}
		db.First(u, g.UserID)
		g.UserLogin = u.Login
	}
	r.Success("grants", gv)
}

// GrantCreateForm from for create grant
type GrantCreateForm struct {
	UserID    uint   `json:"userId"`
	UserLogin string `json:"userLogin"`
	Tag       string `json:"tag"`
	CanSudo   bool   `json:"canSudo"`
	ExpiresIn int64  `json:"expiresIn"`
}

// ExpiresAt convert ExpiresAt to *time.Time
func (f GrantCreateForm) ExpiresAt() *time.Time {
	// convert to time.Time
	if f.ExpiresIn != 0 {
		var t time.Time
		t = time.Now().Add(time.Second * time.Duration(f.ExpiresIn))
		return &t
	}
	return nil
}

// GrantCreate create/update a grant
func GrantCreate(ctx *web.Context, r APIRender, db *models.DB, f GrantCreateForm, a Auth) {
	if f.UserID == 0 {
		if len(f.UserLogin) == 0 {
			r.Fail(ParamsInvalid, "没有指定用户")
			return
		}
		u := &models.User{}
		db.Where("login = ?", f.UserLogin).First(u)
		if db.NewRecord(u) {
			r.Fail(ParamsInvalid, "没有找到用户")
			return
		}
		f.UserID = u.ID
	}
	g := &models.Grant{}
	db.Where("user_id = ? AND tag = ?", f.UserID, f.Tag).First(g)
	if db.NewRecord(g) {
		*g = models.Grant{
			UserID:    f.UserID,
			Tag:       f.Tag,
			CanSudo:   f.CanSudo,
			ExpiresAt: f.ExpiresAt(),
		}
		db.Create(g)
		db.Audit(a.CurrentUser, "grants.create", g)
	} else {
		db.Model(g).Update(map[string]interface{}{"CanSudo": f.CanSudo, "ExpiresAt": f.ExpiresAt()})
		db.Audit(a.CurrentUser, "grants.update", g)
	}
	r.Success("grant", g)
}

// GrantDestroy update a grant
func GrantDestroy(ctx *web.Context, r APIRender, db *models.DB, a Auth) {
	id := ctx.ParamsInt(":id")
	g := &models.Grant{}
	db.First(g, id)
	if db.NewRecord(g) {
		r.Fail(ParamsInvalid, "没有找到记录")
		return
	}
	db.Audit(a.CurrentUser, "grants.destroy", g)
	db.Unscoped().Where("id = ?", id).Delete(&models.Grant{})
	r.Success()
}
