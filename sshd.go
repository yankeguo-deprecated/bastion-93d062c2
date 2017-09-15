package main

import (
	"fmt"
	"io"
	"log"

	gossh "golang.org/x/crypto/ssh"
	"ireul.com/bastion/models"
	"ireul.com/bastion/utils"
	"ireul.com/cli"
	"ireul.com/ssh"
)

// sshdCommand 用来提供 authorized_keys 中的自定义 command
var sshdCommand = cli.Command{
	Name:   "sshd",
	Usage:  "start sshd server",
	Action: execSSHDCommand,
}

func execSSHDCommand(c *cli.Context) (err error) {
	// setup log
	log.SetPrefix("[bastion-sshd] ")

	// decode config
	var cfg *utils.Config
	if cfg, err = utils.ParseConfigFile(c.GlobalString("config")); err != nil {
		log.Fatalln(err)
		return
	}

	// create models.DB
	var db *models.DB
	if db, err = models.NewDB(cfg); err != nil {
		log.Fatalln(err)
		return
	}
	db.AutoMigrate()

	// create sandbox.Manager
	/*
		smc := sandbox.ManagerOptions{
			Image:   cfg.Sandbox.Image,
			DataDir: cfg.Sandbox.DataDir,
		}
		var sm sandbox.Manager
		if sm, err = sandbox.NewManager(smc); err != nil {
			log.Fatalln(err)
			return
		}
	*/

	// handle
	ssh.Handle(func(s ssh.Session) {
		io.WriteString(s, fmt.Sprintf("Hello world, user %d\n", s.Context().Value("UserID")))
	})

	// options
	options := []ssh.Option{
		// set host_key
		ssh.HostKeyFile(cfg.SSHD.HostKeyFile),
		// auth public_key
		ssh.PublicKeyAuth(func(ctx ssh.Context, key ssh.PublicKey) bool {
			fp := gossh.FingerprintSHA256(key)
			k := models.SSHKey{}
			db.Where("fingerprint = ?", fp).First(&k)
			if db.NewRecord(k) {
				log.Println("key not found:", fp)
				return false
			}
			u := models.User{}
			db.First(&u, k.UserID)
			if db.NewRecord(u) || u.IsBlocked {
				log.Println("user id not found or blocked:", k.UserID, ", fingerprint:", fp)
				return false
			}
			log.Println("user signed in:", k.UserID, ", fingerprint:", fp)
			ctx.SetValue("UserID", k.UserID)
			return true
		}),
	}

	log.Printf("Listening at %s:%d\n", cfg.SSHD.Host, cfg.SSHD.Port)
	log.Fatal(ssh.ListenAndServe(fmt.Sprintf("%s:%d", cfg.SSHD.Host, cfg.SSHD.Port), nil, options...))
	return nil
}
