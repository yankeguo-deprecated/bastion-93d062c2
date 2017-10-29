package sandbox

import (
	"bytes"
	"log"
	"text/template"

	"ireul.com/com"
)

const tplSeedSSH = `#!/bin/bash
# create /root/.ssh
mkdir -p /root/.ssh
chmod 700 /root/.ssh
cd /root/.ssh

# write keys
echo "{{.PublicKey}}"  > id_rsa.pub
chmod 644 id_rsa.pub
echo "{{.PrivateKey}}" > id_rsa
chmod 600 id_rsa

# write README
echo "id_rsa 和 id_rsa.pub 受 Bunker 管理，请勿修改" > README
`

const tplSSHConfig = `#!/bin/bash
# remove .ssh/config
rm -f /root/.ssh/config

# create new .ssh/config
{{range .Entries}}
echo "Host {{.Name}}" >> /root/.ssh/config
echo "  HostName {{.Host}}" >> /root/.ssh/config
echo "  Port {{.Port}}" >> /root/.ssh/config
echo "  User {{.User}}" >> /root/.ssh/config
{{end}}
`

// SSHEntry a entry in ssh_config
type SSHEntry struct {
	Name string
	Host string
	Port uint
	User string
}

func createScript(name string, tmpl string, data com.Map) string {
	t, err := template.New(name).Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	buf := &bytes.Buffer{}
	t.Execute(buf, data)
	return buf.String()
}

func scriptSeedSSH(publicKey string, privateKey string) string {
	return createScript(
		"seed-ssh",
		tplSeedSSH,
		com.NewMap("PublicKey", publicKey, "PrivateKey", privateKey),
	)
}

// ScriptSeedSSHConfig create a script for seeding .ssh/config
func ScriptSeedSSHConfig(entries []SSHEntry) string {
	return createScript(
		"seed-ssh-config",
		tplSSHConfig,
		com.NewMap("Entries", entries),
	)
}
