// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysCron is the golang structure of table sys_cron for DAO operations like Where/Data.
type SysCron struct {
	g.Meta    `orm:"table:sys_cron, do:true"`
	Id        any         // 任务ID
	Title     any         // 标题
	Tag       any         // 标签 字典表获取
	ApiUrl    any         // 接口地址
	ApiParam  *gjson.Json // 接口请求参数
	ApiHeader *gjson.Json // 接口请求头
	Pattern   any         // cron表达式
	Status    any         // 任务状态 active运行中/inactive已结束
	Policy    any         // 策略 1并行/2单例/3单次/4多次
	Count     any         // 执行次数 policy=4时有效
	Sort      any         // 排序
	NextRunAt *gtime.Time // 下次预计运行时间
	LastRunAt *gtime.Time // 最后一次运行时间
	Remark    any         // 备注
	CreatedBy any         // 添加人
	CreatedAt *gtime.Time // 创建时间
	UpdatedBy any         // 更新人
	UpdatedAt *gtime.Time // 更新时间
}
