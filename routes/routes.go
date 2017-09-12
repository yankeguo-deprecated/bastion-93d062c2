package routes

import (
	"github.com/pagoda-tech/bastion/routes/token"
	"github.com/pagoda-tech/bastion/utils"
	"github.com/pagoda-tech/binding"
	"github.com/pagoda-tech/macaron"
)

// Mount 将所有路由挂载到 macaron 上
func Mount(m *macaron.Macaron) {
	m.Get("/api", apiAction)
	m.Post("/api/tokens/create", binding.Bind(token.CreateForm{}), token.CreateAction)
}

func apiAction(ctx *macaron.Context, r *utils.Render) {
	r.Success(func(m utils.WildMap) {
		m.Set("name", "bastion").Set("version", ctx.Data["Version"])
	})
}
