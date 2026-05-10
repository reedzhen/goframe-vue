package ${.table.short_name|CaseLower}

import (
    "${.table.module_name}/internal/dao"
    "${.table.module_name}/internal/library/query"
    "${.table.module_name}/internal/model/entity"
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
    "context"
    "github.com/gogf/gf/v2/errors/gerror"
    "github.com/gogf/gf/v2/util/gconv"
)

// s${.table.short_name|CaseCamel} ${.table.comment}
type s${.table.short_name|CaseCamel} struct{}

func init() {
	service.Register${.table.short_name|CaseCamel}(New())
}

func New() *s${.table.short_name|CaseCamel} {
	return &s${.table.short_name|CaseCamel}{}
}

// GetPage 获取${.table.comment}分页
func (s *s${.table.short_name|CaseCamel}) GetPage(ctx context.Context, in dto.${.table.short_name|CaseCamel}PageInput) (out *query.Result,err error) {
	list := make([]*entity.${.table.name|CaseCamel},0)

	out, err = query.Page(dao.${.table.name|CaseCamel}.Ctx(ctx), &in, &list)
	return
}

// GetList 获取${.table.comment}列表
func (s *s${.table.short_name|CaseCamel}) GetList(ctx context.Context, in dto.${.table.short_name|CaseCamel}GetListInput) ([]*entity.${.table.name|CaseCamel},error) {
	var (
    		db  = dao.${.table.name|CaseCamel}.Ctx(ctx)
    		out = make([]*entity.${.table.name|CaseCamel}, 0)
    	)

    	${- range $index, $elem := .table.fields}
        ${ if eq $elem.need_search "true"}
            ${if eq $elem.type "string"}
               if in.${$elem.name|CaseCamel} != "" {
                  db = db.Where("${$elem.name}", in.${$elem.name|CaseCamel})
               }
            ${else}
               if in.${$elem.name|CaseCamel} > 0 {
                  db = db.Where("${$elem.name}", in.${$elem.name|CaseCamel})
               }
            ${end}
        ${end}
        ${- end}
    	if err := db.Scan(&out); err != nil {
    		return nil, err
    	}

    return out, nil
}

// Create 新建${.table.comment}
func (s *s${.table.short_name|CaseCamel}) Create(ctx context.Context, in dto.${.table.short_name|CaseCamel}CreateInput) (out int64, err error) {
	return dao.${.table.name|CaseCamel}.Ctx(ctx).Data(in).InsertAndGetId()
}

// Update 修改${.table.comment}
func (s *s${.table.short_name|CaseCamel}) Update(ctx context.Context, in dto.${.table.short_name|CaseCamel}UpdateInput) (err error) {
    // 校验${.table.comment}是否存在
	if _, err = s.ValidateExists(ctx, in.Id); err != nil {
		return
	}

	// 执行编辑
	param := gconv.Map(in)
	if _, err = dao.${.table.name|CaseCamel}.Ctx(ctx).WherePri(in.Id).FieldsEx(dao.${.table.name|CaseCamel}.Columns().Id).Data(param).Update();err != nil{
	    return
	}
	return
}

// Delete 删除${.table.comment}
func (s *s${.table.short_name|CaseCamel}) Delete(ctx context.Context, id int64) (err error) {
    // 校验${.table.comment}是否存在
	if _, err = s.ValidateExists(ctx, id); err != nil {
		return
	}

    // 执行删除
	if _, err = dao.${.table.name|CaseCamel}.Ctx(ctx).Where(dao.${.table.name|CaseCamel}.Columns().Id, id).Delete();err != nil{
	    return
	}
	return
}

// GetInfo 获取${.table.comment}详情
func (s *s${.table.short_name|CaseCamel}) GetInfo(ctx context.Context, id int64) (out *entity.${.table.name|CaseCamel}, err error) {
	if err = dao.${.table.name|CaseCamel}.Ctx(ctx).WherePri(id).Scan(&out); err != nil {
       return
    }
	return
}

// ValidateExists 校验${.table.comment}是否存在
func (s *s${.table.short_name|CaseCamel}) ValidateExists(ctx context.Context, id int64) (out *entity.${.table.name|CaseCamel}, err error) {
	if err = dao.${.table.name|CaseCamel}.Ctx(ctx).WherePri(id).Scan(&out); err != nil {
		return
	}
	if out == nil {
		return nil, gerror.New("${.table.comment}不存在")
	}
	return
}