package main

import (
	"fmt"
	"github.com/urfave/cli"
)

// sandboxAttachCommand 用来提供 authorized_keys 中的自定义 command
var sandboxAttachCommand = cli.Command{
	Name:   "sandbox:attach",
	Usage:  "handle ssh connection",
	Action: execSandboxAttachCommand,
}

func execSandboxAttachCommand(c *cli.Context) error {
	fmt.Println("handle web")
	return nil
}
