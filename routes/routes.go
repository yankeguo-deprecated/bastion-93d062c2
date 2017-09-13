package routes

import (
	"github.com/pagoda-tech/bastion/utils"
	"github.com/pagoda-tech/binding"
	"github.com/pagoda-tech/macaron"
)

// Mount 将所有路由挂载到 macaron 上
func Mount(m *macaron.Macaron) {
	m.Use(APIRenderer())
	m.Use(Authenticator())
	m.Get( "/api", apiAction)
	m.Post("/api/tokens/create", binding.Bind(TokenCreateForm{}), TokenCreate)
	m.Post("/api/tokens/:id/destroy", RequireAuth(), ResolveCurrentToken(":id"), TokenDestroy)
	m.Get( "/api/users/:userid/tokens", RequireAuth(), ResolveCurrentUser(":userid"), TokenList)
	m.Get( "/api/users/:userid/ssh_keys", RequireAuth(), ResolveCurrentUser(":userid"), SSHKeyList)
	m.Post("/api/users/:userid/ssh_keys/create", binding.Bind(SSHKeyCreateForm{}), RequireAuth(), ResolveCurrentUser(":userid"), SSHKeyCreate)
	m.Get( "/api/users/:id", RequireAuth(), ResolveCurrentUser(":id"), UserShow)
	m.Post("/api/ssh_keys/:id/destroy", RequireAuth(), SSHKeyDestroy)
}

func apiAction(ctx *macaron.Context, r APIRender) {
	r.Success(utils.NewMap("name", "bastion", "version", ctx.Data["Version"]))
}
