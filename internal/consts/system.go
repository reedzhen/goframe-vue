package consts

// CronStatus 定时任务状态
type CronStatus string
type TenantStatus int

const (
	UserDataScopeSuper  = "super"  // 超级后台用户
	UserDataScopeAdmin  = "admin"  // 公司管理员
	UserDataScopeCustom = "custom" // 店面用户 配合 user_shop 使用

	OperationRecordStatusSuc = 1 // 操作日志-成功
	OperationRecordStatusErr = 2 // 操作日志-失败

	MenuTypeMenu   = 1 // 菜单类型-菜单
	MenuTypeButton = 2 // 菜单类型-按钮

	MenuIsHideYes = 1 // 隐藏
	MenuIsHideNo  = 2 // 显示

	MenuShowTypeSuper   = "super"   // 超级管理员专属菜单
	MenuShowTypeCompany = "company" // 公司才能看到的菜单
	//MenuShowTypeAgent   = "agent"   // 待发货公司才能看到的菜单

	LoginRecordTypeSuc = 1 // 登录日志-登录成功
	LoginRecordTypeErr = 2 // 登录日志-登录失败

	CronStatusActive   CronStatus = "active"   // 运行中
	CronStatusInactive CronStatus = "inactive" // 已结束

	CronPolicyParallel = 1 // 并行策略
	CronPolicySingle   = 2 // 单例策略
	CronPolicyOnce     = 3 // 单次策略
	CronPolicyTimes    = 4 // 多次策略

	CronRecordStatusSuccess = "success" // cron日志成功
	CronRecordStatusFailure = "failure" // cron日志失败

	TenantStatusOk      TenantStatus = 1 // 租户状态可用
	TenantStatusDisable TenantStatus = 2 // 租户套餐不可用

	TenantPackageStatusOk      = 1 // 租户套餐状态可用
	TenantPackageStatusDisable = 2 // 租户套餐不可用

)

type OrgStatus int

const (
	OrgStatusOk      OrgStatus = 1 // 组织状态可用
	OrgStatusDisable OrgStatus = 2 // 组织停用
)

// UserStatus 用户状态
type UserStatus int

const (
	UserStatusOk     = 1 // 用户状态可用
	UserStatusFrozen = 2 // 冻结不可登录
)

// UserIsAdmin 是否是超管
type UserIsAdmin int

const (
	UserIsAdminYes = 1 // 超管
	UserIsAdminNo  = 2 // 非超管
)

// RoleStatus 角色状态
type RoleStatus int

const (
	RoleStatusOk      RoleStatus = 1 // 角色状态可用
	RoleStatusDisable RoleStatus = 2 // 角色状态停用
)

// 上传存储驱动
const (
	UploadDriveLocal = "local" // 本地驱动
	UploadDriveOss   = "oss"   // 阿里云oss
	UploadDriveQiNiu = "qiniu" // 七牛云
)

// RoleDataScope 角色数据权限
type RoleDataScope int

// 数据范围常量定义（值越大，范围越宽）
const (
	DataScopeSelf         RoleDataScope = 1 // 仅本人数据
	DataScopeDept         RoleDataScope = 2 // 本部门数据
	DataScopeDeptAndChild RoleDataScope = 3 // 本部门及下级部门
	DataScopeCompany      RoleDataScope = 4 // 全公司数据
	DataScopeCustom       RoleDataScope = 5 // 自定义数据范围
	DataScopeAll          RoleDataScope = 6 // 所有数据（超级管理员）
)
