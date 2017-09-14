package sandbox

import (
	"github.com/docker/docker/client"
)

var cli *client.Client

func init() {
	var err error
	if cli, err = client.NewEnvClient(); err != nil {
		panic(err)
	}
}
