// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysConfigDao is the data access object for the table sys_config.
type SysConfigDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysConfigColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysConfigColumns defines and stores column names for the table sys_config.
type SysConfigColumns struct {
	Id           string //
	Group        string // 配置分组
	Label        string // 中文名称 例如:姓名
	Key          string // 配置键名 例如:age
	Value        string // 配置值 例如:18
	DefaultValue string // 默认值
	Type         string // 类型 string,int,uint,bool,datetime,date
	Desc         string // 描述
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
	TenantId     string // 贴牌Id
}

// sysConfigColumns holds the columns for the table sys_config.
var sysConfigColumns = SysConfigColumns{
	Id:           "id",
	Group:        "group",
	Label:        "label",
	Key:          "key",
	Value:        "value",
	DefaultValue: "default_value",
	Type:         "type",
	Desc:         "desc",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	TenantId:     "tenant_id",
}

// NewSysConfigDao creates and returns a new DAO object for table data access.
func NewSysConfigDao(handlers ...gdb.ModelHandler) *SysConfigDao {
	return &SysConfigDao{
		group:    "default",
		table:    "sys_config",
		columns:  sysConfigColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysConfigDao) Columns() SysConfigColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysConfigDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
