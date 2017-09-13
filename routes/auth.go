package routes

import (
	"github.com/pagoda-tech/bastion/models"
	"github.com/pagoda-tech/macaron"
	"strings"
)

// Auth 认证结果
type Auth struct {
	Code         string
	Message      string
	CurrentToken *models.Token
	CurrentUser  *models.User
}

// SignedIn 是否已经登录
func (a Auth) SignedIn() bool {
	return a.CurrentToken != nil && a.CurrentUser != nil
}

// Authenticator 创建认证中间件
func Authenticator() interface{} {
	return func(ctx *macaron.Context, db *models.DB, r APIRender) {
		a := Auth{}

		k := extractBearer(ctx.Req)

		if len(k) > 0 {
			// find a Token
			t := &models.Token{}
			db.Where("secret = ?", k).Find(t)

			if t.ID > 0 {
				// find a User
				u := &models.User{}
				db.First(u, t.UserID)

				if u.ID > 0 && !u.IsBlocked {
					// touch token and user
					db.Touch(t)
					db.Touch(u)

					// assign user / token
					a.CurrentUser = u
					a.CurrentToken = t
				}
			}

			if a.CurrentUser == nil || a.CurrentToken == nil {
				a.Code = CredentialsInvalid
				a.Message = "无效的凭证"
			}
		} else {
			a.Code = CredentialsMissing
			a.Message = "没有凭证"
		}

		ctx.Map(a)
	}
}

// RequireAuth 检验认证结果
func RequireAuth() interface{} {
	return func(ctx *macaron.Context, a Auth, r APIRender) {
		if !a.SignedIn() {
			r.Fail(a.Code, a.Message)
		}
	}
}

func extractBearer(req macaron.Request) (k string) {
	h := req.Header["Authorization"]
	if h != nil && len(h) > 0 {
		vs := strings.Split(strings.TrimSpace(h[len(h)-1]), " ")
		if len(vs) == 2 && vs[0] == "Bearer" {
			k = vs[1]
		}
	}
	return
}
