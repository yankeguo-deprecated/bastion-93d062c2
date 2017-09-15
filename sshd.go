package main

import (
	"fmt"
	"ireul.com/cli"
)

// sshdCommand 用来提供 authorized_keys 中的自定义 command
var sshdCommand = cli.Command{
	Name:   "sandbox:attach",
	Usage:  "handle ssh connection",
	Action: execSSHDCommand,
}

func execSSHDCommand(c *cli.Context) error {
	fmt.Println("handle web")
	return nil
}
