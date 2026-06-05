package dict

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-vben/internal/library/query"
	"goframe-vben/internal/model/entity"
)

type GetPageReq struct {
	g.Meta `path:"/dict/page" method:"get" tags:"字典管理" summary:"字典分页" noAuth:"true"`
	query.PageParam
}
type GetPageRes struct {
	*query.Result
}

type CreateReq struct {
	g.Meta `path:"/dict/create" method:"post" tags:"字典管理" summary:"新增字典"`
	Code   string `json:"code" v:"required" dc:"字典标识"`
	Name   string `json:"name" v:"required" dc:"字典名称"`
	Note   string `json:"note" dc:"备注"`
}
type CreateRes struct{}

type UpdateReq struct {
	g.Meta `path:"/dict/update" method:"post" tags:"字典管理" summary:"编辑字典"`
	Id     int64  `json:"id"  v:"required" dc:"主键ID"`
	Code   string `json:"code" v:"required" dc:"字典标识"`
	Name   string `json:"name" v:"required" dc:"字典名称"`
	Note   string `json:"note" dc:"备注"`
}
type UpdateRes struct{}

type DeleteReq struct {
	g.Meta `path:"/dict/delete/{Id}" method:"post" tags:"字典管理" summary:"删除字典"`
	Id     int64 `json:"id" v:"required"`
}
type DeleteRes struct{}

type GetListReq struct {
	g.Meta `path:"/dict/list" method:"get" tags:"字典管理" summary:"字典列表" noAuth:"true"`
}
type GetListRes []*entity.SysDict
