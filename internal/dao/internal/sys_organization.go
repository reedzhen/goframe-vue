// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-05-16 22:56:55
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysOrganizationDao is the data access object for the table sys_organization.
type SysOrganizationDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  SysOrganizationColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// SysOrganizationColumns defines and stores column names for the table sys_organization.
type SysOrganizationColumns struct {
	Id        string //
	ParentId  string // 上级id, 0是顶级
	Code      string // 机构代码
	Name      string // 机构名称
	FullName  string // 机构全称
	Type      string // 机构类型
	Status    string // 部门状态 1正常/2停用
	LinkId    string // 负责人id sys_user.id
	LinkMan   string // 联系人
	LinkPhone string // 联系电话
	Level     string // 关系树层级
	Tree      string // 关系树
	Remark    string // 备注
	CreatedBy string // 添加人
	CreatedAt string // 创建时间
	UpdatedBy string // 更新人
	UpdatedAt string // 更新时间
	TenantId  string // 租户ID
}

// sysOrganizationColumns holds the columns for the table sys_organization.
var sysOrganizationColumns = SysOrganizationColumns{
	Id:        "id",
	ParentId:  "parent_id",
	Code:      "code",
	Name:      "name",
	FullName:  "full_name",
	Type:      "type",
	Status:    "status",
	LinkId:    "link_id",
	LinkMan:   "link_man",
	LinkPhone: "link_phone",
	Level:     "level",
	Tree:      "tree",
	Remark:    "remark",
	CreatedBy: "created_by",
	CreatedAt: "created_at",
	UpdatedBy: "updated_by",
	UpdatedAt: "updated_at",
	TenantId:  "tenant_id",
}

// NewSysOrganizationDao creates and returns a new DAO object for table data access.
func NewSysOrganizationDao(handlers ...gdb.ModelHandler) *SysOrganizationDao {
	return &SysOrganizationDao{
		group:    "default",
		table:    "sys_organization",
		columns:  sysOrganizationColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysOrganizationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysOrganizationDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysOrganizationDao) Columns() SysOrganizationColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysOrganizationDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysOrganizationDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysOrganizationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
