package admin

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/service"

	"goframe-vben/api/admin/dict"
)

func (c *ControllerDict) DataGetList(ctx context.Context, req *dict.DataGetListReq) (res *dict.DataGetListRes, err error) {
	list, err := service.DictData().GetList(ctx, dto.DictDataGetListInput{DictCode: req.DictCode})
	if err != nil {
		return
	}

	retList := make([]*dict.DataGetListItem, len(list))
	if err = gconv.Structs(list, &retList); err != nil {
		return
	}

	// 转换数据类型
	for _, v := range retList {
		switch v.CodeType {
		case "string":
			v.Code = gconv.String(v.Code)
		case "int":
			v.Code = gconv.Int(v.Code)
		}
	}

	res = (*dict.DataGetListRes)(&retList)
	return
}
