package main

import (
	"ireul.com/cli"
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
		sshdCommand,
		userCreateCommand,
	}
	app.Run(os.Args)
}
