package admin

import (
	"context"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/dict"
)

func (c *ControllerDict) GetList(ctx context.Context, req *dict.GetListReq) (res *dict.GetListRes, err error) {
	list, err := service.Dict().GetList(ctx)
	if err != nil {
		return
	}

	return (*dict.GetListRes)(&list), nil
}
