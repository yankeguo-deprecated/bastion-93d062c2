package main

import (
	"fmt"
	"github.com/urfave/cli"
)

// handleCommand 用来提供 authorized_keys 中的自定义 command
var handleCommand = cli.Command{
	Name:    "handle",
	Aliases: []string{"n"},
	Usage:   "handle ssh connection",
	Action:  execHandleCommand,
}

func execHandleCommand(c *cli.Context) error {
	fmt.Println("handle web")
	return nil
}
