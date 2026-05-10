// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysCronRecord is the golang structure of table sys_cron_record for DAO operations like Where/Data.
type SysCronRecord struct {
	g.Meta    `orm:"table:sys_cron_record, do:true"`
	Id        any         //
	CronId    any         // 任务Id sys_cron.id
	SpendTime any         // 消耗时间, 单位毫秒
	Output    any         // 执行结果或输出，可能包括错误信息
	Status    any         // 执行状态 success成功/failure失败
	Remark    any         // 备注
	CreatedAt *gtime.Time // 创建时间
}
