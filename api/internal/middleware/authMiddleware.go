package middleware

import (
	"net/http"

	"github.com/me-cs/dianping-gozero/library/rpcx"
	"github.com/me-cs/dianping-gozero/library/rpcx/middlewares"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// AuthMiddleware 登录校验中间件
// 检查context中是否有用户信息（userId），如果没有则返回401
// 必须在RefreshTokenMiddleware之后执行
type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := rpcx.UserFromIncomingContext(r.Context())
		if user == nil {
			logx.Errorf("用户信息不在元数据里")
			httpx.Error(w, middlewares.Unauthorized)
			return
		}
		next(w, r)
	}
}
