package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"ireul.com/bastion/types"
	"ireul.com/etc/passwd"
)

var host = os.Getenv("BASTION_HOST")
var token = os.Getenv("BASTION_TOKEN")
var home = os.Getenv("BASTION_HOME")
var ticker = time.NewTicker(time.Minute)

var debug = len(os.Getenv("BASTION_DEBUG")) > 0

func main() {
	// check is root
	if os.Getuid() != 0 && !debug {
		log.Fatalln("bagent must run as root user!")
		return
	}
	// check variables
	if len(host) == 0 {
		log.Fatalln("environment 'BASTION_HOST' not set!")
		return
	}
	if len(token) == 0 {
		log.Fatalln("environment 'BASTION_TOKEN' not set!")
		return
	}
	if len(home) == 0 {
		home = "/home"
	}
	// start loop
	syncAccounts()
	for {
		<-ticker.C
		syncAccounts()
	}
}

func syncAccounts() {
	// request API
	res, err := makeRequest()
	if debug {
		fmt.Println(res)
	}
	if err != nil {
		log.Println("failed to request bastion host:", err.Error())
		return
	}

	// read data
	as, err := readAccounts()
	if err != nil {
		log.Println("failed to read accounts from /etc/passwd:", err.Error())
		return
	}

	// compute
	data := SyncData{
		BaseDir:  home,
		Accounts: res.Accounts,
	}

	adds := []string{}
	dels := []string{}

	// missing accounts
	for _, a := range res.Accounts {
		if as[a.Account].UID == "" {
			adds = append(adds, a.Account)
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
		dels = append(dels, k)
	}

	data.AccountsAdd = adds
	data.AccountsRemove = dels

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

func makeRequest() (res types.ServerSyncResponse, err error) {
	var req *http.Request
	// make request
	if req, err = http.NewRequest(http.MethodPost, host+"/api/servers/sync", strings.NewReader("")); err != nil {
		return
	}
	req.Header.Set("Authorization", "Bearer "+token)
	// execute request
	var r *http.Response
	if r, err = http.DefaultClient.Do(req); err != nil {
		return
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(&res)
	return
}

func readAccounts() (map[string]passwd.Entry, error) {
	var file = "/etc/passwd"
	if debug {
		file = "passwd"
	}
	all, err := passwd.ParseFile(file)
	if err != nil {
		return nil, err
	}
	ret := make(map[string]passwd.Entry)
	for k, v := range all {
		if strings.HasPrefix(k, types.AccountPrefix) {
			ret[k] = v
		}
	}
	return ret, nil
}
