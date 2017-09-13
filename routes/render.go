package routes

import (
	"github.com/pagoda-tech/bastion/utils"
	"github.com/pagoda-tech/macaron"
)

// APIRender 封装后的 macaron.Render
type APIRender interface {
	macaron.Render
	Success(args ...interface{})
	Fail(code string, message string)
}

// apiRenderImpl 封装后的 macaron.Render
type apiRenderImpl struct {
	macaron.Render
}

// Success 返回 code = 200，并构建 Map
func (r *apiRenderImpl) Success(args ...interface{}) {
	var m utils.Map
	if len(args) == 1 {
		a := args[0]
		if v, ok := a.(utils.Map); ok {
			m = v
		} else if v, ok := a.(map[string]interface{}); ok {
			m = utils.Map(v)
		} else if v, ok := a.(func(utils.Map)); ok {
			m = utils.Map{}
			v(m)
		} else {
			m = utils.Map{}
		}
	} else if len(args) > 0 && len(args)%2 == 0 {
		m = utils.NewMap(args...)
	} else {
		m = utils.Map{}
	}
	m.Set("code", "ok")
	r.JSON(200, m)
}

// Fail 返回 code = 400，并构建 Map
func (r *apiRenderImpl) Fail(code string, msg string) {
	r.JSON(400, utils.NewMap("code", code, "message", msg))
}

// APIRenderer 为 macaron.Context 注入 utils.Render
func APIRenderer() interface{} {
	return func(r macaron.Render, ctx *macaron.Context) {
		render := &apiRenderImpl{r}
		ctx.MapTo(render, (*APIRender)(nil))
	}
}
