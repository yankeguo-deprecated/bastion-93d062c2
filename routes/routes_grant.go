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
	u := &models.User{}
	for _, g := range gv {
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
	} else {
		return nil
	}
}

// GrantCreate create/update a grant
func GrantCreate(ctx *web.Context, r APIRender, db *models.DB, f GrantCreateForm) {
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
	} else {
		db.Model(g).Update(map[string]interface{}{"CanSudo": f.CanSudo, "ExpiresAt": f.ExpiresAt()})
	}
	r.Success("grant", g)
}

// GrantDestroy update a grant
func GrantDestroy(ctx *web.Context, r APIRender, db *models.DB) {
	id := ctx.ParamsInt(":id")
	db.Unscoped().Where("id = ?", id).Delete(&models.Grant{})
	r.Success()
}
