package utils

import (
	"crypto/rsa"
	"errors"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/ssh"

	"ireul.com/bastion/types"
	"ireul.com/toml"
)

// ParseConfigFile parse a config toml file
func ParseConfigFile(file string) (*types.Config, error) {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return ParseConfig(string(bs))
}

// ParseConfig parse a config toml string
func ParseConfig(s string) (*types.Config, error) {
	config := types.Config{}
	if _, err := toml.Decode(s, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

// ValidateConfig validate the Config file
func ValidateConfig(c *types.Config) (err error) {
	// check SSHD.HostKeyFile
	if _, err = os.Stat(c.SSHD.HostKeyFile); err != nil {
		return
	}
	// read master key file
	hk, err := ioutil.ReadFile(c.Bastion.MasterKeyFile)
	if err != nil {
		return
	}
	// parse host key
	vk, err := ssh.ParseRawPrivateKey(hk)
	if err != nil {
		return
	}
	// cast to rsa.PrivateKey
	rvk, ok := vk.(*rsa.PrivateKey)
	if !ok {
		return errors.New("only RSA SSH host key is supported")
	}
	// create ssh.PublicKey
	pk, err := ssh.NewPublicKey(rvk.Public())
	if err != nil {
		return
	}
	// marshal to SSH wired format
	c.Bastion.MasterPublicKey = string(ssh.MarshalAuthorizedKey(pk))
	if _, err = os.Stat(c.Bastion.MasterKeyFile); err != nil {
		return
	}
	return nil
}
