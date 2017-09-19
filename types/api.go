package types

// ServerAccount represents a linux local account on remote server, the name is not actual account name
type ServerAccount struct {
	Account   string
	CanSudo   bool
	PublicKey string
}

// ServerSyncResponse is response for servers/sync
type ServerSyncResponse struct {
	Accounts []ServerAccount
}
