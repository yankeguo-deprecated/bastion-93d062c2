package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"ireul.com/bastion/types"
	"ireul.com/passwd"
)

// Account is something hold a account information
type Account struct {
	Login string
	Sudo  bool
}

var host = os.Getenv("BASTION_HOST")
var token = os.Getenv("BASTION_TOKEN")
var home = os.Getenv("BASTION_HOME")
var ticker = time.NewTicker(time.Minute)

var debug = len(os.Getenv("BASTION_DEBUG")) > 0

func main() {
	// check is root
	if os.Getuid() != 0 {
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
		log.Fatalln("environment 'BASTION_HOME' not set!")
		return
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
	fmt.Println(res)
	if err != nil {
		log.Println("failed to request bastion host: ", err.Error())
		return
	}
	// read /etc/passwd
	entries, err := passwd.ParseFile("passwd")
	if err != nil {
		log.Println("failed to parse /etc/passwd, ", err.Error())
		return
	}
	// prepare script
	script := ""

	// ban not existed users
	for name, e := range entries {
		if strings.HasPrefix(name, types.AccountPrefix) {
			fmt.Println(script, e)
		}
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
