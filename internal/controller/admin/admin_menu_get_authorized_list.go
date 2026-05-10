package admin

import (
	"context"
	"goframe-vben/api/admin/menu"
	"goframe-vben/internal/service"
)

func (c *ControllerMenu) GetAuthorizedList(ctx context.Context, req *menu.GetAuthorizedListReq) (res *menu.GetAuthorizedListRes, err error) {
	list, err := service.Menu().GetAuthorizedList(ctx)
	if err != nil {
		return
	}

	newList := ConvertToTree(list)

	return (*menu.GetAuthorizedListRes)(&newList), nil
}
