package admin

import (
    "context"
    "${.table.module_name}/api/admin/${.table.short_name|CaseLower}"
    "strings"
    "${.table.module_name}/internal/library/contexts"
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

// Update 修改
func (c *Controller${.table.short_name|UcFirstCaseLower}) Update(ctx context.Context, req *${.table.short_name|CaseLower}.UpdateReq) (res *${.table.short_name|CaseLower}.UpdateRes, err error) {
    return nil, service.${.table.short_name|CaseCamel}().Update(ctx, dto.${.table.short_name|CaseCamel}UpdateInput{
    Id: req.Id,
    UpdatedBy: contexts.GetUserId(ctx),
    ${.table.short_name|CaseCamel}CreateUpdateBase: dto.${.table.short_name|CaseCamel}CreateUpdateBase{
        ${- range $index, $elem := .table.fields}
            ${- if eq $elem.type "string"}
              ${$elem.name|CaseCamel}: strings.TrimSpace(req.${$elem.name|CaseCamel}), // ${$elem.comment}
            ${- else}
              ${$elem.name|CaseCamel}: req.${$elem.name|CaseCamel}, // ${$elem.comment}
            ${- end}
        ${- end}
    },
    })
}
