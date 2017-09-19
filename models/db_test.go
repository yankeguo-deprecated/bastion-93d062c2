package models

import (
	"os"
	"testing"
	"time"

	_ "ireul.com/mysql"
)

// init default db for testing

var db *DB

func init() {
	var err error
	db, err = NewDB("test", os.Getenv("TEST_DB_URL"))
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate()
	if err != nil {
		panic(err)
	}
}

// test Touch

func TestTouch(t *testing.T) {
	u := createTestUser()
	db.Touch(u)
	if u.UsedAt == nil {
		t.Error("failed to set usedAt")
	} else {
		dif := time.Now().Sub(*u.UsedAt)
		if dif < 0 || dif > time.Second*10 {
			t.Error("used_at is not set properly")
		}
	}
}
