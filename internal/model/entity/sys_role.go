// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysRole is the golang structure for table sys_role.
type SysRole struct {
	Id         int64       `json:"id"         orm:"id"          ` // 角色id
	ParentId   int64       `json:"parentId"   orm:"parent_id"   ` // 上级角色ID
	Name       string      `json:"name"       orm:"name"        ` // 角色名称
	Code       string      `json:"code"       orm:"code"        ` // 角色标识
	Level      int         `json:"level"      orm:"level"       ` // 关系树等级
	Tree       string      `json:"tree"       orm:"tree"        ` // 关系树
	DataScope  int         `json:"dataScope"  orm:"data_scope"  ` // 数据范围 1全部/2当前部门/3当前及以下部门/4自定义部门
	CustomDept *gjson.Json `json:"customDept" orm:"custom_dept" ` // 自定义部门权限
	Status     int         `json:"status"     orm:"status"      ` // 角色状态 1正常/2停用
	Sort       int         `json:"sort"       orm:"sort"        ` // 排序
	Remark     string      `json:"remark"     orm:"remark"      ` // 备注
	CreatedBy  int64       `json:"createdBy"  orm:"created_by"  ` // 添加人
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  ` // 创建时间
	UpdatedBy  int64       `json:"updatedBy"  orm:"updated_by"  ` // 更新人
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  ` // 更新时间
	TenantId   int64       `json:"tenantId"   orm:"tenant_id"   ` // 租户ID
}
