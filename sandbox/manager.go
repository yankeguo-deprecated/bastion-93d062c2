package sandbox

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"ireul.com/bastion/models"
)

const dirPerm = os.FileMode(0750)

// ManagerOptions 沙箱管理器选项
type ManagerOptions struct {
	Image   string
	DataDir string
}

// Manager 沙箱管理器
type Manager interface {
	// GetOrCreate 获取或者创建一个 Sandbox
	GetOrCreate(u models.User) (Sandbox, error)
	// ExecScript 使用 /bin/bash 执行一段 Shell 脚本
	ExecScript(s Sandbox, sc string) error
	// ExecAttach 执行命令，并连接 stdin / stdout 和 可选的 TTY
	ExecAttach(s Sandbox, opts ExecAttachOptions) error
}

// 沙箱管理工具
type managerImpl struct {
	options ManagerOptions
	mutex   *sync.Mutex
}

// Window 代表一个 Pty 的窗口大小
type Window struct {
	Width  uint
	Height uint
}

// ExecAttachOptions 代表执行 ExecAttach 所需参数
type ExecAttachOptions struct {
	Term       string
	Command    []string
	Reader     io.Reader
	Writer     io.Writer
	IsPty      bool
	WindowChan chan Window
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
		return s, err
	}

	return s, m.create(s)
}

func (m managerImpl) create(s Sandbox) error {
	if err := os.MkdirAll(s.PrivateDir, dirPerm); err != nil {
		return err
	}
	if err := os.MkdirAll(s.SharedDir, dirPerm); err != nil {
		return err
	}
	ccfg := &container.Config{
		Hostname: s.Hostname(),
		Image:    m.options.Image,
	}
	hcfg := &container.HostConfig{
		Binds: []string{
			fmt.Sprintf("%s:/root", s.PrivateDir),
			fmt.Sprintf("%s:/shared", s.SharedDir),
		},
		RestartPolicy: container.RestartPolicy{
			Name: "always",
		},
	}

	if _, err := cli.ContainerCreate(context.Background(), ccfg, hcfg, nil, s.ContainerName()); err != nil {
		return err
	}

	if err := cli.ContainerStart(context.Background(), s.ContainerName(), types.ContainerStartOptions{}); err != nil {
		return err
	}

	return m.ExecScript(s, scriptSeedSSH(s.UserPublicKey, s.UserPrivateKey))
}

func (m managerImpl) ExecScript(s Sandbox, sc string) error {
	id, err := cli.ContainerExecCreate(
		context.Background(),
		s.ContainerName(),
		types.ExecConfig{
			AttachStdin: true,
			Cmd: []string{
				"/bin/bash",
			},
		},
	)
	if err != nil {
		return err
	}

	r, err := cli.ContainerExecAttach(context.Background(), id.ID, types.ExecConfig{})
	if err != nil {
		return err
	}
	defer r.Close()

	scr := bytes.NewReader([]byte(sc))
	_, err = io.Copy(r.Conn, scr)

	return err
}

func (m managerImpl) ExecAttach(s Sandbox, opts ExecAttachOptions) (err error) {
	execCfg := types.ExecConfig{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          opts.IsPty,
		Cmd:          opts.Command,
	}

	if opts.IsPty {
		execCfg.Env = []string{
			fmt.Sprintf("TERM=%s", opts.Term),
		}
	}

	// create ExecID
	id, err := cli.ContainerExecCreate(
		context.Background(),
		s.ContainerName(),
		execCfg,
	)

	if err != nil {
		return
	}

	ctx := context.Background()

	// attach with ExecID
	r, err := cli.ContainerExecAttach(ctx, id.ID, execCfg)
	if err != nil {
		return err
	}
	defer r.Close()

	// pipe WindowChan
	if opts.IsPty {
		go func() {
			for {
				w, live := <-opts.WindowChan
				if !live {
					break
				}
				cli.ContainerExecResize(
					context.Background(),
					id.ID,
					types.ResizeOptions{
						Width:  w.Width,
						Height: w.Height,
					})
			}
		}()
	}

	// stream
	stream := ExecStream{Resp: r, In: opts.Reader, Out: opts.Writer, IsPty: opts.IsPty}
	err = stream.Run(ctx)

	return
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
