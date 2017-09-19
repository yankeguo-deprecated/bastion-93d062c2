package models

import (
	"time"

	"ireul.com/orm"
)

// Grant is a grant given to a user
type Grant struct {
	orm.Model
	UserID    uint       `orm:"unique_index:uix_tag_user_id" json:"userId"`
	Tag       string     `orm:"unique_index:uix_tag_user_id" json:"tag"`
	ExpiresAt *time.Time `orm:"index" json:"expiresAt"`
	CanSudo   bool       `json:"canSudo"`
}

// IsExpiredNow check if Grant is expired now
func (n Grant) IsExpiredNow() bool {
	return n.ExpiresAt != nil && n.IsExpired(time.Now())
}

// IsExpired check if Grant is expired
func (n Grant) IsExpired(t time.Time) bool {
	return n.ExpiresAt != nil && t.After(*n.ExpiresAt)
}

// CompactSliceGrant returns a array of Grant with unique UserID and larger permission (CanSudo), but ID is removed
func CompactSliceGrant(ins []Grant) map[uint]Grant {
	now := time.Now()
	outs := make(map[uint]Grant, len(ins))
	for _, in := range ins {
		if in.IsExpired(now) {
			continue
		}
		if !outs[in.UserID].CanSudo {
			outs[in.UserID] = Grant{
				UserID:  in.UserID,
				CanSudo: in.CanSudo,
			}
		}
	}
	return outs
}
