package admin

import (
	"context"
	"goframe-vben/internal/consts"
	"goframe-vben/internal/library/storage"

	"goframe-vben/api/admin/common"
)

func (c *ControllerCommon) FileUpload(ctx context.Context, req *common.FileUploadReq) (res *common.FileUploadRes, err error) {
	url, err := storage.New(consts.UploadDriveOss).UploadFile(ctx, storage.UploadParam{
		Dir:        req.Dir,
		File:       req.File,
		RandomName: true,
	})
	if err != nil {
		return nil, err
	}
	return &common.FileUploadRes{Url: url}, nil
}
