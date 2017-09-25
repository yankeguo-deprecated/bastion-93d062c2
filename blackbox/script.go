package main

import (
	"bytes"
	"text/template"

	"ireul.com/bastion/types"
)

const tmplSync = `#!/bin/bash
## Create Account

{{range $.AccountsAdd}}
# create account
useradd --create-home --home-dir "{{$.BaseDir}}/{{.}}" --shell /bin/bash {{.}}
{{end}}

## Update Account

{{range $.Accounts}}
# lock the passwd 
passwd -l {{.Account}}
# unlock account
usermod --expiredate "" {{.}}
# update authorized_keys
mkdir -p "{{$.BaseDir}}/{{.Account}}/.ssh"
echo "{{.PublicKey}}" > "{{$.BaseDir}}/{{.Account}}/.ssh/authorized_keys"
chown {{.Account}}:{{.Account}} "{{$.BaseDir}}/{{.Account}}/.ssh"
chown {{.Account}}:{{.Account}} "{{$.BaseDir}}/{{.Account}}/.ssh/authorized_keys"
chmod 700 "{{$.BaseDir}}/{{.Account}}/.ssh"
chmod 600 "{{$.BaseDir}}/{{.Account}}/.ssh/authorized_keys"
# update sudo status
gpasswd {{if .CanSudo}}-a{{else}}-d{{end}} {{.Account}} sudo || true
{{end}}

## Remove Account

{{range $.AccountsRemove}}
# lock account
usermod --expiredate 1 {{.}}
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
