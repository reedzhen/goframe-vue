package dto

import (
	"github.com/gogf/gf/v2/database/gdb"
	"goframe-vben/internal/library/query"
)

// CronRecordPageInput 定时任务日志分页
type CronRecordPageInput struct {
	query.PageParam
	CronId int64
}

func (q *CronRecordPageInput) Cond(m *gdb.Model) *gdb.Model {
	if q.CronId > 0 {
		m = m.Where("cron_id", q.CronId)
	}
	return m
}

// CronRecordCreateInput 定时任务日志新增
type CronRecordCreateInput struct {
	CronId    int64  // 任务Id sys_cron.id
	SpendTime int    // 消耗时间, 单位毫秒
	Output    string // 执行结果或输出，可能包括错误信息
	Status    string // 执行状态 success成功/failure失败
	Remark    string // 备注
}
