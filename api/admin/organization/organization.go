package organization

import (
	"github.com/gogf/gf/v2/frame/g"
	"goframe-vben/internal/model/entity"
)

type GetListReq struct {
	g.Meta `path:"/organization/list" method:"get" tags:"组织机构" summary:"获取组织机构列表" noAuth:"true"`
	Name   string `json:"name" in:"query" dc:"机构名称"`
}
type GetListRes []*entity.SysOrganization

type CreateReq struct {
	g.Meta `path:"/organization/create" method:"post" tags:"组织机构" summary:"新建组织机构"`
	CreateUpdateBase
}
type CreateRes struct {
	Id int64 `json:"id" dc:"组织机构Id"`
}

type UpdateReq struct {
	g.Meta `path:"/organization/update" method:"post" tags:"组织机构" summary:"修改组织机构"`
	Id     int64 `json:"id" v:"required|min:1#请选择要修改的组织机构|请选择要修改的组织机构" dc:"主键ID"`
	CreateUpdateBase
}
type UpdateRes struct{}

type CreateUpdateBase struct {
	ParentId int64  `json:"parentId" v:"required#请输入上级id," dc:"上级id, 0是顶级"`
	Code     string `json:"code"  dc:"机构代码"`
	Name     string `json:"name" v:"required#请输入机构名称" dc:"机构名称"`
	FullName string `json:"fullName" v:"required#请输入机构全称" dc:"机构全称"`
	Type     int    `json:"type" v:"required#请输入机构类型" dc:"机构类型"`
	Status   int    `json:"status" v:"required#请输入部门状态" d:"1" dc:"部门状态 1正常/2停用"`
	//LinkId    int64  `json:"linkId" v:"required#请输入负责人id" dc:"负责人id sys_user.id"`
	//LinkMan   string `json:"linkMan" v:"required#请输入联系人" dc:"联系人"`
	//LinkPhone string `json:"linkPhone" v:"required#请输入联系电话" dc:"联系电话"`
	Remark string `json:"remark" dc:"备注"`
}

type DeleteReq struct {
	g.Meta `path:"/organization/delete/{Id}" method:"post" tags:"组织机构" summary:"删除组织机构"`
	Id     int64 `in:"path" v:"required|min:1#请选择要删除的组织机构|请选择要删除的组织机构" dc:"主键ID"`
}
type DeleteRes struct{}

type GetInfoReq struct {
	g.Meta `path:"/organization/info/{Id}" method:"get" tags:"组织机构" summary:"获取组织机构详情"`
	Id     int64 `in:"path" v:"required|min:1#请选择要查看的组织机构|请选择要查看的组织机构" dc:"主键ID"`
}
type GetInfoRes struct {
	*entity.SysOrganization
}
