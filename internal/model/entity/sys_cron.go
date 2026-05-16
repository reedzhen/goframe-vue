// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysCron is the golang structure for table sys_cron.
type SysCron struct {
	Id        int64       `json:"id"        orm:"id"          ` // 任务ID
	Title     string      `json:"title"     orm:"title"       ` // 标题
	Tag       string      `json:"tag"       orm:"tag"         ` // 标签 字典表获取
	ApiUrl    string      `json:"apiUrl"    orm:"api_url"     ` // 接口地址
	ApiParam  *gjson.Json `json:"apiParam"  orm:"api_param"   ` // 接口请求参数
	ApiHeader *gjson.Json `json:"apiHeader" orm:"api_header"  ` // 接口请求头
	Pattern   string      `json:"pattern"   orm:"pattern"     ` // cron表达式
	Status    string      `json:"status"    orm:"status"      ` // 任务状态 active运行中/inactive已结束
	Policy    int         `json:"policy"    orm:"policy"      ` // 策略 1并行/2单例/3单次/4多次
	Count     int         `json:"count"     orm:"count"       ` // 执行次数 policy=4时有效
	Sort      int         `json:"sort"      orm:"sort"        ` // 排序
	NextRunAt *gtime.Time `json:"nextRunAt" orm:"next_run_at" ` // 下次预计运行时间
	LastRunAt *gtime.Time `json:"lastRunAt" orm:"last_run_at" ` // 最后一次运行时间
	Remark    string      `json:"remark"    orm:"remark"      ` // 备注
	CreatedBy int64       `json:"createdBy" orm:"created_by"  ` // 添加人
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at"  ` // 创建时间
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by"  ` // 更新人
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at"  ` // 更新时间
}
