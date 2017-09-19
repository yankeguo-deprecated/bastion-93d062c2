package models

import "testing"

func TestTokenBeforeSave(t *testing.T) {
	u := createTestUser()

	tk := &Token{UserID: u.ID, Desc: "TEST"}
	db.Create(tk)

	if len(tk.Secret) == 0 {
		t.Error("failed to set token.secret automatically")
	}
}
