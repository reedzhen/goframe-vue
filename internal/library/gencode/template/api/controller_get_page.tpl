package admin

import (
    "context"
    "${.table.module_name}/api/admin/${.table.short_name|CaseLower}"
    "strings"
	 ${- if ne .table.addon_name ""}
	  "${.table.module_name}/addons/${.table.addon_name}/model/dto"
   ${- else}
	   "${.table.module_name}/internal/model/dto"
	 ${- end}
	 ${- if ne .table.addon_name ""}
	  "${.table.module_name}/addons/${.table.addon_name}/service"
	 ${- else}
	   "${.table.module_name}/internal/service"
	 ${- end}
)

// GetPage 分页
func (c *Controller${.table.short_name|UcFirstCaseLower}) GetPage(ctx context.Context, req *${.table.short_name|CaseLower}.GetPageReq) (res *${.table.short_name|CaseLower}.GetPageRes, err error) {
	data, err := service.${.table.short_name|CaseCamel}().GetPage(ctx, dto.${.table.short_name|CaseCamel}PageInput{
	   	PageParam: req.PageParam,
        ${- range $index, $elem := .table.fields}
            ${ if eq $elem.need_search "true"}
               ${if eq $elem.type "string"}
                  ${$elem.name|CaseCamel}: strings.TrimSpace(req.${$elem.name|CaseCamel}), // ${$elem.comment}
               ${else}
                  ${$elem.name|CaseCamel}: req.${$elem.name|CaseCamel}, // ${$elem.comment}
               ${end}
            ${end}
        ${- end}
	})
    return &${.table.short_name|CaseLower}.GetPageRes{Result: data}, err
}
