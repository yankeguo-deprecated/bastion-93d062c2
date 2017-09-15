package main

import (
	"github.com/urfave/cli"
	_ "ireul.com/mysql"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "bastion"
	app.Usage = "Enterprise Bastion System"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Value: "config.toml",
			Usage: "config file",
		},
	}
	app.Commands = []cli.Command{
		webCommand,
		sandboxAttachCommand,
		userCreateCommand,
	}
	app.Run(os.Args)
}
