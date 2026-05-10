// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserDao is the data access object for the table sys_user.
type SysUserDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysUserColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysUserColumns defines and stores column names for the table sys_user.
type SysUserColumns struct {
	Id             string //
	OrganizationId string // 机构Id
	Nickname       string // 昵称
	Username       string // 账号
	Password       string // 密码
	Salt           string // 盐加密
	RoleId         string // 角色Id
	Phone          string // 手机号
	Avatar         string // 头像
	Email          string // 邮箱
	EmailVerified  string // 邮箱是否验证:  1是 2否
	RealName       string // 真实姓名
	IdCard         string // 身份证号
	Birthday       string // 出生日期
	Introduction   string // 个人简介
	Status         string // 状态 1正常/2冻结
	LastLoginAt    string // 最新一次登录时间
	ParentId       string // 上级用户Id
	Level          string // 关系树等级
	Tree           string // 关系树
	IsAdmin        string // 是否超级管理员 1是/2否 默认：2
	CreatedBy      string // 添加人
	CreatedAt      string // 创建时间
	UpdatedBy      string // 更新人
	UpdatedAt      string // 更新时间
	TenantId       string // 租户ID
}

// sysUserColumns holds the columns for the table sys_user.
var sysUserColumns = SysUserColumns{
	Id:             "id",
	OrganizationId: "organization_id",
	Nickname:       "nickname",
	Username:       "username",
	Password:       "password",
	Salt:           "salt",
	RoleId:         "role_id",
	Phone:          "phone",
	Avatar:         "avatar",
	Email:          "email",
	EmailVerified:  "email_verified",
	RealName:       "real_name",
	IdCard:         "id_card",
	Birthday:       "birthday",
	Introduction:   "introduction",
	Status:         "status",
	LastLoginAt:    "last_login_at",
	ParentId:       "parent_id",
	Level:          "level",
	Tree:           "tree",
	IsAdmin:        "is_admin",
	CreatedBy:      "created_by",
	CreatedAt:      "created_at",
	UpdatedBy:      "updated_by",
	UpdatedAt:      "updated_at",
	TenantId:       "tenant_id",
}

// NewSysUserDao creates and returns a new DAO object for table data access.
func NewSysUserDao(handlers ...gdb.ModelHandler) *SysUserDao {
	return &SysUserDao{
		group:    "default",
		table:    "sys_user",
		columns:  sysUserColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysUserDao) Columns() SysUserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysUserDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
