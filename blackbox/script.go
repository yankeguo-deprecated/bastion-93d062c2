package main

import (
	"bytes"
	"text/template"

	"ireul.com/bastion/types"
)

const tmplSync = `#!/bin/bash
set -e
set -u

# Global Variables

HOME_BASE="{{$.BaseDir}}"
SUDOERS_TMP=/tmp/bastion-sudoers.tmp

# Functions

create_user() {
	useradd --create-home --home-dir "$HOME_BASE/$1" $1
}

clean_sudoers_tmp() {
	rm -f $SUDOERS_TMP
	touch -f $SUDOERS_TMP
}

apply_sudoers_tmp() {
	mv -f $SUDOERS_TMP /etc/sudoers.d/bastion-sudoers
}

unlock_account() {
	usermod -L -e "" $1
}

update_account_authorized_keys() {
	SSH_DIR="$HOME_BASE/$1/.ssh"
	mkdir -p "$SSH_DIR"
	pushd "$SSH_DIR"
	echo "$2" > authorized_keys
	chown -R $1:$1 .
	chmod 700 .
	chmod 600 authorized_keys
	popd
}

add_account_to_sudoers_tmp() {
	echo "$1 ALL=(ALL:ALL) NOPASSWD:ALL" >> $SUDOERS_TMP
}

lock_account() {
	usermod -L -e 1 $1
}

{{range $.AccountsAdd}}
create_user {{.}}
{{end}}

clean_sudoers_tmp

{{range $.Accounts}}
unlock_account {{.Account}}
update_account_authorized_keys {{.Account}} "{{.PublicKey}}"
{{if .CanSudo}}
add_account_to_sudoers_tmp {{.Account}}
{{end}}
{{end}}

apply_sudoers_tmp

{{range $.AccountsRemove}}
lock_account {{.}}
{{end}}
`

// SyncData data for script template
type SyncData struct {
	BaseDir        string
	AccountsAdd    []string        // add missing accounts
	Accounts       []types.Account // update authorized_keys, lock password, sudo
	AccountsRemove []string        // disable missing accounts
}

// GenerateSyncScript generate a script
func GenerateSyncScript(data SyncData) string {
	tmpl, err := template.New("sync.sh").Parse(tmplSync)
	if err != nil {
		panic(err)
	}
	buf := &bytes.Buffer{}
	tmpl.Execute(buf, data)
	return buf.String()
}
