package dto

import (
	"goframe-vben/internal/model/entity"
)

// OrganizationPageOutput 组织分页返回
type OrganizationPageOutput struct {
	*entity.SysOrganization
	ParentOrg *entity.SysOrganization `json:"parentOrg" orm:"with:id=parent_id"`
}

// OrganizationGetListInput 组织机构列表
type OrganizationGetListInput struct {
	Name string // 机构名称
}

// OrganizationCreateUpdateBase 组织机构新建或修改
type OrganizationCreateUpdateBase struct {
	ParentId  int64  // 上级id, 0是顶级
	Code      string // 机构代码
	Name      string // 机构名称
	FullName  string // 机构全称
	Type      int    // 机构类型
	Status    int    // 部门状态 1正常/2停用
	LinkId    int64  // 负责人id sys_user.id
	LinkMan   string // 联系人
	LinkPhone string // 联系电话
	Level     int    // 关系树层级
	Tree      string // 关系树
	Remark    string // 备注
}

// OrganizationCreateInput 组织机构新建
type OrganizationCreateInput struct {
	CreatedBy int64 // 创建人
	OrganizationCreateUpdateBase
}

// OrganizationUpdateInput 组织机构修改
type OrganizationUpdateInput struct {
	Id        int64
	UpdatedBy int64 // 更新人
	OrganizationCreateUpdateBase
}
