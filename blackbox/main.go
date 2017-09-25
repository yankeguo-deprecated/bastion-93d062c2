package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"ireul.com/bastion/types"
	"ireul.com/etc/passwd"
	"ireul.com/yaml"
)

// ConfigFilePath path to config file
const ConfigFilePath = "/etc/blackbox.yaml"

// Config config struct
type Config struct {
	Host  string `yaml:"host"`
	Token string `yaml:"token"`
	Home  string `yaml:"home"`
}

var debug = len(os.Getenv("BLACKBOX_DEBUG")) > 0

func main() {
	// setup log
	log.SetPrefix("[blackbox] ")

	// check is root user
	if os.Getuid() != 0 && !debug {
		log.Fatalln("blackbox must run as root user!")
		return
	}

	// decode config
	cp := ConfigFilePath
	if debug {
		cp = "blackbox.yaml"
	}

	c := Config{}

	data, err := ioutil.ReadFile(cp)
	if err != nil {
		log.Fatalln("failed to read", cp, err.Error())
		return
	}

	if err = yaml.Unmarshal(data, &c); err != nil {
		log.Fatalln("failed to decode", cp, err.Error())
		return
	}

	// check config
	if len(c.Host) == 0 {
		log.Fatalln("'host' not set in", cp)
		return
	}
	if len(c.Token) == 0 {
		log.Fatalln("'token' not set in", cp)
		return
	}
	if len(c.Home) == 0 {
		c.Home = "/home"
	}

	syncAccounts(c)
}

func syncAccounts(c Config) {
	// request API
	res, err := requestAPI(c)
	if debug {
		fmt.Println(res)
	}
	if err != nil {
		log.Println("failed to request bastion host:", err.Error())
		return
	}

	// read data
	as, err := readSystemAccounts()
	if err != nil {
		log.Println("failed to read accounts from /etc/passwd:", err.Error())
		return
	}

	// compute
	data := SyncData{
		BaseDir:        c.Home,
		Accounts:       res.Accounts,
		AccountsAdd:    []string{},
		AccountsRemove: []string{},
	}

	// missing accounts
	for _, a := range res.Accounts {
		if as[a.Account].UID == "" {
			data.AccountsAdd = append(data.AccountsAdd, a.Account)
		}
	}

	// disabled accounts
OUTER:
	for k := range as {
		for _, a := range res.Accounts {
			if a.Account == k {
				continue OUTER
			}
		}
		data.AccountsRemove = append(data.AccountsRemove, k)
	}

	// execute script
	script := GenerateSyncScript(data)

	if debug {
		log.Println("--- DEBUG sync script output")
		log.Println(script)
		log.Println("--- DEBUG sync script output")
		return
	}

	cmd := exec.Command("/bin/bash")
	cmd.Stdin = strings.NewReader(script)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println("failed to execute sync shell:", err.Error())
		log.Println(string(out))
	}
}

func requestAPI(c Config) (res types.ServerSyncResponse, err error) {
	var req *http.Request
	// make request
	if req, err = http.NewRequest(http.MethodPost, c.Host+"/api/servers/sync", strings.NewReader("")); err != nil {
		return
	}
	req.Header.Set("Authorization", "Bearer "+c.Token)
	// execute request
	var r *http.Response
	if r, err = http.DefaultClient.Do(req); err != nil {
		return
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&res)
	return
}

func readSystemAccounts() (map[string]passwd.Entry, error) {
	var file = "/etc/passwd"
	if debug {
		file = "passwd"
	}
	all, err := passwd.ParseFile(file)
	if err != nil {
		return nil, err
	}
	// filter accounts not start with 'bastion-'
	ret := make(map[string]passwd.Entry)
	for k, v := range all {
		if strings.HasPrefix(k, types.AccountPrefix) {
			ret[k] = v
		}
	}
	return ret, nil
}
