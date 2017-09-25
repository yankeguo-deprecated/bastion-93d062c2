package main

import (
	"fmt"
	"testing"

	"ireul.com/bastion/types"
)

func TestScript(t *testing.T) {
	d := SyncData{
		BaseDir:     "/tmp",
		AccountsAdd: []string{"bastion-ryan1", "bastion-ryan2"},
		Accounts: []types.Account{{
			Account:   "bastion-ryan",
			CanSudo:   true,
			PublicKey: "ssh-rsa test1",
		}, {
			Account:   "bastion-ryan1",
			CanSudo:   true,
			PublicKey: "ssh-rsa test2",
		}, {
			Account:   "bastion-ryan2",
			CanSudo:   false,
			PublicKey: "ssh-rsa test3",
		}},
		AccountsRemove: []string{"bastion-ryan3"},
	}
	fmt.Println("#------------------")
	fmt.Println(GenerateSyncScript(d))
	fmt.Println("#------------------")
}
