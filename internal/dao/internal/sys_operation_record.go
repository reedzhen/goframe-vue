// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysOperationRecordDao is the data access object for the table sys_operation_record.
type SysOperationRecordDao struct {
	table    string                    // table is the underlying table name of the DAO.
	group    string                    // group is the database configuration group name of the current DAO.
	columns  SysOperationRecordColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler        // handlers for customized model modification.
}

// SysOperationRecordColumns defines and stores column names for the table sys_operation_record.
type SysOperationRecordColumns struct {
	Id         string // 主键
	UserId     string // 用户id
	Username   string // 账号
	Nickname   string // 昵称
	Url        string // 请求地址
	Method     string // 请求方式
	Module     string // 接口模块 例如：用户管理
	Summary    string // 接口描述 例如：添加用户
	Param      string // 请求参数
	JsonResult string // 返回结果
	ErrorMsg   string // 异常信息
	SpendTime  string // 消耗时间, 单位毫秒
	TraceId    string // trace_id
	Status     string // 状态 1成功/2异常
	Platform   string // 平台 admin/api/open
	UserAgent  string // 请求头User-Agent
	Ip         string // 主机地址
	Remark     string // 备注
	CreatedAt  string // 创建时间
	TenantId   string // 租户ID
}

// sysOperationRecordColumns holds the columns for the table sys_operation_record.
var sysOperationRecordColumns = SysOperationRecordColumns{
	Id:         "id",
	UserId:     "user_id",
	Username:   "username",
	Nickname:   "nickname",
	Url:        "url",
	Method:     "method",
	Module:     "module",
	Summary:    "summary",
	Param:      "param",
	JsonResult: "json_result",
	ErrorMsg:   "error_msg",
	SpendTime:  "spend_time",
	TraceId:    "trace_id",
	Status:     "status",
	Platform:   "platform",
	UserAgent:  "user_agent",
	Ip:         "ip",
	Remark:     "remark",
	CreatedAt:  "created_at",
	TenantId:   "tenant_id",
}

// NewSysOperationRecordDao creates and returns a new DAO object for table data access.
func NewSysOperationRecordDao(handlers ...gdb.ModelHandler) *SysOperationRecordDao {
	return &SysOperationRecordDao{
		group:    "default",
		table:    "sys_operation_record",
		columns:  sysOperationRecordColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysOperationRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysOperationRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysOperationRecordDao) Columns() SysOperationRecordColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysOperationRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysOperationRecordDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *SysOperationRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
