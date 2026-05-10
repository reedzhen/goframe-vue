// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysMenuTempDao is the data access object for the table sys_menu_temp.
type SysMenuTempDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysMenuTempColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysMenuTempColumns defines and stores column names for the table sys_menu_temp.
type SysMenuTempColumns struct {
	Id        string // 菜单id
	ParentId  string // 上级id, 0是顶级
	Title     string // 菜单名称
	Path      string // 菜单路由地址
	Component string // 菜单组件地址, 目录可为空
	MenuType  string // 类型 menu菜单/button按钮
	Sort      string // 排序号
	Authority string // 按钮权限
	Icon      string // 菜单图标
	Hide      string // 是否隐藏 1是/2否(仅注册路由不显示在左侧菜单)
	MenuMeta  string // 路由元信息
	ApiPath   string // 接口权限 默认存数组
	DeletedAt string // 删除时间
	CreatedBy string // 添加人
	CreatedAt string // 创建时间
	UpdatedBy string // 更新人
	UpdatedAt string // 更新时间
}

// sysMenuTempColumns holds the columns for the table sys_menu_temp.
var sysMenuTempColumns = SysMenuTempColumns{
	Id:        "id",
	ParentId:  "parent_id",
	Title:     "title",
	Path:      "path",
	Component: "component",
	MenuType:  "menu_type",
	Sort:      "sort",
	Authority: "authority",
	Icon:      "icon",
	Hide:      "hide",
	MenuMeta:  "menu_meta",
	ApiPath:   "api_path",
	DeletedAt: "deleted_at",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
}

// NewSysMenuTempDao creates and returns a new DAO object for table data access.
func NewSysMenuTempDao(handlers ...gdb.ModelHandler) *SysMenuTempDao {
	return &SysMenuTempDao{
		group:    "default",
		table:    "sys_menu_temp",
		columns:  sysMenuTempColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysMenuTempDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysMenuTempDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysMenuTempDao) Columns() SysMenuTempColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysMenuTempDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysMenuTempDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysMenuTempDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
