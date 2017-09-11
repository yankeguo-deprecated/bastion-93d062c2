package utils

import (
	"fmt"
	"gopkg.in/macaron.v1"
	"log"
	"net/http"
	"reflect"
	"time"
)

var (
	// LogTimeFormat 日志日期格式
	LogTimeFormat = "2006-01-02 15:04:05"
)

// LoggerInvoker is an inject.FastInvoker wrapper of func(ctx *Context, log *log.Logger).
type LoggerInvoker func(ctx *macaron.Context, log *log.Logger)

// Invoke 接口实现
func (invoke LoggerInvoker) Invoke(params []interface{}) ([]reflect.Value, error) {
	invoke(params[0].(*macaron.Context), params[1].(*log.Logger))
	return nil, nil
}

// Logger returns a middleware handler that logs the request as it goes in and the response as it goes out.
func Logger() macaron.Handler {
	return func(ctx *macaron.Context, log *log.Logger) {
		start := time.Now()

		log.Printf("%s: Started %s %s for %s", time.Now().Format(LogTimeFormat), ctx.Req.Method, ctx.Req.RequestURI, ctx.RemoteAddr())

		rw := ctx.Resp.(macaron.ResponseWriter)
		ctx.Next()

		content := fmt.Sprintf("%s: Completed %s %s %v %s in %v", time.Now().Format(LogTimeFormat), ctx.Req.Method, ctx.Req.RequestURI, rw.Status(), http.StatusText(rw.Status()), time.Since(start))
		log.Println(content)
	}
}
