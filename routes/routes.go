package routes

import (
	"github.com/pagoda-tech/bastion/routes/middlewares"
	"github.com/pagoda-tech/bastion/routes/token"
	"github.com/pagoda-tech/bastion/routes/user"
	"github.com/pagoda-tech/bastion/utils"
	"github.com/pagoda-tech/binding"
	"github.com/pagoda-tech/macaron"
)

// Mount 将所有路由挂载到 macaron 上
func Mount(m *macaron.Macaron) {
	m.Use(middlewares.Renderer())
	m.Get("/api", apiAction)
	m.Post("/api/tokens/create", binding.Bind(token.CreateForm{}), token.Create)
	m.Get("/api/users/:id", middlewares.Authenticate(), user.Show)
}

func apiAction(ctx *macaron.Context, r *middlewares.Render) {
	r.Success(func(m utils.Map) {
		m.Set("name", "bastion").Set("version", ctx.Data["Version"])
	})
}
