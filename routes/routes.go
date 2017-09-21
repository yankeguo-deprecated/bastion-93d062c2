package routes

import (
	"ireul.com/com"
	"ireul.com/web"
	"ireul.com/web/binding"
)

// Mount 将所有路由挂载到 web 上
func Mount(m *web.Web) {
	m.Use(APIRenderer())
	m.Use(Authenticator())
	m.Get("/api", apiAction)
	m.Post("/api/tokens/create", binding.Bind(TokenCreateForm{}), TokenCreate)
	m.Post("/api/tokens/:id/destroy", RequireAuth(), ResolveCurrentToken(":id"), TokenDestroy)
	m.Get("/api/users/:userid/tokens", RequireAuth(), ResolveCurrentUser(":userid"), TokenList)
	m.Get("/api/users/:userid/ssh_keys", RequireAuth(), ResolveCurrentUser(":userid"), SSHKeyList)
	m.Post("/api/users/:userid/ssh_keys/create", binding.Bind(SSHKeyCreateForm{}), RequireAuth(), ResolveCurrentUser(":userid"), SSHKeyCreate)
	m.Get("/api/users", RequireAdmin(), UserList)
	m.Post("/api/users/create", RequireAdmin(), binding.Bind(UserCreateForm{}), UserCreate)
	m.Get("/api/users/:id", RequireAuth(), ResolveCurrentUser(":id"), UserShow)
	m.Post("/api/users/:id/update", RequireAuth(), ResolveCurrentUser(":id"), binding.Bind(UserUpdateForm{}), UserUpdate)
	m.Post("/api/users/:id/update_password", RequireAuth(), ResolveCurrentUser(":id"), binding.Bind(UserUpdatePasswordForm{}), UserUpdatePassword)
	m.Post("/api/users/:id/update_authority", RequireAdmin(), binding.Bind(UserUpdateAuthorityForm{}), UserUpdateAuthority)
	m.Get("/api/users/:id/audit_logs", RequireAuth(), ResolveCurrentUser(":id"), AuditLogsListByUser)
	m.Get("/api/servers", RequireAdmin(), ServerList)
	m.Post("/api/servers/sync", RequireServerToken(), ServerSync)
	m.Post("/api/servers/create", RequireAdmin(), binding.Bind(ServerCreateForm{}), ServerCreate)
	m.Post("/api/servers/:id/update", RequireAdmin(), binding.Bind(ServerUpdateForm{}), ServerUpdate)
	m.Post("/api/servers/:id/destroy", RequireAdmin(), ServerDestroy)
	m.Post("/api/ssh_keys/:id/destroy", RequireAuth(), SSHKeyDestroy)
	m.Get("/api/tags/:tag/grants", RequireAdmin(), GrantList)
	m.Post("/api/grants/create", RequireAdmin(), binding.Bind(GrantCreateForm{}), GrantCreate)
	m.Post("/api/grants/:id/destroy", RequireAdmin(), GrantDestroy)
	m.Get("/api/tags", RequireAdmin(), TagList)
}

func apiAction(ctx *web.Context, r APIRender) {
	r.Success(com.NewMap(
		"name",
		"bastion",
		"version",
		ctx.Data["Version"],
	))
}
