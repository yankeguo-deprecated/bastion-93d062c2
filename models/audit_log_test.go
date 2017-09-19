package models

import "testing"

func TestAudit(t *testing.T) {
	testUser := createTestUser()

	var err error
	tk := &Token{UserID: testUser.ID, Desc: "TEST"}
	err = db.Create(tk).Error
	if err != nil {
		t.Error(err)
	}

	a, err := db.Audit(testUser, "test", tk)
	if err != nil {
		t.Error(err)
	}
	if a.TargetDetail != "TEST" {
		t.Error("sourc_detail not set")
	}
	if a.CreatedAt.IsZero() {
		t.Error("created_at not set")
	}
}
