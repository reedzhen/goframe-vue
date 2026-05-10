// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. Created at 2026-03-29 15:49:32
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysDept is the golang structure for table sys_dept.
type SysDept struct {
	Id        int64       `json:"id"        orm:"id"         ` //
	ParentId  int64       `json:"parentId"  orm:"parent_id"  ` // 上级Id 0是顶级
	DeptCode  string      `json:"deptCode"  orm:"dept_code"  ` // 部门编码
	DeptName  string      `json:"deptName"  orm:"dept_name"  ` // 部门名称
	Status    int         `json:"status"    orm:"status"     ` // 部门状态 1正常/2停用
	Level     int         `json:"level"     orm:"level"      ` // 关系树层级
	Tree      string      `json:"tree"      orm:"tree"       ` // 关系树
	Sort      int         `json:"sort"      orm:"sort"       ` // 排序
	Remark    string      `json:"remark"    orm:"remark"     ` // 备注
	CreatedBy int64       `json:"createdBy" orm:"created_by" ` // 添加人
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" ` // 创建时间
	UpdatedBy int64       `json:"updatedBy" orm:"updated_by" ` // 更新人
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" ` // 更新时间
	TenantId  int64       `json:"tenantId"  orm:"tenant_id"  ` // 租户ID
}
