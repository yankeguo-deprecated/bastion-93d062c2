package utils

import "testing"

func TestConfig(t *testing.T) {
	const c = `
[bastion]
env = "a"
sandbox_dir = "b"
master_key_file = "c"
authorized_keys_file = "d"
sandbox_image = "d1"
[db]
url = "e"
[redis]
url = "f"
[web]
domain = "g"
host = "h"
port = 9
	`
	cfg, err := ParseConfig(c)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Bastion.Env != "a" {
		t.Fatal("bastion.env failed")
	}
	if cfg.Bastion.SandboxDir != "b" {
		t.Fatal("bastion.sandbox_dir failed")
	}
	if cfg.Bastion.MasterKeyFile != "c" {
		t.Fatal("bastion.master_key_file failed")
	}
	if cfg.Bastion.AuthorizedKeysFile != "d" {
		t.Fatal("bastion.authorized_keys_file failed")
	}
	if cfg.Database.URL != "e" {
		t.Fatal("db.url failed")
	}
	if cfg.Redis.URL != "f" {
		t.Fatal("redis.url failed")
	}
	if cfg.Web.Domain != "g" {
		t.Fatal("web.domain failed")
	}
	if cfg.Web.Host != "h" {
		t.Fatal("web.host failed")
	}
	if cfg.Web.Port != 9 {
		t.Fatal("web.port failed")
	}
	if cfg.Bastion.SandboxImage != "d1" {
		t.Fatal("bastion.sandbox_image failed")
	}
}
