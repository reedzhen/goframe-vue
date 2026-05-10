package admin

import (
	"context"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/menu"
)

func (c *ControllerMenu) GetList(ctx context.Context, req *menu.GetListReq) (res *menu.GetListRes, err error) {
	list, err := service.Menu().GetList(ctx, dto.MenuGetListInput{
		Title:     req.Title,
		Path:      req.Path,
		Authority: req.Authority,
		MenuType:  req.MenuType,
		ParentId:  req.ParentId,
	})
	if err != nil {
		return
	}

	temp := ConvertToTree(list)

	return (*menu.GetListRes)(&temp), nil
}
