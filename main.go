package main

import (
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "bastion"
	app.Usage = "Enterprise Bastion System"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "conf",
			Value: "config.toml",
			Usage: "config file",
		},
	}
	app.Commands = []cli.Command{
		webCommand,
		handleCommand,
		newUserCommand,
	}
	app.Run(os.Args)
}
