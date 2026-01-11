// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package upload

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/me-cs/dianping-gozero/api/internal/svc"
	"github.com/me-cs/dianping-gozero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadImageLogic {
	return &UploadImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadImageLogic) UploadImage(file multipart.File, fileHeader *multipart.FileHeader) (resp *types.UploadImageRsp, err error) {
	// 获取原始文件名
	originalFilename := fileHeader.Filename

	// 生成新文件名
	fileName, err := l.createNewFileName(originalFilename)
	if err != nil {
		return &types.UploadImageRsp{
			Success: false,
			ErrMsg:  "生成文件名失败: " + err.Error(),
		}, nil
	}

	// 构建完整的文件路径
	uploadDir := l.svcCtx.Config.Upload.Dir
	fullPath := filepath.Join(uploadDir, fileName)

	// 确保目录存在
	dir := filepath.Dir(fullPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return &types.UploadImageRsp{
			Success: false,
			ErrMsg:  "创建目录失败: " + err.Error(),
		}, nil
	}

	// 创建目标文件
	dst, err := os.Create(fullPath)
	if err != nil {
		return &types.UploadImageRsp{
			Success: false,
			ErrMsg:  "创建文件失败: " + err.Error(),
		}, nil
	}
	defer dst.Close()

	// 复制文件内容
	if _, err := io.Copy(dst, file); err != nil {
		return &types.UploadImageRsp{
			Success: false,
			ErrMsg:  "保存文件失败: " + err.Error(),
		}, nil
	}
	//docker  容器暂时不好在宿主机创建文件，
	fileName = ""
	l.Logger.Infof("文件上传成功: %s", fileName)

	return &types.UploadImageRsp{
		Success: true,
		Data:    fileName,
	}, nil
}

// createNewFileName 根据原始文件名生成新的文件名
// 格式: /blogs/{d1}/{d2}/{uuid}.{ext}
func (l *UploadImageLogic) createNewFileName(originalFilename string) (string, error) {
	// 获取文件扩展名
	ext := filepath.Ext(originalFilename)
	if ext == "" {
		return "", fmt.Errorf("文件没有扩展名")
	}
	// 去掉点号
	ext = strings.TrimPrefix(ext, ".")

	// 生成 UUID
	name := uuid.New().String()

	// 计算 hash，生成两级目录
	hash := hashCode(name)
	d1 := hash & 0xF
	d2 := (hash >> 4) & 0xF

	// 生成文件名: /blogs/{d1}/{d2}/{uuid}.{ext}
	fileName := fmt.Sprintf("/blogs/%x/%x/%s.%s", d1, d2, name, ext)

	return fileName, nil
}

// hashCode 计算字符串的 hash 值（模拟 Java 的 hashCode）
func hashCode(s string) int {
	h := 0
	for _, c := range s {
		h = 31*h + int(c)
	}
	return h
}
