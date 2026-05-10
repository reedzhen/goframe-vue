package cron

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-vben/internal/library/query"
)

type RecordPageReq struct {
	g.Meta `path:"/cron/record/page" method:"get" tags:"定时任务" summary:"获取定时任务日志分页"`
	query.PageParam
	CronId int64 `json:"cron_id" in:"query" dc:"定时任务ID"`
}
type RecordPageRes struct {
	*query.Result
}
