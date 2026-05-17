// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"goframe-vben/api"
	"goframe-vben/internal/library/query"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/model/entity"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IConfig interface {
		// ModuleList 获取配置模块列表
		ModuleList(ctx context.Context, in dto.ConfigModuleListInput) (out []*entity.SysConfigModule, err error)
		// ModuleCreate 新增配置模块
		ModuleCreate(ctx context.Context, in dto.ConfigModuleCreateInput) error
		// ModuleUpdate 编辑配置模块
		ModuleUpdate(ctx context.Context, in dto.ConfigModuleUpdateInput) error
		// ModuleDelete 删除配置模块
		ModuleDelete(ctx context.Context, id int64) error
		// ItemPage 获取配置项分页
		ItemPage(ctx context.Context, in dto.ConfigItemPageInput) (out *query.Result, err error)
		// ItemCreate 新增配置项
		ItemCreate(ctx context.Context, in dto.ConfigItemCreateInput) error
		// ItemUpdate 编辑配置项
		ItemUpdate(ctx context.Context, in dto.ConfigItemUpdateInput) error
		// ItemDelete 删除配置项
		ItemDelete(ctx context.Context, id int64) error
		// GetValuesByModuleCode 根据模块编码获取配置值
		GetValuesByModuleCode(ctx context.Context, moduleCode string) (out *dto.ConfigGetValuesOutput, err error)
		// SaveValuesByModuleCode 保存模块配置值
		SaveValuesByModuleCode(ctx context.Context, in dto.ConfigSaveValuesInput) error
		// GetUpload 获取上传配置
		GetUpload(ctx context.Context) (out *dto.ConfigUploadOutput, err error)
	}
	ICron interface {
		// Page 分页
		Page(ctx context.Context, in dto.CronPageInput) (out *query.Result, err error)
		// Create 新增
		Create(ctx context.Context, in dto.CronCreateInput) (err error)
		// Update 编辑
		Update(ctx context.Context, in dto.CronUpdateInput) error
		// UpdateLastRunAt 编辑最新执行时间
		UpdateLastRunAt(ctx context.Context, in dto.CronUpdateRunAtInput) error
		// ExecuteOnce 在线执行一次
		ExecuteOnce(ctx context.Context, cronId int64) error
		// Delete 删除
		Delete(ctx context.Context, id int64) error
		// GetInfo 详情
		GetInfo(ctx context.Context, id int64) (out *entity.SysCron, err error)
		// GetList 列表
		GetList(ctx context.Context, in dto.CronGetListInput) (out []*entity.SysCron, err error)
		// Start 启动所有cron
		Start(ctx context.Context)
		// Remove 移除所有cron
		Remove()
		// ChangeStatus 修改状态
		ChangeStatus(ctx context.Context, in dto.CronChangeStatusInput) error
	}
	ICronRecord interface {
		// Create 新增
		Create(ctx context.Context, in dto.CronRecordCreateInput) (err error)
		// Page 分页
		Page(ctx context.Context, in dto.CronRecordPageInput) (out *query.Result, err error)
		// DeleteByCronId 删除
		DeleteByCronId(ctx context.Context, cronId int64) error
	}
	IDict interface {
		// Page 字典分页
		Page(ctx context.Context, in dto.DictPageInput) (*query.Result, error)
		// Create 字典新增
		Create(ctx context.Context, in dto.DictCreateInput) (err error)
		// Update 字典编辑
		Update(ctx context.Context, in dto.DictUpdateInput) (err error)
		// Delete 字典删除
		Delete(ctx context.Context, id int64) (err error)
		// GetList 字典列表
		GetList(ctx context.Context) (out []*entity.SysDict, err error)
	}
	IDictData interface {
		// Page 字典项分页
		Page(ctx context.Context, in dto.DictDataPageInput) (*query.Result, error)
		// Create 字典项新增
		Create(ctx context.Context, in dto.DictDataCreateInput) (err error)
		// Update 字典项编辑
		Update(ctx context.Context, in dto.DictDataUpdateInput) (err error)
		// Delete 字典项删除
		Delete(ctx context.Context, id int64) (err error)
		// GetList 获取字典项列表
		GetList(ctx context.Context, in dto.DictDataGetListInput) (out []*entity.SysDictData, err error)
	}
	ILoginRecord interface {
		// AsyncCreate 异步调用登录日志
		AsyncCreate(ctx context.Context, in dto.LoginRecordCreateInput)
		// Create 新建登录日志
		Create(ctx context.Context, in dto.LoginRecordCreateInput) (err error)
		// Page 登录日志分页
		Page(ctx context.Context, in dto.LoginRecordPageInput) (out *query.Result, err error)
	}
	IMenu interface {
		// GetList 获取菜单列表
		GetList(ctx context.Context, in dto.MenuGetListInput) (out []*entity.SysMenu, err error)
		// Create 新建菜单
		Create(ctx context.Context, in dto.MenuCreateInput) (out int64, err error)
		// ValidateExists 获取并校验详情
		ValidateExists(ctx context.Context, id int64) (out *entity.SysMenu, err error)
		// Update 编辑菜单
		Update(ctx context.Context, in dto.MenuUpdateInput) (err error)
		// Delete 删除菜单
		Delete(ctx context.Context, id int64) (err error)
		// GetAuthorizedList 获取当前用户拥有的菜单列表
		GetAuthorizedList(ctx context.Context) ([]*entity.SysMenu, error)
		// LoadRolePermissions 加载角色权限到缓存
		LoadRolePermissions(ctx context.Context, roleId int64) (err error)
		// ClearRolePermissionCache 清除角色权限缓存
		ClearRolePermissionCache(ctx context.Context, roleId int64) (err error)
		// CheckRoleApiPermission 检查角色是否拥有指定接口权限
		CheckRoleApiPermission(ctx context.Context, roleId int64, route string, method string) (bool, error)
		// CheckRolePermissionFast 使用缓存检查用户权限
		CheckRolePermissionFast(ctx context.Context, roleId int64, authKey string) (bool, error)
		// CheckRolePermission 检查角色是否拥有指定权限
		CheckRolePermission(ctx context.Context, roleId int64, authKey string) (bool, error)
	}
	IOperationRecord interface {
		// Page 操作日志分页
		Page(ctx context.Context, in dto.OperationRecordPageInput) (out *query.Result, err error)
		// AsyncCreate 异步调用操作日志
		AsyncCreate(ctx context.Context, in dto.OperationRecordCreateInput)
		// Create 新建操作日志
		Create(ctx context.Context, in dto.OperationRecordCreateInput) (err error)
		// Delete 删除超过n天的日志
		Delete(ctx context.Context, day int) (err error)
	}
	IOrganization interface {
		// Create 新建组织
		Create(ctx context.Context, in dto.OrganizationCreateInput) (out int64, err error)
		// Update 编辑组织
		Update(ctx context.Context, in dto.OrganizationUpdateInput) error
		// Delete 删除组织
		Delete(ctx context.Context, id int64) error
		// GetList 获取组织列表
		GetList(ctx context.Context, in dto.OrganizationGetListInput) ([]*entity.SysOrganization, error)
		// GetInfo 详情
		GetInfo(ctx context.Context, id int64) (out *entity.SysOrganization, err error)
		// ValidateExists 获取并校验详情
		ValidateExists(ctx context.Context, id int64) (out *entity.SysOrganization, err error)
		// GetCountByTenantId 获取数量
		GetCountByTenantId(ctx context.Context, tenantId int64) int
	}
	IRole interface {
		// Page 角色分页
		Page(ctx context.Context, in dto.RolePageInput) (out *query.Result, err error)
		// Create 新建角色
		Create(ctx context.Context, in dto.RoleCreateInput) (out int64, err error)
		// Update 编辑角色
		Update(ctx context.Context, in dto.RoleUpdateInput) (err error)
		// Delete 删除角色
		Delete(ctx context.Context, id int64) error
		// DeleteByTenantId 删除角色
		DeleteByTenantId(ctx context.Context, tenantId int64) error
		// GetChildCount 获取数量
		GetChildCount(ctx context.Context, parentId int64) int
		// GetList 获取角色列表(默认获取当前用户角色的下属角色列表)
		GetList(ctx context.Context, in dto.RoleGetListInput) ([]*entity.SysRole, error)
		// GetInfo 获取详情
		GetInfo(ctx context.Context, id int64) (out *entity.SysRole, err error)
		// ValidateExists 验证并获取详情
		ValidateExists(ctx context.Context, id int64) (out *entity.SysRole, err error)
		// GetRoleMenuList 获取角色下所有菜单
		GetRoleMenuList(ctx context.Context, roleId int64) ([]*api.MenuItem, error)
		// SaveRoleMenuList 保存角色下所有菜单
		SaveRoleMenuList(ctx context.Context, in dto.RoleSaveMenuInput) (err error)
		// VerifyRoleId 验证传入角色Id在不在当前登录用户的下级角色里
		VerifyRoleId(ctx context.Context, roleId int64) error
		// GetRoleMenuListByCond 获取角色下所有菜单
		GetRoleMenuListByCond(ctx context.Context, roleIds []int64) (out []*entity.SysRoleMenu, err error)
		// ChangeDataScope 修改角色数据权限
		ChangeDataScope(ctx context.Context, roleId int64, dataScope int) (err error)
		// GetSubIds 获取下级角色的Id
		GetSubIds(ctx context.Context, roleId int64) []int64
	}
	ITenant interface {
		// GetPage 获取租户分页
		GetPage(ctx context.Context, in dto.TenantPageInput) (res *query.Result, err error)
		// GetList 获取租户列表
		GetList(ctx context.Context, in dto.TenantGetListInput) ([]*entity.SysTenant, error)
		// Create 新增租户
		Create(ctx context.Context, in dto.TenantCreateInput) error
		// Update 编辑租户
		Update(ctx context.Context, in dto.TenantUpdateInput) error
		// Delete 删除租户
		Delete(ctx context.Context, id int64) error
		// DeleteTenantCache 删除租户redis缓存
		DeleteTenantCache(ctx context.Context, tenant *entity.SysTenant) error
		// GetInfo 获取租户详情
		GetInfo(ctx context.Context, id int64) (out *entity.SysTenant, err error)
		// GetInfoWithCache 获取租户详情
		GetInfoWithCache(r *ghttp.Request) (out *entity.SysTenant, err error)
	}
	ITenantPackage interface {
		// GetPage 获取租户套餐分页
		GetPage(ctx context.Context, in dto.TenantPackagePageInput) (res *query.Result, err error)
		// GetList 获取租户套餐列表
		GetList(ctx context.Context, in dto.TenantPackageGetListInput) ([]*entity.SysTenantPackage, error)
		// Create 新建租户套餐
		Create(ctx context.Context, in dto.TenantPackageCreateInput) (err error)
		// Update 修改租户套餐
		Update(ctx context.Context, in dto.TenantPackageUpdateInput) (err error)
		// Delete 删除租户套餐
		Delete(ctx context.Context, id int64) error
		// GetInfo 获取租户套餐详情
		GetInfo(ctx context.Context, id int64) (out *entity.SysTenantPackage, err error)
		// GetMenuList 获取所有菜单,套餐已包含则checked==true
		GetMenuList(ctx context.Context, packageId int64) ([]*api.MenuItem, error)
		// SyncTenantPermission 给使用此套餐的租户同步权限（套餐可能会加减权限）
		SyncTenantPermission(ctx context.Context, id int64) (err error)
	}
	IUser interface {
		// ValidateExists 获取个人信息
		ValidateExists(ctx context.Context, id int64) (out *entity.SysUser, err error)
		// GetInfoByUsername 获取个人信息
		GetInfoByUsername(ctx context.Context, username string) (out *entity.SysUser, err error)
		// Page 用户分页
		Page(ctx context.Context, in dto.UserPageInput) (out *query.Result, err error)
		// Create 新建用户
		Create(ctx context.Context, in dto.UserCreateInput) (out int64, err error)
		// Update 修改用户
		Update(ctx context.Context, in dto.UserUpdateInput) (err error)
		// Delete 删除用户
		Delete(ctx context.Context, id int64) (err error)
		// DeleteByTenantId 通过贴牌Id删除用户
		DeleteByTenantId(ctx context.Context, tenantId int64) (err error)
		// UpdatePwd 修改密码
		UpdatePwd(ctx context.Context, oldPwd string, newPwd string) (err error)
		// ResetPwd 重置密码 123456
		ResetPwd(ctx context.Context, userId int64, newPwd string) (err error)
		// ChangeStatus 修改用户状态
		ChangeStatus(ctx context.Context, userId int64, status int) (err error)
		// ChangeLastLoginAt 登录成功记录登录时间
		ChangeLastLoginAt(ctx context.Context, userId int64) (err error)
		// GetCountByRoleId 获取数量
		GetCountByRoleId(ctx context.Context, roleId int64) int
		// GetUserIdsByDeptIds 获取用户Ids
		GetUserIdsByDeptIds(ctx context.Context, deptIds []int64) ([]int64, error)
		// GetList 获取用户列表
		GetList(ctx context.Context, in dto.UserGetListInput) ([]*entity.SysUser, error)
		// GetRoleIdByUserId 获取角色Id
		GetRoleIdByUserId(ctx context.Context, userId int64) (int64, error)
		// CheckFieldExist 检测给定的字段是否唯一
		CheckFieldExist(ctx context.Context, field string, value string, id ...int64) (err error)
	}
)

