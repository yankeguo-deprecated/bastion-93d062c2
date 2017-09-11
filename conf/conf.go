package conf

import (
	"github.com/BurntSushi/toml"
)

// Config 配置
type Config struct {
	Web      Web      `toml:"web"`
	Redis    Redis    `toml:"redis"`
	Database Database `toml:"db"`
	Bastion  Bastion  `toml:"bastion"`
}

// Web web 接口配置
type Web struct {
	// Domain 域名地址，用于展示
	Domain string `toml:"domain"`
	// Host 主机，用于 IP 绑定
	Host string `toml:"host"`
	// Port 端口号
	Port int `toml:"port"`
}

// Redis Redis 数据库配置
type Redis struct {
	// URL Redis 数据库链接
	URL string `toml:"url"`
}

// Database 数据库相关配置
type Database struct {
	// URL MySQL 数据库链接
	URL string `toml:"url"`
}

// Bastion 相关配置
type Bastion struct {
	// Env 运行环境，可以是 development, production, test
	Env string `toml:"env"`
	// SandboxDir 沙箱的数据地址
	SandboxDir string `toml:"sandbox_dir"`
	// MasterKeyFile Bastion 主密钥的地址
	MasterKeyFile string `toml:"master_key_file"`
	// AuthorizedKeys Bastion 要写入的 autorized_keys 文件
	AuthorizedKeysFile string `toml:"authorized_keys_file"`
}

// DecodeFile 加载一个 TOML 配置文件
func DecodeFile(file string) (*Config, error) {
	config := Config{}
	if _, err := toml.DecodeFile(file, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
