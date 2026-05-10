// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysOperationRecord is the golang structure of table sys_operation_record for DAO operations like Where/Data.
type SysOperationRecord struct {
	g.Meta     `orm:"table:sys_operation_record, do:true"`
	Id         any         // 主键
	UserId     any         // 用户id
	Username   any         // 账号
	Nickname   any         // 昵称
	Url        any         // 请求地址
	Method     any         // 请求方式
	Module     any         // 接口模块 例如：用户管理
	Summary    any         // 接口描述 例如：添加用户
	Param      any         // 请求参数
	JsonResult any         // 返回结果
	ErrorMsg   any         // 异常信息
	SpendTime  any         // 消耗时间, 单位毫秒
	TraceId    any         // trace_id
	Status     any         // 状态 1成功/2异常
	Platform   any         // 平台 admin/api/open
	UserAgent  any         // 请求头User-Agent
	Ip         any         // 主机地址
	Remark     any         // 备注
	CreatedAt  *gtime.Time // 创建时间
	TenantId   any         // 租户ID
}
