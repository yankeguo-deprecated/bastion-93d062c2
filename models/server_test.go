package models

import (
	"fmt"
	"sync"
	"testing"
)

// helper methods to create a Server

var createTestServerLoginMutex = sync.Mutex{}

func createTestServerLogin() string {
	createTestServerLoginMutex.Lock()
	defer createTestServerLoginMutex.Unlock()
	u := &Server{}
	if err := db.Order("id DESC").First(u).Error; err != nil {
		panic(err)
	}
	return "test-" + fmt.Sprintf("%d", u.ID+1)
}

func createTestServer() *Server {
	u := &Server{}
	u.Name = createTestServerLogin()
	if err := db.Create(u).Error; err != nil {
		panic(err)
	}
	return u
}

func TestCreateTestServer(t *testing.T) {
	s := createTestServer()

	if len(s.Token) == 0 {
		t.Error("failed to set server token")
	}
	if len(s.Tags) == 0 {
		t.Error("failed to set tags")
	} else {
		if s.Tags[0] != ServerTagDefault {
			t.Error("failed to set default tag")
		}
	}
}
