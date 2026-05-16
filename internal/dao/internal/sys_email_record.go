// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysEmailRecordDao is the data access object for the table sys_email_record.
type SysEmailRecordDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  SysEmailRecordColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// SysEmailRecordColumns defines and stores column names for the table sys_email_record.
type SysEmailRecordColumns struct {
	Id           string // 主键id
	Title        string // 邮件标题
	Content      string // 邮件内容
	Receiver     string // 收件邮箱
	Sender       string // 发件邮箱
	CreateUserId string // 创建人
	Note         string // 备注
	TenantId     string // 租户id
	CreateTime   string // 创建时间
	UpdateTime   string // 修改时间
}

// sysEmailRecordColumns holds the columns for the table sys_email_record.
var sysEmailRecordColumns = SysEmailRecordColumns{
	Id:           "id",
	Title:        "title",
	Content:      "content",
	Receiver:     "receiver",
	Sender:       "sender",
	CreateUserId: "create_user_id",
	Note:         "note",
	TenantId:     "tenant_id",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
}

// NewSysEmailRecordDao creates and returns a new DAO object for table data access.
func NewSysEmailRecordDao(handlers ...gdb.ModelHandler) *SysEmailRecordDao {
	return &SysEmailRecordDao{
		group:    "default",
		table:    "sys_email_record",
		columns:  sysEmailRecordColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysEmailRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysEmailRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysEmailRecordDao) Columns() SysEmailRecordColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysEmailRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysEmailRecordDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysEmailRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
