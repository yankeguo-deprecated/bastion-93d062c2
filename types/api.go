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
