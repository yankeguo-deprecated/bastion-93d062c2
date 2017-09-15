package main

import (
	"github.com/urfave/cli"
	"ireul.com/bastion/models"
	"ireul.com/bastion/sandbox"
	"ireul.com/bastion/utils"
	"log"
)

// userCreateCommand 用来启动 Web 服务
var userCreateCommand = cli.Command{
	Name:   "user:create",
	Usage:  "create a user",
	Action: execNewUserCommand,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "login",
			Usage: "login name of new user",
		},
		cli.StringFlag{
			Name:  "password",
			Usage: "password of new user",
		},
		cli.BoolFlag{
			Name:  "is_admin",
			Usage: "if new user is an admin",
		},
	},
}

func execNewUserCommand(c *cli.Context) (err error) {
	// setup log
	log.SetPrefix("[bastion-cli] ")

	// login
	login := c.String("login")

	if !models.UserLoginRegexp.MatchString(login) {
		log.Fatalln("invalid user login")
		return
	}

	password := c.String("password")
	if len(password) < models.UserPasswordMinLen {
		log.Fatalln("user password is too short")
		return
	}

	// decode config
	var cfg *utils.Config
	if cfg, err = utils.ParseConfigFile(c.GlobalString("config")); err != nil {
		log.Fatalln(err)
		return
	}

	// db
	var db *models.DB
	if db, err = models.NewDB(cfg); err != nil {
		log.Fatalln(err)
		return
	}
	db.AutoMigrate()

	user := &models.User{
		Login: login,
	}
	if err = user.SetPassword(password); err != nil {
		log.Fatalln(err)
		return
	}
	if err = user.GenerateSSHKey(); err != nil {
		log.Fatalln(err)
		return
	}
	user.Nickname = user.Login
	user.IsAdmin = c.Bool("is_admin")

	if err = db.Create(user).Error; err != nil {
		log.Fatalln(err)
		return
	}

	smc := sandbox.ManagerOptions{
		Image:   cfg.Sandbox.Image,
		DataDir: cfg.Sandbox.DataDir,
	}

	sbm, err := sandbox.NewManager(smc)

	if err != nil {
		log.Fatalln(err)
	}

	_, err = sbm.GetOrCreate(*user)
	if err != nil {
		log.Fatalln(err)
	}

	return
}
