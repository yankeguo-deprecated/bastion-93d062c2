package main

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/crypto/ssh"
	"ireul.com/bastion/models"
	"ireul.com/bastion/sandbox"
	"ireul.com/bastion/utils"
	"ireul.com/cli"
	"ireul.com/sshd"
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
	smc := sandbox.ManagerOptions{
		Image:   cfg.Sandbox.Image,
		DataDir: cfg.Sandbox.DataDir,
	}
	var sm sandbox.Manager
	if sm, err = sandbox.NewManager(smc); err != nil {
		log.Fatalln(err)
		return
	}

	// handle
	sshd.Handle(func(s sshd.Session) {
		// extract User
		u := s.Context().Value("User").(models.User)
		// ensure command
		cmd := s.Command()
		if len(cmd) == 0 {
			cmd = []string{"/bin/bash"}
		}
		// get sandbox
		snb, err := sm.GetOrCreate(u)
		if err != nil {
			io.WriteString(s, fmt.Sprintf("Internal Error: %s\n", err.Error()))
			s.Exit(1)
			return
		}
		// attach sandbox
		pty, sshwinch, isPty := s.Pty()

		// create opts
		opts := sandbox.ExecAttachOptions{
			Command: cmd,
			Reader:  s,
			Writer:  s,
			IsPty:   isPty,
			Term:    pty.Term,
		}

		if isPty {
			snbwinch := make(chan sandbox.Window, 1)
			opts.WindowChan = snbwinch

			// convert channel sshwinch -> snbwinch
			go func() {
				for {
					s, live := <-sshwinch
					if live {
						snbwinch <- sandbox.Window{Height: uint(s.Height), Width: uint(s.Width)}
					} else {
						close(snbwinch)
						break
					}
				}
			}()
		}

		err = sm.ExecAttach(snb, opts)

		if err != nil {
			log.Printf("ERROR: Sandbox ExecAttach Failed: %s\n", err.Error())
		}
	})

	// options
	options := []sshd.Option{
		// set host_key
		sshd.HostKeyFile(cfg.SSHD.HostKeyFile),
		// auth public_key
		sshd.PublicKeyAuth(func(ctx sshd.Context, key sshd.PublicKey) bool {
			// get fingerprint
			fp := ssh.FingerprintSHA256(key)
			// find SSHKey
			k := models.SSHKey{}
			db.Where("fingerprint = ?", fp).First(&k)
			if db.NewRecord(k) {
				log.Printf("ERROR: Invalid Key, FP=%s", fp)
				return false
			}
			// find User
			u := models.User{}
			db.First(&u, k.UserID)
			if db.NewRecord(u) || u.IsBlocked {
				log.Printf("ERROR: User Not Found / Blocked, UserID=%d, FP=%s\n", k.UserID, fp)
				return false
			}
			// set User.ID
			log.Printf("Signed In, UserID=%d, FP=%s\n", k.UserID, fp)
			ctx.SetValue("User", u)
			return true
		}),
	}

	log.Printf("Listening at %s:%d\n", cfg.SSHD.Host, cfg.SSHD.Port)
	log.Fatal(sshd.ListenAndServe(fmt.Sprintf("%s:%d", cfg.SSHD.Host, cfg.SSHD.Port), nil, options...))
	return nil
}
