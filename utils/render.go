package utils

import (
	"github.com/pagoda-tech/macaron"
)

// Render 封装后的 macaron.Render
type Render struct {
	macaron.Render
}

// Success 返回 code = 200，并构建 WildMap
func (r *Render) Success(builder func(m WildMap)) {
	m := Map("code", "OK")
	if builder != nil {
		builder(m)
	}
	r.JSON(200, m)
}

// Fail 返回 code = 400，并构建 WildMap
func (r *Render) Fail(code string, msg string) {
	m := Map("code", code).Set("message", msg)
	r.JSON(400, m)
}

// Renderer 为 macaron.Context 注入 utils.Render
func Renderer() interface{} {
	return func(r macaron.Render, ctx *macaron.Context) {
		render := &Render{r}
		ctx.Map(render)
	}
}
