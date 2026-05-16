// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysConfigModule is the golang structure for table sys_config_module.
type SysConfigModule struct {
	Id          int64       `json:"id"          orm:"id"          ` // 主键ID
	Code        string      `json:"code"        orm:"code"        ` // 模块编码，代码中使用的唯一标识
	Name        string      `json:"name"        orm:"name"        ` // 模块名称
	Description string      `json:"description" orm:"description" ` // 模块说明
	Sort        int         `json:"sort"        orm:"sort"        ` // 排序值，越小越靠前
	Status      int         `json:"status"      orm:"status"      ` // 状态 1启用/0禁用
	CreatedBy   int64       `json:"createdBy"   orm:"created_by"  ` // 添加人
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  ` // 创建时间
	UpdatedBy   int64       `json:"updatedBy"   orm:"updated_by"  ` // 更新人
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  ` // 更新时间
	TenantId    int64       `json:"tenantId"    orm:"tenant_id"   ` // 租户ID
}
