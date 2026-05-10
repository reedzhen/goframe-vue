package admin

import (
    "context"
    "${.table.module_name}/api/admin/${.table.short_name|CaseLower}"
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
	 "strings"
)

// Create 新建
func (c *Controller${.table.short_name|UcFirstCaseLower}) Create(ctx context.Context, req *${.table.short_name|CaseLower}.CreateReq) (res *${.table.short_name|CaseLower}.CreateRes, err error) {
    lastInsertId, err :=service.${.table.short_name|CaseCamel}().Create(ctx, dto.${.table.short_name|CaseCamel}CreateInput{
      CreatedBy: contexts.GetUserId(ctx),
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
    if err != nil {
    	return
    }
    return &${.table.short_name|CaseLower}.CreateRes{Id: lastInsertId}, nil
}
