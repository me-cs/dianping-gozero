// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/me-cs/dianping-gozero/api/internal/logic/user"
	"github.com/me-cs/dianping-gozero/api/internal/svc"
	"github.com/me-cs/dianping-gozero/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetMeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetMeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 从header获取token
		token := r.Header.Get("authorization")
		if token == "" {
			httpx.WriteJson(w, http.StatusUnauthorized, types.GetMeRsp{
				Success: false,
				ErrMsg:  "未登录",
			})
			return
		}

		// 从Redis获取用户信息
		userInfoStr, err := svcCtx.Dao.Redis.GetCtx(r.Context(), "login:token:"+token)
		if err != nil {
			httpx.WriteJson(w, http.StatusUnauthorized, types.GetMeRsp{
				Success: false,
				ErrMsg:  "登录已过期",
			})
			return
		}

		// 解析用户信息获取user_id
		var userInfo map[string]interface{}
		if err := json.Unmarshal([]byte(userInfoStr), &userInfo); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		userId, ok := userInfo["id"].(float64)
		if !ok {
			httpx.WriteJson(w, http.StatusInternalServerError, types.GetMeRsp{
				Success: false,
				ErrMsg:  "用户信息格式错误",
			})
			return
		}

		// 将userId放入context
		ctx := context.WithValue(r.Context(), "userId", int64(userId))

		l := user.NewGetMeLogic(ctx, svcCtx)
		resp, err := l.GetMe(&req)
		if err != nil {
			httpx.ErrorCtx(ctx, w, err)
		} else {
			httpx.OkJsonCtx(ctx, w, resp)
		}
	}
}
