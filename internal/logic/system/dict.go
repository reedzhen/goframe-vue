package system

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-vben/internal/codes"
	"goframe-vben/internal/dao"
	"goframe-vben/internal/library/contexts"
	"goframe-vben/internal/library/query"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/model/entity"
	"goframe-vben/internal/service"
)

type sDict struct{}

func init() {
	service.RegisterDict(NewDict())
}

func NewDict() *sDict {
	return &sDict{}
}

// Page 字典分页
func (s *sDict) Page(ctx context.Context, in dto.DictPageInput) (*query.Result, error) {
	var items []*entity.SysDict
	res, err := query.Page(dao.SysDict.Ctx(ctx), &in, &items)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Create 字典新增
func (s *sDict) Create(ctx context.Context, in dto.DictCreateInput) (err error) {
	var dict entity.SysDict
	if err = gconv.Struct(in, &dict); err != nil {
		return
	}
	dict.CreatedBy = contexts.GetUserId(ctx)
	_, err = dao.SysDict.Ctx(ctx).Data(dict).OmitEmptyData().Insert()
	return
}

// Update 字典编辑
func (s *sDict) Update(ctx context.Context, in dto.DictUpdateInput) (err error) {
	param := gconv.Map(in)
	param[dao.SysDict.Columns().UpdatedBy] = contexts.GetUserId(ctx)
	_, err = dao.SysDict.Ctx(ctx).WherePri(in.Id).FieldsEx(dao.SysDict.Columns().Id).Data(param).Update()
	return
}

// Delete 字典删除
func (s *sDict) Delete(ctx context.Context, id int64) (err error) {
	cnt, err := dao.SysDictData.Ctx(ctx).Where(dao.SysDictData.Columns().DictId, id).Count()
	if err != nil {
		return
	}
	if cnt > 0 {
		return gerror.NewCode(codes.CodeDictHasChild)
	}
	_, err = dao.SysDict.Ctx(ctx).Where(dao.SysDict.Columns().Id, id).Delete()
	return
}

// GetList 字典列表
func (s *sDict) GetList(ctx context.Context) (out []*entity.SysDict, err error) {
	db := dao.SysDict.Ctx(ctx)
	err = db.Ctx(ctx).Scan(&out)
	return
}
