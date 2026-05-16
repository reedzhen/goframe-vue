// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysCronRecord is the golang structure for table sys_cron_record.
type SysCronRecord struct {
	Id        int64       `json:"id"        orm:"id"         ` //
	CronId    int64       `json:"cronId"    orm:"cron_id"    ` // 任务Id sys_cron.id
	SpendTime int         `json:"spendTime" orm:"spend_time" ` // 消耗时间, 单位毫秒
	Output    string      `json:"output"    orm:"output"     ` // 执行结果或输出，可能包括错误信息
	Status    string      `json:"status"    orm:"status"     ` // 执行状态 success成功/failure失败
	Remark    string      `json:"remark"    orm:"remark"     ` // 备注
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
}
