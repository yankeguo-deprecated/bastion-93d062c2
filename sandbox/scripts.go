package sandbox

import (
	"bytes"
	"log"
	"text/template"

	"ireul.com/bastion/utils"
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

func createScript(name string, tmpl string, data utils.Map) string {
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
		utils.NewMap("PublicKey", publicKey, "PrivateKey", privateKey),
	)
}
