package dict

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-vben/internal/library/query"
)

type DataGetPageReq struct {
	g.Meta `path:"/dict-data/page" method:"get" tags:"字典管理" summary:"获取字典项分页" noAuth:"true"`
	query.PageParam
	DictId   uint   `json:"dictId" in:"query" dc:"dict_id"`
	Keywords string `json:"keywords" in:"query" dc:"关键字"`
}
type DataGetPageRes struct {
	*query.Result
}

type DataCreateReq struct {
	g.Meta `path:"/dict-data/create" method:"post" tags:"字典管理" summary:"新增字典项" noAuth:"true"`
	DictId int64  `json:"dictId" v:"required" dc:"字典id"`
	Code   string `json:"code" v:"required" dc:"字典项标识"`
	Name   string `json:"name" v:"required" dc:"字典项名称"`
	Sort   int    `json:"sort" v:"required" dc:"排序号"`
	Note   string `json:"note"  dc:"备注"`
}
type DataCreateRes struct{}

type DataUpdateReq struct {
	g.Meta `path:"/dict-data/update" method:"post" tags:"字典管理" summary:"编辑字典项" noAuth:"true"`
	Id     int64  `json:"id"  v:"required" dc:"主键ID"`
	DictId int64  `json:"dictId" v:"required" dc:"字典id"`
	Code   string `json:"code" v:"required" dc:"字典项标识"`
	Name   string `json:"name" v:"required" dc:"字典项名称"`
	Sort   int    `json:"sort" v:"required" dc:"排序号"`
	Note   string `json:"note"  dc:"备注"`
}
type DataUpdateRes struct{}

type DataDeleteReq struct {
	g.Meta `path:"/dict-data/delete/{Id}" method:"post" tags:"字典管理" summary:"删除字典项" noAuth:"true"`
	Id     int64 `json:"id" v:"required"`
}
type DataDeleteRes struct{}

type DataGetListReq struct {
	g.Meta   `path:"/dict-data/list" method:"get" tags:"字典管理" summary:"获取字典项列表" noAuth:"true"`
	DictCode string `json:"dictCode" v:"required"`
}
type DataGetListRes []*DataGetListItem

type DataGetListItem struct {
	Id       int64  `json:"id"       dc:"字典项id" `
	DictId   int64  `json:"dictId"    dc:"字典id"    `
	Code     any    `json:"code"      dc:"字典项标识"       `
	CodeType string `json:"codeType"  dc:"类型 string/int"  `
	Name     string `json:"name"      dc:"字典项名称"       `
	Sort     int    `json:"sort"      dc:"排序号"       `
	Note     string `json:"note"      dc:"备注"       `
}
