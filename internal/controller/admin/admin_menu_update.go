package admin

import (
	"context"
	"goframe-vben/internal/library/contexts"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/menu"
)

func (c *ControllerMenu) Update(ctx context.Context, req *menu.UpdateReq) (res *menu.UpdateRes, err error) {
	return nil, service.Menu().Update(ctx, dto.MenuUpdateInput{
		Id:        req.Id,
		UpdatedBy: contexts.GetUserId(ctx),
		MenuCreateUpdateBase: dto.MenuCreateUpdateBase{
			ParentId:  req.ParentId,
			Title:     req.Title,
			Path:      req.Path,
			Component: req.Component,
			MenuType:  req.MenuType,
			Sort:      req.Sort,
			Authority: req.Authority,
			Icon:      req.Icon,
			Hide:      req.Hide,
			MenuMeta:  req.MenuMeta,
			ApiPath:   req.ApiPath,
		},
	})
}
