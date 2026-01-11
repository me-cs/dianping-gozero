package middleware

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/me-cs/dianping-gozero/library/auth"
	"github.com/me-cs/dianping-gozero/library/metadatax"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"google.golang.org/grpc/metadata"
)

const (
	// LOGIN_USER_KEY Redis中用户token的key前缀
	LOGIN_USER_KEY = "login:token:"
	// LOGIN_USER_TTL token有效期（30分钟）
	LOGIN_USER_TTL = 30 * time.Minute
)

// RefreshTokenMiddleware Token刷新中间件
// 从请求头获取token，如果存在则从Redis获取用户信息并保存到context
// 刷新token有效期，总是放行
type RefreshTokenMiddleware struct {
	redisClient *redis.Redis // Redis客户端，需要从ServiceContext注入
}

func NewRefreshTokenMiddleware(redisClient *redis.Redis) *RefreshTokenMiddleware {
	return &RefreshTokenMiddleware{
		redisClient: redisClient,
	}
}

func (m *RefreshTokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. 获取请求头中的token
		token := r.Header.Get("authorization")
		if token == "" {
			token = r.Header.Get("Authorization")
		}

		// 去除 "Bearer " 前缀（如果有）
		token = strings.TrimPrefix(token, "Bearer ")
		token = strings.TrimSpace(token)

		// 2. 如果token为空，直接放行
		if token == "" {
			logx.Infof(token)
			next(w, r)
			return
		}

		// 3. 基于TOKEN获取redis中的用户
		key := LOGIN_USER_KEY + token
		userJson, err := m.redisClient.Get(key)
		if err != nil {
			logx.Errorf("Failed to get user from redis: %v", err)
			// 即使Redis出错也放行，让后续的AuthMiddleware处理
			next(w, r)
			return
		}

		// 4. 判断用户是否存在
		if len(userJson) == 0 {
			logx.Infof("用户不存在")
			// 用户不存在，直接放行（token可能已过期）
			next(w, r)
			return
		}
		userInfo := map[string]interface{}{}
		err = json.Unmarshal([]byte(userJson), &userInfo)
		if err != nil {
			logx.Errorf("json.Unmarshal failed %v", err)
			// 即使Redis出错也放行，让后续的AuthMiddleware处理
			next(w, r)
			return
		}
		ctx := r.Context()

		userId, ok := userInfo["id"]
		if !ok {
			logx.Errorf("id不存在")
			next(w, r)
			return
		}
		userIdF64, ok := userId.(float64)
		if !ok {
			logx.Errorf("id类型不对")
			next(w, r)
			return
		}
		userIdI64 := int64(userIdF64)
		userStr, err := auth.EncodeUser(auth.NewUser(userIdI64))
		if err != nil {
			logx.Errorf("auth.EncodeUser,err=%v", err)
			next(w, r)
			return
		}
		md := metadata.Pairs(metadatax.UserInner, userStr)
		ctx = metadata.NewIncomingContext(ctx, md)

		// 6. 刷新token有效期
		err = m.redisClient.Expire(key, int(LOGIN_USER_TTL.Seconds()))
		if err != nil {
			logx.Errorf("Failed to refresh token TTL: %v", err)
			// 刷新失败不影响本次请求，继续放行

		}

		// 7. 放行
		next(w, r.WithContext(ctx))
	}
}
