package dto

import (
	"github.com/gogf/gf/v2/database/gdb"
	"goframe-vben/internal/library/query"
)

// DictPageInput 字典分页
type DictPageInput struct {
	query.PageParam
}

func (q *DictPageInput) Cond(m *gdb.Model) *gdb.Model {
	return m
}

// DictCreateInput 字典新增
type DictCreateInput struct {
	Code string
	Name string
	Sort uint
	Note string
}

// DictUpdateInput 字典编辑
type DictUpdateInput struct {
	Id   int64
	Code string
	Name string
	Sort uint
	Note string
}
