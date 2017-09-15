package main

import (
	"log"

	"ireul.com/bastion/models"
	"ireul.com/bastion/routes"
	"ireul.com/bastion/sandbox"
	"ireul.com/bastion/utils"
	"ireul.com/cli"
	"ireul.com/redis"
	"ireul.com/web"
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

	// create web instance
	m := web.Classic()
	m.Use(web.Renderer())
	m.Use(func(ctx *web.Context) {
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

	// create sandbox.Manager
	smc := sandbox.ManagerOptions{
		Image:   cfg.Sandbox.Image,
		DataDir: cfg.Sandbox.DataDir,
	}
	var sm sandbox.Manager
	if sm, err = sandbox.NewManager(smc); err != nil {
		log.Fatalln(err)
		return
	}
	m.Map(sm)

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

	// run web instance
	m.Run(cfg.Web.Host, cfg.Web.Port)
	return
}
