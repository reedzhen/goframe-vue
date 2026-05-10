package ${.table.short_name|CaseLower}

import (
    "github.com/gogf/gf/v2/frame/g"
    "${.table.module_name}/internal/library/query"
    "${.table.module_name}/internal/model/entity"
    ${- if eq .table.has_time "true"}
        "github.com/gogf/gf/v2/os/gtime"
    ${- end}
)

type GetPageReq struct {
	g.Meta `path:"/${.table.short_name|CaseKebab}/page" method:"get" tags:"[后台]${.table.comment}" summary:"获取${.table.comment}分页"`
	query.PageParam
${- range $index, $elem := .table.fields}
${- if eq $elem.need_search "true"}
    ${$elem.name|CaseCamel} ${$elem.type} `json:"${$elem.json_name}" in:"query" dc:"${$elem.comment}"`
${- end}
${- end}
}
type GetPageRes struct {
	*query.Result
}

type GetListReq struct {
	g.Meta `path:"/${.table.short_name|CaseKebab}/list" method:"get" tags:"[后台]${.table.comment}" summary:"获取${.table.comment}列表"`
${- range $index, $elem := .table.fields}
${- if eq $elem.need_search "true"}
    ${$elem.name|CaseCamel} ${$elem.type} `json:"${$elem.json_name}" in:"query" dc:"${$elem.comment}"`
${- end}
${- end}
}
type GetListRes []*entity.${.table.name|CaseCamel}

type CreateReq struct {
    g.Meta `path:"/${.table.short_name|CaseKebab}/create" method:"post" tags:"[后台]${.table.comment}" summary:"新建${.table.comment}"`
	CreateUpdateBase
}
type CreateRes struct{
    Id int64 `json:"id" dc:"${.table.comment}Id"`
}

type UpdateReq struct {
	g.Meta `path:"/${.table.short_name|CaseKebab}/update" method:"post" tags:"[后台]${.table.comment}" summary:"修改${.table.comment}"`
	Id   int64   `json:"id" v:"required|min:1#请选择要修改的${.table.comment}|请选择要修改的${.table.comment}" dc:"主键ID"`
    CreateUpdateBase
}
type UpdateRes struct{}

type CreateUpdateBase struct {
${- range $index, $elem := .table.fields}
    ${$elem.name|CaseCamel} ${$elem.type} `json:"${$elem.json_name}" v:"required#请输入${$elem.comment_short}" dc:"${$elem.comment}"`
${- end}
}

type DeleteReq struct {
	g.Meta `path:"/${.table.short_name|CaseKebab}/delete/{Id}" method:"post" tags:"[后台]${.table.comment}" summary:"删除${.table.comment}"`
	Id  int64 `in:"path" v:"required|min:1#请选择要删除的${.table.comment}|请选择要删除的${.table.comment}" dc:"主键ID"`
}
type DeleteRes struct{}

type GetInfoReq struct {
	g.Meta `path:"/${.table.short_name|CaseKebab}/info/{Id}" method:"get" tags:"[后台]${.table.comment}" summary:"获取${.table.comment}详情"`
	Id  int64 `in:"path" v:"required|min:1#请选择要查看的${.table.comment}|请选择要查看的${.table.comment}" dc:"主键ID"`
}
type GetInfoRes struct {
    *entity.${.table.name|CaseCamel}
}