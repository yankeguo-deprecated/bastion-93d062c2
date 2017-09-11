package utils

import "testing"

func TestGenerateSSHKeyPair(t *testing.T) {
	_, _, _, err := GenerateSSHKeyPair()
	if err != nil {
		t.Error(err)
		return
	}
}
