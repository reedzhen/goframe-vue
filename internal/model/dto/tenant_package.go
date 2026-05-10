package dto

import (
	"github.com/gogf/gf/v2/database/gdb"
	"goframe-vben/internal/library/query"
)

// TenantPackagePageInput 租户套餐分页
type TenantPackagePageInput struct {
	query.PageParam

	Name   string // 套餐名
	Status int    // 租户状态 1正常/2停用
}

func (q *TenantPackagePageInput) Cond(m *gdb.Model) *gdb.Model {
	if q.Name != "" {
		m = m.Where("name", q.Name)
	}
	if q.Status > 0 {
		m = m.Where("status", q.Status)
	}
	return m
}

// TenantPackageGetListInput 租户套餐列表
type TenantPackageGetListInput struct {
	Name   string // 套餐名
	Status int    // 租户状态 1正常/2停用
}

// TenantPackageCreateUpdateBase 租户套餐新建或修改
type TenantPackageCreateUpdateBase struct {
	Name    string // 套餐名
	Status  int    // 租户状态 1正常/2停用
	MenuIds string // 菜单ID 逗号分隔
	Remark  string // 备注
}

// TenantPackageCreateInput 租户套餐新建
type TenantPackageCreateInput struct {
	CreatedBy int64 // 创建人
	TenantPackageCreateUpdateBase
}

// TenantPackageUpdateInput 租户套餐修改
type TenantPackageUpdateInput struct {
	Id        int64
	UpdatedBy int64 // 更新人
	TenantPackageCreateUpdateBase
}
