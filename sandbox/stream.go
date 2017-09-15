package sandbox

import (
	"io"

	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/stdcopy"
)

// ExecStream 用来封装 Docker 容器 与 外部 io.Reader, io.Writer 逻辑的
type ExecStream struct {
	Resp  types.HijackedResponse
	In    io.Reader
	Out   io.Writer
	IsPty bool
}

func (s ExecStream) setupInput() chan error {
	ch := make(chan error, 1)
	go func() {
		_, e := io.Copy(s.Resp.Conn, s.In)
		s.Resp.CloseWrite()
		ch <- e
	}()
	return ch
}

func (s ExecStream) setupOutput() chan error {
	ch := make(chan error, 1)
	go func() {
		var e error
		if s.IsPty {
			_, e = io.Copy(s.Out, s.Resp.Reader)
		} else {
			_, e = stdcopy.StdCopy(s.Out, s.Out, s.Resp.Reader)
		}
		ch <- e
	}()
	return ch
}

// Run 在一个 context 上执行 stream
func (s ExecStream) Run(ctx context.Context) (err error) {
	inDone := s.setupInput()
	outDone := s.setupOutput()
	for {
		select {
		case <-inDone:
			{
				select {
				case e := <-outDone:
					return e
				case <-ctx.Done():
					return ctx.Err()
				}
			}
		case e := <-outDone:
			{
				return e
			}
		case <-ctx.Done():
			{
				return ctx.Err()
			}
		}
	}
	return
}
