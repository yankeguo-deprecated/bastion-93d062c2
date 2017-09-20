package types

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
