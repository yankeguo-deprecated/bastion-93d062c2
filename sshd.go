package main

import (
	"fmt"
	"io"
	"log"

	"golang.org/x/crypto/ssh"
	"ireul.com/bastion/models"
	"ireul.com/bastion/sandbox"
	"ireul.com/bastion/types"
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
	var cfg *types.Config
	if cfg, err = utils.ParseConfigFile(c.GlobalString("config")); err != nil {
		log.Fatalln(err)
		return
	}
	if err = utils.ValidateConfig(cfg); err != nil {
		log.Fatalln(err)
		return
	}

	// create models.DB
	var db *models.DB
	if db, err = models.NewDB(cfg.Bastion.Env, cfg.Database.URL); err != nil {
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
		k := s.Context().Value("SSHKey").(models.SSHKey)
		// get sandbox
		snb, err := sm.GetOrCreate(u)
		if err != nil {
			io.WriteString(s, fmt.Sprintf("Internal Error: %s\n", err.Error()))
			s.Exit(1)
			return
		}
		// check u.IsBlocked
		if u.IsBlocked {
			db.Audit(k, "ssh.blocked", snb)
			io.WriteString(s, fmt.Sprintf("User(%d:%s) is blocked, this incident will be reported.\n", u.ID, u.Login))
			s.Exit(1)
			return
		}
		// touch and audit
		db.Touch(u)
		db.Touch(k)
		db.Audit(k, "ssh.success", snb)
		// update .ssh/config
		ss := []models.Server{}
		db.Find(&ss)
		entries := []sandbox.SSHEntry{}
		for _, s := range ss {
			entries = append(entries, sandbox.SSHEntry{
				Name: s.Name,
				Port: s.Port,
				Host: s.Address,
				User: types.AccountPrefix + u.Login,
			})
		}
		err = sm.ExecScript(snb, sandbox.ScriptSeedSSHConfig(entries))
		if err != nil {
			io.WriteString(s, "failed to seed .ssh/config")
		}
		// ensure command
		cmd := s.Command()
		if len(cmd) == 0 {
			cmd = []string{"/bin/bash"}
		}
		// create opts
		pty, sshwinch, isPty := s.Pty()
		opts := sandbox.ExecAttachOptions{
			Command: cmd,
			Reader:  s,
			Writer:  s,
			IsPty:   isPty,
			Term:    pty.Term,
		}
		// convert channel sshwinch -> snbwinch
		if isPty {
			snbwinch := make(chan sandbox.Window, 1)
			opts.WindowChan = snbwinch

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
		// attach sandbox
		err = sm.ExecAttach(snb, opts)
		if err != nil {
			log.Printf("ERROR: Sandbox ExecAttach Failed: %s\n", err.Error())
			io.WriteString(s, fmt.Sprintf("Failed to attach sandbox %s, %s\n", snb.ContainerName(), err.Error()))
			s.Exit(1)
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
			if db.NewRecord(u) {
				log.Printf("ERROR: User Not Found, UserID=%d, FP=%s\n", k.UserID, fp)
				return false
			}
			// set User / SSHKey
			log.Printf("Signed In, UserID=%d, FP=%s\n", k.UserID, fp)
			ctx.SetValue("User", u)
			ctx.SetValue("SSHKey", k)
			return true
		}),
	}

	log.Printf("Listening at %s:%d\n", cfg.SSHD.Host, cfg.SSHD.Port)
	log.Fatal(sshd.ListenAndServe(fmt.Sprintf("%s:%d", cfg.SSHD.Host, cfg.SSHD.Port), nil, options...))
	return nil
}
