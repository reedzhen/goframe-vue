package dto

import (
	"github.com/gogf/gf/v2/database/gdb"
	"goframe-vben/internal/library/query"
)

// DeptPageInput 部门管理分页
type DeptPageInput struct {
	query.PageParam
	DeptCode string // 部门编码
	DeptName string // 部门名称
}

func (q *DeptPageInput) Cond(m *gdb.Model) *gdb.Model {
	if q.DeptCode != "" {
		m = m.Where("dept_code", q.DeptCode)
	}
	if q.DeptName != "" {
		m = m.Where("dept_name", q.DeptName)
	}
	return m
}

// DeptGetListInput 部门管理列表
type DeptGetListInput struct {
	DeptCode string // 部门编码
	DeptName string // 部门名称
	IsTotal  bool   // 是否获取全部数据
}

// DeptCreateInput 部门管理新建
type DeptCreateInput struct {
	CreatedBy int64  // 创建人
	ParentId  int64  // 上级Id 0是顶级
	DeptCode  string // 部门编码
	DeptName  string // 部门名称
	Status    int    // 部门状态 1正常/2停用
	Level     int    // 关系树层级
	Tree      string // 关系树
	Sort      int    // 排序
	Remark    string // 备注
}

// DeptUpdateInput 部门管理修改
type DeptUpdateInput struct {
	Id        int64
	UpdatedBy int64  // 更新人
	ParentId  int64  // 上级Id 0是顶级
	DeptCode  string // 部门编码
	DeptName  string // 部门名称
	Status    int    // 部门状态 1正常/2停用
	//Level     int    // 关系树层级
	//Tree      string // 关系树
	Sort   int    // 排序
	Remark string // 备注
}
