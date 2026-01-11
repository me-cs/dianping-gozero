// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package shop-type

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/me-cs/dianping-gozero/api/internal/logic/shop-type"
	"github.com/me-cs/dianping-gozero/api/internal/svc"
	"github.com/me-cs/dianping-gozero/api/internal/types"

)


func QueryTypeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryTypeListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := shop-type.NewQueryTypeListLogic(r.Context(), svcCtx)
		resp, err := l.QueryTypeList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
