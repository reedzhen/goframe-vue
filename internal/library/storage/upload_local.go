package storage

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"path/filepath"
)

// LocalDrive 本地驱动
type LocalDrive struct {
}

// UploadFile 上传到本地
func (d *LocalDrive) UploadFile(ctx context.Context, in UploadParam) (url string, err error) {
	serverRoot := g.Cfg().MustGet(ctx, "server.serverRoot").String()
	baseUrl := g.Cfg().MustGet(ctx, `app.domain`).String()

	fileName, err := in.File.Save(filepath.Join(serverRoot, in.Dir), in.RandomName)
	if err != nil {
		return "", err
	}
	return filepath.Join(baseUrl, in.Dir, fileName), nil // 直接访问地址
}
