// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOperationRecord is the golang structure for table sys_operation_record.
type SysOperationRecord struct {
	Id         int64       `json:"id"         orm:"id"          ` // 主键
	UserId     int64       `json:"userId"     orm:"user_id"     ` // 用户id
	Username   string      `json:"username"   orm:"username"    ` // 账号
	Nickname   string      `json:"nickname"   orm:"nickname"    ` // 昵称
	Url        string      `json:"url"        orm:"url"         ` // 请求地址
	Method     string      `json:"method"     orm:"method"      ` // 请求方式
	Module     string      `json:"module"     orm:"module"      ` // 接口模块 例如：用户管理
	Summary    string      `json:"summary"    orm:"summary"     ` // 接口描述 例如：添加用户
	Param      string      `json:"param"      orm:"param"       ` // 请求参数
	JsonResult string      `json:"jsonResult" orm:"json_result" ` // 返回结果
	ErrorMsg   string      `json:"errorMsg"   orm:"error_msg"   ` // 异常信息
	SpendTime  int64       `json:"spendTime"  orm:"spend_time"  ` // 消耗时间, 单位毫秒
	TraceId    string      `json:"traceId"    orm:"trace_id"    ` // trace_id
	Status     int         `json:"status"     orm:"status"      ` // 状态 1成功/2异常
	Platform   string      `json:"platform"   orm:"platform"    ` // 平台 admin/api/open
	UserAgent  string      `json:"userAgent"  orm:"user_agent"  ` // 请求头User-Agent
	Ip         string      `json:"ip"         orm:"ip"          ` // 主机地址
	Remark     string      `json:"remark"     orm:"remark"      ` // 备注
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  ` // 创建时间
	TenantId   int64       `json:"tenantId"   orm:"tenant_id"   ` // 租户ID
}
