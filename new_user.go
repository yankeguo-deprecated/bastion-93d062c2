package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/pagoda-tech/bastion/conf"
	"github.com/pagoda-tech/bastion/models"
	"github.com/urfave/cli"
	"log"
)

// newUserCommand 用来启动 Web 服务
var newUserCommand = cli.Command{
	Name:   "new-user",
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
	var cfg *conf.Config
	if cfg, err = conf.DecodeFile(c.GlobalString("conf")); err != nil {
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

	return
}
