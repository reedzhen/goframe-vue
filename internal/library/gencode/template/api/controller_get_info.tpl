package admin

import (
    "context"
    "${.table.module_name}/api/admin/${.table.short_name|CaseLower}"
	 ${- if ne .table.addon_name ""}
	  "${.table.module_name}/addons/${.table.addon_name}/service"
	 ${- else}
	   "${.table.module_name}/internal/service"
	 ${- end}
)

// GetInfo 详情
func (c *Controller${.table.short_name|UcFirstCaseLower}) GetInfo(ctx context.Context, req *${.table.short_name|CaseLower}.GetInfoReq) (res *${.table.short_name|CaseLower}.GetInfoRes, err error) {
    data, err := service.${.table.short_name|CaseCamel}().GetInfo(ctx, req.Id)
    if err != nil {
        return
    }

    return &${.table.short_name|CaseLower}.GetInfoRes{${.table.name|CaseCamel}: data}, nil
}
