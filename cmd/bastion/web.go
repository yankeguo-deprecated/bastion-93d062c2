package main

import (
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pagoda-tech/bastion/conf"
	"github.com/pagoda-tech/bastion/models"
	"github.com/pagoda-tech/macaron"
	"github.com/urfave/cli"
	"log"
)

// webCommand 用来启动 Web 服务
var webCommand = cli.Command{
	Name:    "web",
	Aliases: []string{"w"},
	Usage:   "start the web server",
	Action:  execWebCommand,
}

func execWebCommand(c *cli.Context) (err error) {
	// setup log
	log.SetPrefix("[bastion-web] ")

	// decode config
	var cfg *conf.Config
	if cfg, err = conf.DecodeFile(c.GlobalString("conf")); err != nil {
		log.Fatalln(err)
		return
	}

	// create macaron instance
	m := macaron.Classic()
	m.SetEnv(cfg.Bastion.Env)

	// map config
	m.Map(cfg)

	// create xorm engine and map
	var db *gorm.DB
	if db, err = gorm.Open("mysql", cfg.Database.URL); err != nil {
		log.Fatalln(err)
		return
	}
	if m.Env() == macaron.DEV {
		db.LogMode(true)
	}
	models.AutoMigrate(db)
	m.Map(db)

	// create redis and map
	var ro *redis.Options
	if ro, err = redis.ParseURL(cfg.Redis.URL); err != nil {
		log.Fatalln(err)
		return
	}
	r := redis.NewClient(ro)
	m.Map(r)

	// routes
	mountRoutes(m)

	// run macaron instance
	m.Run(cfg.Web.Host, cfg.Web.Port)
	return
}

func mountRoutes(m *macaron.Macaron) {
	m.Get("/", func() string {
		return "Hello world!"
	})
}
