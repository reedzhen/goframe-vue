// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysTenantPackage is the golang structure for table sys_tenant_package.
type SysTenantPackage struct {
	Id        int64       `json:"id"        orm:"id"         ` // 套餐编号
	Name      string      `json:"name"      orm:"name"       ` // 套餐名
	Status    int         `json:"status"    orm:"status"     ` // 套餐状态 1正常/2停用
	MenuIds   string      `json:"menuIds"   orm:"menu_ids"   ` // 菜单ID 逗号分隔
	Remark    string      `json:"remark"    orm:"remark"     ` // 备注
	CreatedBy int64       `json:"createdBy" orm:"created_by" ` // 添加人
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by" ` // 更新人
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
}
