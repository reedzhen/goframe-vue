package dto

import (
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gtime"
	"goframe-vben/internal/library/query"
	"goframe-vben/utility/tools"
)

// OperationRecordPageInput 操作日志分页
type OperationRecordPageInput struct {
	query.PageParam
	Username  string
	Module    string
	CreatedAt []string
}

// OperationRecordCreateInput 操作日志创建参数
type OperationRecordCreateInput struct {
	UserId     int64  // 用户ID
	Username   string // 用户账号
	Nickname   string // 用户昵称
	Url        string // 请求地址
	Method     string // 请求方式
	Module     string // 接口模块
	Summary    string // 接口描述
	Param      string // 请求参数
	JsonResult string // 返回结果
	ErrorMsg   string // 异常信息
	SpendTime  int64  // 消耗时间，单位毫秒
	TraceId    string // 链路ID
	Status     int    // 状态 1成功/2异常
	Platform   string // 平台 admin/api/open
	UserAgent  string // 请求头 User-Agent
	Ip         string // 客户端IP
	Remark     string // 备注
}

func (q *OperationRecordPageInput) Cond(m *gdb.Model) *gdb.Model {
	if q.Username != "" {
		m = m.WhereLike("username", tools.TrimLike(q.Username))
	}
	if q.Module != "" {
		m = m.WhereLike("module", tools.TrimLike(q.Module))
	}
	if len(q.CreatedAt) == 2 && q.CreatedAt[0] != "" && q.CreatedAt[1] != "" {
		m = m.WhereBetween("created_at", q.CreatedAt[0], gtime.NewFromStr(q.CreatedAt[1]).EndOfDay().String())
	}
	return m
}