var (
	localConfig          IConfig
	localCron            ICron
	localCronRecord      ICronRecord
	localDict            IDict
	localDictData        IDictData
	localLoginRecord     ILoginRecord
	localMenu            IMenu
	localOperationRecord IOperationRecord
	localOrganization    IOrganization
	localRole            IRole
	localTenant          ITenant
	localTenantPackage   ITenantPackage
	localUser            IUser
)

func Config() IConfig {
	if localConfig == nil {
		panic("implement not found for interface IConfig, forgot register?")
	}
	return localConfig
}

func RegisterConfig(i IConfig) {
	localConfig = i
}

func Cron() ICron {
	if localCron == nil {
		panic("implement not found for interface ICron, forgot register?")
	}
	return localCron
}

func RegisterCron(i ICron) {
	localCron = i
}

func CronRecord() ICronRecord {
	if localCronRecord == nil {
		panic("implement not found for interface ICronRecord, forgot register?")
	}
	return localCronRecord
}

func RegisterCronRecord(i ICronRecord) {
	localCronRecord = i
}

func Dict() IDict {
	if localDict == nil {
		panic("implement not found for interface IDict, forgot register?")
	}
	return localDict
}

func RegisterDict(i IDict) {
	localDict = i
}

func DictData() IDictData {
	if localDictData == nil {
		panic("implement not found for interface IDictData, forgot register?")
	}
	return localDictData
}

