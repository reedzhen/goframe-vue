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

// GetList 列表
func (c *Controller${.table.short_name|UcFirstCaseLower}) GetList(ctx context.Context, req *${.table.short_name|CaseLower}.GetListReq) (res *${.table.short_name|CaseLower}.GetListRes, err error) {
    list, err := service.${.table.short_name|CaseCamel}().GetList(ctx, dto.${.table.short_name|CaseCamel}GetListInput{
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
    if err != nil {
    	return nil, err
    }

    return (*${.table.short_name|CaseLower}.GetListRes)(&list), nil
}
