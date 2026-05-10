package dto

import (
    "${.table.module_name}/internal/library/query"
    "github.com/gogf/gf/v2/database/gdb"
${- if eq .table.has_time "true"}
    "github.com/gogf/gf/v2/os/gtime"
${- end}
)

// ${.table.short_name|CaseCamel}PageInput ${.table.comment}分页
type ${.table.short_name|CaseCamel}PageInput struct {
query.PageParam
${- range $index, $elem := .table.fields}
${- if eq $elem.need_search "true"}
    ${$elem.name|CaseCamel} ${$elem.type} // ${$elem.comment}
${- end}
${- end}
}
func (q *${.table.short_name|CaseCamel}PageInput) Cond(m *gdb.Model) *gdb.Model {
	${- range $index, $elem := .table.fields}
    ${ if eq $elem.need_search "true"}
        ${if eq $elem.type "string"}
           if q.${$elem.name|CaseCamel} != "" {
              m = m.Where("${$elem.name}", q.${$elem.name|CaseCamel})
           }
        ${else}
           if q.${$elem.name|CaseCamel} > 0 {
              m = m.Where("${$elem.name}", q.${$elem.name|CaseCamel})
           }
        ${end}
    ${end}
    ${- end}
	return m
}

// ${.table.short_name|CaseCamel}GetListInput ${.table.comment}列表
type ${.table.short_name|CaseCamel}GetListInput struct {
${- range $index, $elem := .table.fields}
${- if eq $elem.need_search "true"}
    ${$elem.name|CaseCamel} ${$elem.type} // ${$elem.comment}
${- end}
${- end}
}

// ${.table.short_name|CaseCamel}CreateUpdateBase ${.table.comment}新建或修改
type ${.table.short_name|CaseCamel}CreateUpdateBase struct {
${- range $index, $elem := .table.fields}
    ${$elem.name|CaseCamel} ${$elem.type}  // ${$elem.comment}
${- end}
}

// ${.table.short_name|CaseCamel}CreateInput ${.table.comment}新建
type ${.table.short_name|CaseCamel}CreateInput struct {
    CreatedBy int64  // 创建人
    ${.table.short_name|CaseCamel}CreateUpdateBase
}

// ${.table.short_name|CaseCamel}UpdateInput ${.table.comment}修改
type ${.table.short_name|CaseCamel}UpdateInput struct {
	Id   int64
	UpdatedBy int64 // 更新人
    ${.table.short_name|CaseCamel}CreateUpdateBase
}
