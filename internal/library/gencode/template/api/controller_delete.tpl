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

// Delete 删除
func (c *Controller${.table.short_name|UcFirstCaseLower}) Delete(ctx context.Context, req *${.table.short_name|CaseLower}.DeleteReq) (res *${.table.short_name|CaseLower}.DeleteRes, err error) {
	return nil, service.${.table.short_name|CaseCamel}().Delete(ctx, req.Id)
}
