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
	m.Get( "/api/users/:userId/tokens", RequireAuth(), TokenList)
	m.Post("/api/tokens/:id/destroy", RequireAuth(), TokenDestroy)
	m.Get( "/api/users/:id", RequireAuth(), UserShow)
}

func apiAction(ctx *macaron.Context, r APIRender) {
	r.Success(utils.NewMap("name", "bastion", "version", ctx.Data["Version"]))
}
