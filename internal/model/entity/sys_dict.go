// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDict is the golang structure for table sys_dict.
type SysDict struct {
	Id        int64       `json:"id"        orm:"id"         ` // 字典id
	Code      string      `json:"code"      orm:"code"       ` // 字典标识
	Name      string      `json:"name"      orm:"name"       ` // 字典名称
	Sort      uint        `json:"sort"      orm:"sort"       ` // 排序号
	Note      string      `json:"note"      orm:"note"       ` // 备注
	CreatedBy int64       `json:"createdBy" orm:"created_by" ` // 添加人
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by" ` // 更新人
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
}
