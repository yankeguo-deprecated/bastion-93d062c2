package sandbox

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/pagoda-tech/bastion/models"
	"os"
	"sync"
)

const dirPerm = os.FileMode(0750)

// ManagerOptions 沙箱管理器选项
type ManagerOptions struct {
	ImageName string
	DataDir   string
}

// Manager 沙箱管理器
type Manager interface {
	GetOrCreate(u models.User) (Sandbox, error)
}

// 沙箱管理工具
type managerImpl struct {
	options ManagerOptions
	mutex   *sync.Mutex
}

func (m managerImpl) exists(s Sandbox) (bool, error) {
	_, err := cli.ContainerInspect(context.Background(), s.ContainerName())
	if err != nil {
		if client.IsErrContainerNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (m managerImpl) GetOrCreate(u models.User) (Sandbox, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	s := newSandbox(u, m.options.DataDir)

	if e, err := m.exists(s); e || err != nil {
		return Sandbox{}, err
	}

	return s, m.create(s)
}

func (m managerImpl) create(s Sandbox) error {
	if err := os.MkdirAll(s.RootDir, dirPerm); err != nil {
		return err
	}
	if err := os.MkdirAll(s.SharedDir, dirPerm); err != nil {
		return err
	}
	ccfg := &container.Config{
		Hostname: s.Hostname(),
		Image:    m.options.ImageName,
	}
	hcfg := &container.HostConfig{
		Binds: []string{
			fmt.Sprintf("%s:/root", s.RootDir),
			fmt.Sprintf("%s:/shared", s.SharedDir),
		},
		RestartPolicy: container.RestartPolicy{
			Name: "always",
		},
	}

	ret, err := cli.ContainerCreate(context.Background(), ccfg, hcfg, nil, s.ContainerName())

	if err != nil {
		return err
	}

	return cli.ContainerStart(context.Background(), ret.ID, types.ContainerStartOptions{})
}

// NewManager 创建
func NewManager(opt ManagerOptions) (Manager, error) {
	if err := os.MkdirAll(opt.DataDir, dirPerm); err != nil {
		return nil, err
	}
	m := managerImpl{
		options: opt,
		mutex:   &sync.Mutex{},
	}
	return m, nil
}
