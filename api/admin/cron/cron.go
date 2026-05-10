package cron

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-vben/internal/consts"
	"goframe-vben/internal/library/query"
)

type PageReq struct {
	g.Meta `path:"/cron/page" method:"get" tags:"定时任务" summary:"定时任务分页"`
	query.PageParam
	Title string `json:"title" in:"query" dc:"标题"`
	Tag   string `json:"tag" in:"query" dc:"标签 字典表获取"`
}
type PageRes struct {
	*query.Result
}

type CreateReq struct {
	g.Meta `path:"/cron/create" method:"post" tags:"定时任务" summary:"创建定时任务"`
	CreateUpdateBase
}
type CreateRes struct{}

type UpdateReq struct {
	g.Meta `path:"/cron/update" method:"post" tags:"定时任务" summary:"编辑定时任务"`
	Id     int64 `json:"id" v:"required" dc:"主键ID"`
	CreateUpdateBase
}
type UpdateRes struct{}

type CreateUpdateBase struct {
	Title     string `json:"title" v:"required" dc:"标题"`
	Tag       string `json:"tag" v:"required" dc:"标签 字典表获取"`
	ApiUrl    string `json:"api_url" v:"required" dc:"接口地址"`
	ApiParam  string `json:"api_param"  dc:"接口请求参数"`
	ApiHeader string `json:"api_header"  dc:"接口请求头"`
	Pattern   string `json:"pattern" v:"required" dc:"cron表达式"`
	Policy    int    `json:"policy" v:"required" dc:"策略 1并行/2单例/3单次/4多次"`
	Count     int    `json:"count" dc:"执行次数"`
	Sort      int    `json:"sort" v:"required" dc:"排序"`
	Remark    string `json:"remark"  dc:"备注"`
}

type DeleteReq struct {
	g.Meta `path:"/cron/delete/{Id}" method:"post" tags:"定时任务" summary:"删除定时任务"`
	Id     int64 `in:"path"`
}
type DeleteRes struct{}

type ChangeStatusReq struct {
	g.Meta `path:"/cron/change-status" method:"post" tags:"定时任务" summary:"编辑定时任务状态"`
	Id     int64             `json:"id"  v:"required" dc:"主键ID"`
	Status consts.CronStatus `json:"status" v:"enums" dc:"状态"`
}
type ChangeStatusRes struct{}

type ExecuteOnceReq struct {
	g.Meta `path:"/cron/execute-once" method:"post" tags:"定时任务" summary:"直接执行一次定时任务"`
	Id     int64 `json:"id"  v:"required" dc:"主键ID"`
}
type ExecuteOnceRes struct{}
