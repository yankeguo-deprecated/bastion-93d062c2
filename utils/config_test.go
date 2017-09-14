package utils

import "testing"

func TestConfig(t *testing.T) {
	const c = `
[bastion]
env = "a"
master_key_file = "c"
authorized_keys_file = "d"
[db]
url = "e"
[redis]
url = "f"
[web]
domain = "g"
host = "h"
port = 9
[sandbox]
data_dir = "b"
image = "d1"
`
	cfg, err := ParseConfig(c)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.Bastion.Env != "a" {
		t.Fatal("bastion.env failed")
	}
	if cfg.Sandbox.DataDir != "b" {
		t.Fatal("sandbox.data_dir failed")
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
	if cfg.Sandbox.Image != "d1" {
		t.Fatal("sandbox.image failed")
	}
}
