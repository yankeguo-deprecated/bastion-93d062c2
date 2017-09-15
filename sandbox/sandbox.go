package sandbox

import (
	"fmt"
	"ireul.com/bastion/models"
	"path"
)

func newSandbox(u models.User, dataDir string) Sandbox {
	s := Sandbox{
		UserID:         u.ID,
		UserLogin:      u.Login,
		UserPublicKey:  u.PublicKey,
		UserPrivateKey: u.PrivateKey,
	}

	s.SharedDir = path.Join(dataDir, "shared")
	s.PrivateDir = path.Join(dataDir, s.ContainerName())
	return s
}

// Sandbox 沙箱环境
type Sandbox struct {
	UserID         uint
	UserLogin      string
	UserPublicKey  string
	UserPrivateKey string

	PrivateDir string
	SharedDir  string
}

// ContainerName 返回沙箱的容器名
func (s Sandbox) ContainerName() string {
	return fmt.Sprintf("sandbox-%s", s.UserLogin)
}

// Hostname 返回沙箱的内部主机名
func (s Sandbox) Hostname() string {
	return fmt.Sprintf("%s.sandbox", s.UserLogin)
}
