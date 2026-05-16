// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysConfigItemDao is the data access object for the table sys_config_item.
type SysConfigItemDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  SysConfigItemColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// SysConfigItemColumns defines and stores column names for the table sys_config_item.
type SysConfigItemColumns struct {
	Id           string // 主键ID
	ModuleId     string // 所属配置模块ID
	Name         string // 配置项名称，后台展示用
	ConfigKey    string // 配置项键名，代码读取用唯一标识
	ConfigValue  string // 配置项当前值
	DefaultValue string // 默认值，读取失败或未配置时兜底
	ValueType    string // 值类型 string/int/uint/bool/json/datetime/date
	InputType    string // 输入类型 1输入框/2范围/3下拉/4单选/5开关/6多选
	InputParams  string // 控件参数，选项或范围配置，例如 A-1|B-2|C-3
	Description  string // 配置项说明
	Sort         string // 排序值，越小越靠前
	Status       string // 状态 1启用/0禁用
	IsSystem     string // 是否系统内置 1是/0否，内置配置通常由研发维护
	CreatedBy    string // 添加人
	CreatedAt    string // 创建时间
	UpdatedBy    string // 更新人
	UpdatedAt    string // 更新时间
	TenantId     string // 租户ID
}

// sysConfigItemColumns holds the columns for the table sys_config_item.
var sysConfigItemColumns = SysConfigItemColumns{
	Id:           "id",
	ModuleId:     "module_id",
	Name:         "name",
	ConfigKey:    "config_key",
	ConfigValue:  "config_value",
	DefaultValue: "default_value",
	ValueType:    "value_type",
	InputType:    "input_type",
	InputParams:  "input_params",
	Description:  "description",
	Sort:         "sort",
	Status:       "status",
	IsSystem:     "is_system",
	CreatedBy:    "created_by",
	CreatedAt:    "created_at",
	UpdatedBy:    "updated_by",
	UpdatedAt:    "updated_at",
	TenantId:     "tenant_id",
}

// NewSysConfigItemDao creates and returns a new DAO object for table data access.
func NewSysConfigItemDao(handlers ...gdb.ModelHandler) *SysConfigItemDao {
	return &SysConfigItemDao{
		group:    "default",
		table:    "sys_config_item",
		columns:  sysConfigItemColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysConfigItemDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysConfigItemDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysConfigItemDao) Columns() SysConfigItemColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysConfigItemDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysConfigItemDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysConfigItemDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
