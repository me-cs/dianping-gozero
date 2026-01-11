// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package upload

import (
	"net/http"

	"github.com/me-cs/dianping-gozero/api/internal/logic/upload"
	"github.com/me-cs/dianping-gozero/api/internal/svc"
	"github.com/me-cs/dianping-gozero/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 解析 multipart form，限制最大 32MB
		if err := r.ParseMultipartForm(32 << 20); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		// 获取上传的文件
		file, fileHeader, err := r.FormFile("file")
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, &types.UploadImageRsp{
				Success: false,
				ErrMsg:  "获取上传文件失败: " + err.Error(),
			})
			return
		}
		defer file.Close()

		l := upload.NewUploadImageLogic(r.Context(), svcCtx)
		resp, err := l.UploadImage(file, fileHeader)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
