package admin

import (
	"context"
	"goframe-vben/internal/library/contexts"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/menu"
)

func (c *ControllerMenu) Create(ctx context.Context, req *menu.CreateReq) (res *menu.CreateRes, err error) {
	_, err = service.Menu().Create(ctx, dto.MenuCreateInput{
		CreatedBy: contexts.GetUserId(ctx),
		MenuCreateUpdateBase: dto.MenuCreateUpdateBase{
			Component: req.Component,
			MenuMeta:  req.MenuMeta,
			Title:     req.Title,
			Path:      req.Path,
			ParentId:  req.ParentId,
			Hide:      req.Hide,
			MenuType:  req.MenuType,
			Sort:      req.Sort,
			Authority: req.Authority,
			Icon:      req.Icon,
			ApiPath:   req.ApiPath,
		},
	})
	return nil, err
}
