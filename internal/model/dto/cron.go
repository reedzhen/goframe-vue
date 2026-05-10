package dto

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"goframe-vben/internal/consts"
	"goframe-vben/internal/library/query"
	"goframe-vben/utility/tools"
)

// CronPageInput 定时任务分页
type CronPageInput struct {
	query.PageParam
	Title string // 标题
	Tag   string // 标签 字典表获取
}

func (q *CronPageInput) Cond(m *gdb.Model) *gdb.Model {
	if q.Title != "" {
		m = m.WhereLike("title", tools.TrimLike(q.Title))
	}
	if q.Tag != "" {
		m = m.WhereLike("tag", tools.TrimLike(q.Tag))
	}
	return m
}

// CronCreateUpdateBase 定时任务新增或编辑
type CronCreateUpdateBase struct {
	Title     string // 标题
	Tag       string // 标签 字典表获取
	ApiUrl    string // 接口地址
	ApiParam  string // 接口请求参数
	ApiHeader string // 接口请求头
	Pattern   string // cron表达式
	Policy    int    // 策略 1并行/2单例/3单次/4多次
	Count     int    // 执行次数
	Sort      int    // 排序
	Remark    string // 备注
}

// CronCreateInput 定时任务新增
type CronCreateInput struct {
	CronCreateUpdateBase
	CreatedBy int64             // 创建人
	Status    consts.CronStatus // 任务状态 active运行中/inactive已结束
}

// CronUpdateInput 定时任务编辑
type CronUpdateInput struct {
	Id        int64
	UpdatedBy int64 // 更新人
	CronCreateUpdateBase
}

// CronGetListInput 定时任务列表
type CronGetListInput struct {
	Title string // 标题
	Tag   string // 标签 字典表获取
}

// CronChangeStatusInput 定时任务编辑状态
type CronChangeStatusInput struct {
	Id        int64
	Status    consts.CronStatus // 任务状态 active运行中/inactive已结束
	UpdatedBy int64             // 更新人
}

// CronUpdateRunAtInput 定时任务编辑
type CronUpdateRunAtInput struct {
	Id        int64
	Policy    int // 策略 1并行/2单例/3单次/4多次
	NextRunAt *gtime.Time
	LastRunAt *gtime.Time
}
