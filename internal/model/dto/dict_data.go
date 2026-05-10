package dto

import (
	"github.com/gogf/gf/v2/database/gdb"
	"goframe-vben/internal/library/query"
	"strings"
)

// DictDataPageInput 字典项分页
type DictDataPageInput struct {
	query.PageParam
	DictId   uint
	Keywords string
}

func (q *DictDataPageInput) Cond(m *gdb.Model) *gdb.Model {
	if q.DictId > 0 {
		m = m.Where("dict_id", q.DictId)
	}
	if q.Keywords != "" {
		m = m.WhereLike("name", "%"+strings.TrimSpace(q.Keywords)+"%")
	}
	return m
}

// DictDataCreateInput 字典项新增
type DictDataCreateInput struct {
	DictId int64
	Code   string
	Name   string
	Sort   int
	Note   string
}

// DictDataUpdateInput 字典项编辑
type DictDataUpdateInput struct {
	Id     int64
	DictId int64
	Code   string
	Name   string
	Sort   int
	Note   string
}

// DictDataGetListInput 获取字典项列表
type DictDataGetListInput struct {
	DictCode string
}

//
//// DictDataGetListOutput 字典项列表返回
//type DictDataGetListOutput struct {
//	entity.SysDictData
//}
