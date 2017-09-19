package models

import (
	"fmt"
	"sync"
	"testing"
)

// helper methods to create a User

var createTestUserLoginMutex = sync.Mutex{}

func createTestUserLogin() string {
	createTestUserLoginMutex.Lock()
	defer createTestUserLoginMutex.Unlock()
	u := &User{}
	if err := db.Order("id DESC").First(u).Error; err != nil {
		panic(err)
	}
	return "test-" + fmt.Sprintf("%d", u.ID+1)
}

func createTestUser() *User {
	u := &User{}
	u.Login = createTestUserLogin()
	if err := db.Create(u).Error; err != nil {
		panic(err)
	}
	return u
}

// test helper it-self

func TestCreateTestUser(t *testing.T) {
	u := createTestUser()
	if u.Nickname != u.Login {
		t.Error("failed to set nickname automatically")
	}
	if len(u.Fingerprint) == 0 || len(u.PublicKey) == 0 || len(u.PrivateKey) == 0 {
		t.Error("public key / fingerprint / private key is not set automatically")
	}
	if u.UsedAt != nil {
		t.Error("newly created user shall not have UsedAt set")
	}
}
