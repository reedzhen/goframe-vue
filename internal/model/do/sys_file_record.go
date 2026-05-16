// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysFileRecord is the golang structure of table sys_file_record for DAO operations like Where/Data.
type SysFileRecord struct {
	g.Meta    `orm:"table:sys_file_record, do:true"`
	Id        any         // 主键id
	Name      any         // 文件名称
	Path      any         // 文件存储路径
	Length    any         // 文件大小
	Type      any         // 文件类型：img图片/file文件
	Note      any         // 备注
	CreatedBy any         // 创建人
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
}
