package utils

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

// Config 配置
type Config struct {
	Web struct {
		// Domain 域名地址，用于展示
		Domain string `toml:"domain"`
		// Host 主机，用于 IP 绑定
		Host string `toml:"host"`
		// Port 端口号
		Port int `toml:"port"`
	} `toml:"web"`
	Redis struct {
		// URL Redis 数据库链接
		URL string `toml:"url"`
	} `toml:"redis"`
	Database struct {
		// URL MySQL 数据库链接
		URL string `toml:"url"`
	} `toml:"db"`
	Bastion struct {
		// Env 运行环境，可以是 development, production, test
		Env string `toml:"env"`
		// SandboxDir 沙箱的数据地址
		SandboxDir string `toml:"sandbox_dir"`
		// MasterKeyFile Bastion 主密钥的地址
		MasterKeyFile string `toml:"master_key_file"`
		// AuthorizedKeys Bastion 要写入的 autorized_keys 文件
		AuthorizedKeysFile string `toml:"authorized_keys_file"`
	} `toml:"bastion"`
}

// ParseConfigFile 加载一个 TOML 配置文件
func ParseConfigFile(file string) (*Config, error) {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return ParseConfig(string(bs))
}

// ParseConfig 解析 TOML
func ParseConfig(s string) (*Config, error) {
	config := Config{}
	if _, err := toml.Decode(s, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
