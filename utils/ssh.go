package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"golang.org/x/crypto/ssh"
	"strings"
)

// GenerateSSHKeyPair 创建 RSA 2048 SSH 密钥对
func GenerateSSHKeyPair() (fp string, pub string, priv string, err error) {
	var k *rsa.PrivateKey
	if k, err = rsa.GenerateKey(rand.Reader, 2048); err != nil {
		return
	}
	var sk ssh.PublicKey
	if sk, err = ssh.NewPublicKey(k.Public()); err != nil {
		return
	}

	fp = ssh.FingerprintSHA256(sk)
	pub = strings.TrimSpace(string(ssh.MarshalAuthorizedKey(sk)))

	kb := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(k)}
	priv = string(pem.EncodeToMemory(kb))

	return
}
