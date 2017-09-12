package main

import (
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pagoda-tech/bastion/models"
	"github.com/pagoda-tech/bastion/routes"
	"github.com/pagoda-tech/bastion/utils"
	"github.com/pagoda-tech/macaron"
	"github.com/urfave/cli"
	"log"
)

// webCommand 用来启动 Web 服务
var webCommand = cli.Command{
	Name:   "web",
	Usage:  "start the web server",
	Action: execWebCommand,
}

func execWebCommand(c *cli.Context) (err error) {
	// setup log
	log.SetPrefix("[bastion-web] ")

	// create macaron instance
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Use(func(ctx *macaron.Context) {
		ctx.Data["Version"] = VERSION
	})

	// decode config
	var cfg *utils.Config
	if cfg, err = utils.ParseConfigFile(c.GlobalString("config")); err != nil {
		log.Fatalln(err)
		return
	}

	// map config
	m.SetEnv(cfg.Bastion.Env)
	m.Map(cfg)

	// map DB
	var db *models.DB
	if db, err = models.NewDB(cfg); err != nil {
		log.Fatalln(err)
		return
	}
	db.AutoMigrate()
	m.Map(db)

	// map redis client
	var ro *redis.Options
	if ro, err = redis.ParseURL(cfg.Redis.URL); err != nil {
		log.Fatalln(err)
		return
	}
	r := redis.NewClient(ro)
	m.Map(r)

	// routes
	routes.Mount(m)

	// run macaron instance
	m.Run(cfg.Web.Host, cfg.Web.Port)
	return
}
