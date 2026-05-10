package admin

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"
	"strings"

	"goframe-vben/api/admin/user"
)

func (c *ControllerUser) GetList(ctx context.Context, req *user.GetListReq) (res *user.GetListRes, err error) {
	list, err := service.User().GetList(ctx, dto.UserGetListInput{
		OrganizationId: req.OrganizationId,              // 部门Id
		Nickname:       strings.TrimSpace(req.Nickname), // 昵称
		Username:       strings.TrimSpace(req.Username), // 账号
	})
	if err != nil {
		return nil, err
	}

	// 案例 将数据库查询出来的数据转换成数组返回前端
	temp := make([]*user.UserItem, 0)
	if err = gconv.Struct(list, &temp); err != nil {
		return
	}
	return (*user.GetListRes)(&temp), nil
}
