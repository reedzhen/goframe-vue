// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDictData is the golang structure for table sys_dict_data.
type SysDictData struct {
	Id        int64       `json:"id"        orm:"id"         ` // 字典项id
	DictId    int64       `json:"dictId"    orm:"dict_id"    ` // 字典id
	Code      string      `json:"code"      orm:"code"       ` // 字典项标识
	CodeType  string      `json:"codeType"  orm:"code_type"  ` // 类型 string/int
	Name      string      `json:"name"      orm:"name"       ` // 字典项名称
	Sort      int         `json:"sort"      orm:"sort"       ` // 排序号
	Note      string      `json:"note"      orm:"note"       ` // 备注
	CreatedBy int64       `json:"createdBy" orm:"created_by" ` // 添加人
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by" ` // 更新人
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
}
