package utils

import (
	"io/ioutil"
	"os"

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
	_, err = os.Stat(c.SSHD.HostKeyFile)
	return
}
