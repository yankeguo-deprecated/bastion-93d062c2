package routes

import (
	"time"

	"ireul.com/bastion/models"
	"ireul.com/web"
)

// GrantList list all grants
func GrantList(ctx *web.Context, r APIRender, db *models.DB) {
	gs := []models.Grant{}
	db.Find(&gs)
	r.Success("grants", models.ConvertGrantResolved(gs))
}

// GrantCreateForm from for create grant
type GrantCreateForm struct {
	UserID    uint   `json:"userId"`
	Tag       string `json:"tag"`
	CanSudo   bool   `json:"canSudo"`
	ExpiresIn int64  `json:"expiresIn"`
}

// ExpiresAt convert ExpiresAt to *time.Time
func (f GrantCreateForm) ExpiresAt() *time.Time {
	// convert to time.Time
	var t *time.Time
	if f.ExpiresIn != 0 {
		*t = time.Now().Add(time.Second * time.Duration(f.ExpiresIn))
	}
	return t
}

// GrantCreate create/update a grant
func GrantCreate(ctx *web.Context, r APIRender, db *models.DB, f GrantCreateForm) {
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
