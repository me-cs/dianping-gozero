// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package blog

import (
	"net/http"

	"github.com/me-cs/dianping-gozero/api/internal/logic/blog"
	"github.com/me-cs/dianping-gozero/api/internal/svc"
	"github.com/me-cs/dianping-gozero/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryBlogLikesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryBlogLikesReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := blog.NewQueryBlogLikesLogic(r.Context(), svcCtx)
		resp, err := l.QueryBlogLikes(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
