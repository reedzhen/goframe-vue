// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysLoginRecordDao is the data access object for the table sys_login_record.
type SysLoginRecordDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  SysLoginRecordColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// SysLoginRecordColumns defines and stores column names for the table sys_login_record.
type SysLoginRecordColumns struct {
	Id        string // 主键
	Username  string // 用户账号
	Os        string // 操作系统
	Device    string // 设备名
	Browser   string // 浏览器类型
	Ip        string // ip地址
	IpCity    string // ip归属地
	LoginType string // 操作类型 1登录成功/2登录失败/3退出登录
	Remark    string // 备注
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	TenantId  string // 租户ID
}

// sysLoginRecordColumns holds the columns for the table sys_login_record.
var sysLoginRecordColumns = SysLoginRecordColumns{
	Id:        "id",
	Username:  "username",
	Os:        "os",
	Device:    "device",
	Browser:   "browser",
	Ip:        "ip",
	IpCity:    "ip_city",
	LoginType: "login_type",
	Remark:    "remark",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	TenantId:  "tenant_id",
}

// NewSysLoginRecordDao creates and returns a new DAO object for table data access.
func NewSysLoginRecordDao(handlers ...gdb.ModelHandler) *SysLoginRecordDao {
	return &SysLoginRecordDao{
		group:    "default",
		table:    "sys_login_record",
		columns:  sysLoginRecordColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysLoginRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysLoginRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysLoginRecordDao) Columns() SysLoginRecordColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysLoginRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysLoginRecordDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysLoginRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
