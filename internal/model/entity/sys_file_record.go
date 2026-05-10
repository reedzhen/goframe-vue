// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysFileRecord is the golang structure for table sys_file_record.
type SysFileRecord struct {
	Id        int64       `json:"id"        orm:"id"         ` // 主键id
	Name      string      `json:"name"      orm:"name"       ` // 文件名称
	Path      string      `json:"path"      orm:"path"       ` // 文件存储路径
	Length    int         `json:"length"    orm:"length"     ` // 文件大小
	Type      string      `json:"type"      orm:"type"       ` // 文件类型：img图片/file文件
	Note      string      `json:"note"      orm:"note"       ` // 备注
	CreatedBy int64       `json:"createdBy" orm:"created_by" ` // 创建人
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 修改时间
}
