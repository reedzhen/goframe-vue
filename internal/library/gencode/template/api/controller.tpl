package admin

import (
     ${- if ne .table.addon_name ""}
          admin "${.table.module_name}/addons/${.table.addon_name}/api/admin/${.table.short_name}"
       ${- else}
           admin "${.table.module_name}/api/admin/${.table.short_name}"
     ${- end}
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
    "context"
)

var (
	${.table.short_name|CaseCamel} = ControllerAdmin{}
)

// ControllerAdmin ${.table.comment}
type ControllerAdmin struct{}

// Create 新建
func (c *ControllerAdmin) Create(ctx context.Context, req *admin.CreateReq) (res *admin.CreateRes, err error) {
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
    return &admin.CreateRes{Id: lastInsertId}, nil
}

// Delete 删除
func (c *ControllerAdmin) Delete(ctx context.Context, req *admin.DeleteReq) (res *admin.DeleteRes, err error) {
	return nil, service.${.table.short_name|CaseCamel}().Delete(ctx, req.Id)
}

// GetInfo 详情
func (c *ControllerAdmin) GetInfo(ctx context.Context, req *admin.GetInfoReq) (res *admin.GetInfoRes, err error) {
    data, err := service.${.table.short_name|CaseCamel}().GetInfo(ctx, req.Id)
    if err != nil {
        return
    }

    return &admin.GetInfoRes{${.table.name|CaseCamel}: data}, nil
}

// GetList 列表
func (c *ControllerAdmin) GetList(ctx context.Context, req *admin.GetListReq) (res *admin.GetListRes, err error) {
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

    return (*admin.GetListRes)(&list), nil
}

// GetPage 分页
func (c *ControllerAdmin) GetPage(ctx context.Context, req *admin.GetPageReq) (res *admin.GetPageRes, err error) {
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
    return &admin.GetPageRes{Result: data}, err
}

// Update 修改
func (c *ControllerAdmin) Update(ctx context.Context, req *admin.UpdateReq) (res *admin.UpdateRes, err error) {
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