func RegisterDictData(i IDictData) {
	localDictData = i
}

func LoginRecord() ILoginRecord {
	if localLoginRecord == nil {
		panic("implement not found for interface ILoginRecord, forgot register?")
	}
	return localLoginRecord
}

func RegisterLoginRecord(i ILoginRecord) {
	localLoginRecord = i
}

func Menu() IMenu {
	if localMenu == nil {
		panic("implement not found for interface IMenu, forgot register?")
	}
	return localMenu
}

func RegisterMenu(i IMenu) {
	localMenu = i
}

func OperationRecord() IOperationRecord {
	if localOperationRecord == nil {
		panic("implement not found for interface IOperationRecord, forgot register?")
	}
	return localOperationRecord
}

func RegisterOperationRecord(i IOperationRecord) {
	localOperationRecord = i
}

func Organization() IOrganization {
	if localOrganization == nil {
		panic("implement not found for interface IOrganization, forgot register?")
	}
	return localOrganization
}

func RegisterOrganization(i IOrganization) {
	localOrganization = i
}

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}

func Tenant() ITenant {
	if localTenant == nil {
		panic("implement not found for interface ITenant, forgot register?")
	}
	return localTenant
}

func RegisterTenant(i ITenant) {
	localTenant = i
}

func TenantPackage() ITenantPackage {
	if localTenantPackage == nil {
		panic("implement not found for interface ITenantPackage, forgot register?")
	}
	return localTenantPackage
}

func RegisterTenantPackage(i ITenantPackage) {
	localTenantPackage = i
}

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
