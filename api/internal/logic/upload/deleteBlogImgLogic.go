// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package upload

import (
	"context"
	"os"
	"path/filepath"

	"github.com/me-cs/dianping-gozero/api/internal/svc"
	"github.com/me-cs/dianping-gozero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteBlogImgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteBlogImgLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteBlogImgLogic {
	return &DeleteBlogImgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteBlogImgLogic) DeleteBlogImg(req *types.DeleteBlogImgReq) (resp *types.DeleteBlogImgRsp, err error) {
	// 构建文件路径
	uploadDir := l.svcCtx.Config.Upload.Dir
	filePath := filepath.Join(uploadDir, req.Name)

	// 检查文件信息
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return &types.DeleteBlogImgRsp{
				Success: false,
				ErrMsg:  "文件不存在",
			}, nil
		}
		return &types.DeleteBlogImgRsp{
			Success: false,
			ErrMsg:  "获取文件信息失败: " + err.Error(),
		}, nil
	}

	// 检查是否是目录
	if fileInfo.IsDir() {
		return &types.DeleteBlogImgRsp{
			Success: false,
			ErrMsg:  "错误的文件名称",
		}, nil
	}

	// 删除文件
	if err := os.Remove(filePath); err != nil {
		return &types.DeleteBlogImgRsp{
			Success: false,
			ErrMsg:  "删除文件失败: " + err.Error(),
		}, nil
	}

	l.Logger.Infof("文件删除成功: %s", req.Name)

	return &types.DeleteBlogImgRsp{
		Success: true,
	}, nil
}
