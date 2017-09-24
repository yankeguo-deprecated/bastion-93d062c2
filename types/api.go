package types

// Account represents a linux local account on remote server, the name is not actual account name
type Account struct {
	Account   string
	CanSudo   bool
	PublicKey string
}

// ServerSyncResponse is response for servers/sync
type ServerSyncResponse struct {
	Accounts []Account
}

// Dashboard response
type Dashboard struct {
	Sandbox DashboardSandbox  `json:"sandbox"`
	Servers []DashboardServer `json:"servers"`
}

// DashboardSandbox sandbox information
type DashboardSandbox struct {
	IsKeyMissing bool   `json:"isKeyMissing"`
	Address      string `json:"address"`
}

// DashboardServer server in Dashboard API
type DashboardServer struct {
	ID      uint     `json:"id"`
	Name    string   `json:"name"`
	Account string   `json:"account"`
	Address string   `json:"address"`
	Port    uint     `json:"port"`
	CanSudo bool     `json:"canSudo"`
	Tags    []string `json:"tags"`
}
