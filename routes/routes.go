package routes

import (
	"github.com/pagoda-tech/bastion/routes/accesstoken"
	"github.com/pagoda-tech/binding"
	"github.com/pagoda-tech/macaron"
)

// Mount 将所有路由挂载到 macaron 上
func Mount(m *macaron.Macaron) {
	m.Get("/", func(ctx *macaron.Context) {
		ctx.Render.HTML(200, "index", nil)
	})
	m.Get("/api", func(ctx *macaron.Context) {
		ctx.Render.JSON(200, map[string]interface{}{
			"code": "OK",
		})
	})
	m.Post("/api/access_tokens/create", binding.Bind(accesstoken.CreateForm{}), accesstoken.CreateAction)
}
