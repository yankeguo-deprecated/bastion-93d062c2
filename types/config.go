package types

import (
	"io/ioutil"
	"os"

	"ireul.com/toml"
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
		// MasterKeyFile Bastion 主密钥的地址
		MasterKeyFile string `toml:"master_key_file"`
	} `toml:"bastion"`
	Sandbox struct {
		// 沙箱的镜像名称
		Image string `toml:"image"`
		// 沙箱的数据地址
		DataDir string `toml:"data_dir"`
	} `toml:"sandbox"`
	SSHD struct {
		// SSHD 服务的绑定地址
		Host string `toml:"host"`
		// SSHD 服务的端口号
		Port int `toml:"port"`
		// SSHD 的主机密钥地址，RSA
		HostKeyFile string `toml:"host_key_file"`
	}
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

// Validate validate the Config file
func (c Config) Validate() (err error) {
	if _, err = os.Stat(c.SSHD.HostKeyFile); err != nil {
		return
	}
	if _, err = os.Stat(c.Bastion.MasterKeyFile); err != nil {
		return
	}
	return nil
}
