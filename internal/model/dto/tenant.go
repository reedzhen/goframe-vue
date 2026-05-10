package dto

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"goframe-vben/internal/library/query"
)

// TenantPageInput 租户分页
type TenantPageInput struct {
	query.PageParam
	TenantName string // 租户名
	LinkMan    string // 联系人姓名
	LinkPhone  string // 联系人手机
	Status     int    // 租户状态 1正常/2停用
}

func (q *TenantPageInput) Cond(m *gdb.Model) *gdb.Model {
	if q.TenantName != "" {
		m = m.Where("tenant_name", q.TenantName)
	}
	if q.LinkMan != "" {
		m = m.Where("link_man", q.LinkMan)
	}
	if q.LinkPhone != "" {
		m = m.Where("link_phone", q.LinkPhone)
	}
	if q.Status > 0 {
		m = m.Where("status", q.Status)
	}
	return m
}

// TenantGetListInput 租户列表
type TenantGetListInput struct {
	TenantName string // 租户名
	LinkMan    string // 联系人姓名
	LinkPhone  string // 联系人手机
	PackageId  int64  // 租户套餐ID
}

// TenantCreateUpdateBase 租户新增或编辑
type TenantCreateUpdateBase struct {
	TenantName string // 租户名
	//PackageId    int64       // 租户套餐编号
	LinkMan string // 联系人姓名
	//LinkPhone    string      // 联系人手机
	Status       int         // 租户状态 1正常/2停用
	Website      string      // 绑定域名
	ExpireAt     *gtime.Time // 过期时间
	AccountCount int         // 账号数量
}

// TenantCreateInput 租户新增
type TenantCreateInput struct {
	CreatedBy int64  // 创建人
	PackageId int64  // 租户套餐编号
	LinkPhone string // 联系人手机
	TenantCreateUpdateBase
}

// TenantUpdateInput 租户编辑
type TenantUpdateInput struct {
	Id        int64
	UpdatedBy int64 // 更新人
	TenantCreateUpdateBase
}
