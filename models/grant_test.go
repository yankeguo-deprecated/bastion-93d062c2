package models

import (
	"testing"
	"time"
)

func TestCreateGrant(t *testing.T) {
	u := createTestUser()

	exp := time.Now().Add(-time.Hour * 2)
	n := &Grant{
		UserID:    u.ID,
		Tag:       "test",
		ExpiresAt: &exp,
	}
	db.Create(n)

	if !n.IsExpiredNow() {
		t.Error("should be expired now")
	}
	if n.IsExpired(exp.Add(-time.Second)) {
		t.Error("should be not expired compared with given time")
	}

	n.ExpiresAt = nil
	db.Save(n)

	if n.IsExpired(time.Now()) {
		t.Error("should not be expired with expiresAt = NULL")
	}
}

func TestCompactSliceGrant(t *testing.T) {
	ins := []Grant{
		{
			UserID:  1,
			CanSudo: false,
		},
		{
			UserID:  1,
			CanSudo: true,
		},
		{
			UserID:  2,
			CanSudo: false,
		},
		{
			UserID:  2,
			CanSudo: false,
		},
	}
	ns := CompactSliceGrant(ins)
	if ns[1].CanSudo != true {
		t.Error("failed to override can sudo")
	}
	if ns[2].UserID != 2 {
		t.Error("failed to set")
	}
}
