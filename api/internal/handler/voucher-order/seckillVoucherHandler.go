// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package voucher-order

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/me-cs/dianping-gozero/api/internal/logic/voucher-order"
	"github.com/me-cs/dianping-gozero/api/internal/svc"
	"github.com/me-cs/dianping-gozero/api/internal/types"

)


func SeckillVoucherHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SeckillVoucherReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := voucher-order.NewSeckillVoucherLogic(r.Context(), svcCtx)
		resp, err := l.SeckillVoucher(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
