package routes

import (
	"ireul.com/com"
	"ireul.com/web"
)

// APIRender 封装后的 web.Render
type APIRender interface {
	web.Render
	Success(args ...interface{})
	Fail(code string, message string)
}

// apiRenderImpl 封装后的 web.Render
type apiRenderImpl struct {
	web.Render
}

// Success 返回 code = 200，并构建 Map
func (r *apiRenderImpl) Success(args ...interface{}) {
	var m com.Map
	if len(args) == 1 {
		a := args[0]
		if v, ok := a.(com.Map); ok {
			m = v
		} else if v, ok := a.(map[string]interface{}); ok {
			m = com.Map(v)
		} else if v, ok := a.(func(com.Map)); ok {
			m = com.Map{}
			v(m)
		} else {
			m = com.Map{}
		}
	} else if len(args) > 0 && len(args)%2 == 0 {
		m = com.NewMap(args...)
	} else {
		m = com.Map{}
	}
	m.Set("code", "ok")
	r.JSON(200, m)
}

// Fail 返回 code = 400，并构建 Map
func (r *apiRenderImpl) Fail(code string, msg string) {
	r.JSON(400, com.NewMap("code", code, "message", msg))
}

// APIRenderer 为 web.Context 注入 com.Render
func APIRenderer() interface{} {
	return func(r web.Render, ctx *web.Context) {
		render := &apiRenderImpl{r}
		ctx.MapTo(render, (*APIRender)(nil))
	}
}
