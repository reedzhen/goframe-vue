package system

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-vben/internal/dao"
	"goframe-vben/internal/library/contexts"
	"goframe-vben/internal/library/query"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/model/entity"
	"goframe-vben/internal/service"
	"strings"
)

type sDictData struct{}

func init() {
	service.RegisterDictData(NewDictData())
}

func NewDictData() *sDictData {
	return &sDictData{}
}

// Page 字典项分页
func (s *sDictData) Page(ctx context.Context, in dto.DictDataPageInput) (*query.Result, error) {
	var items []*entity.SysDictData
	db := dao.SysDictData.Ctx(ctx).OrderAsc("sort")
	res, err := query.Page(db, &in, &items)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Create 字典项新增
func (s *sDictData) Create(ctx context.Context, in dto.DictDataCreateInput) (err error) {
	var dictData entity.SysDictData
	if err = gconv.Struct(in, &dictData); err != nil {
		return
	}
	dictData.CreatedBy = contexts.GetUserId(ctx)
	_, err = dao.SysDictData.Ctx(ctx).Data(dictData).OmitEmptyData().Insert()
	return
}

// Update 字典项编辑
func (s *sDictData) Update(ctx context.Context, in dto.DictDataUpdateInput) (err error) {
	param := gconv.Map(in)
	param[dao.SysDictData.Columns().UpdatedBy] = contexts.GetUserId(ctx)
	_, err = dao.SysDictData.Ctx(ctx).WherePri(in.Id).FieldsEx(dao.SysDictData.Columns().Id).Data(param).Update()
	return
}

// Delete 字典项删除
func (s *sDictData) Delete(ctx context.Context, id int64) (err error) {
	_, err = dao.SysDictData.Ctx(ctx).Where(dao.SysDictData.Columns().Id, id).Delete()
	return
}

// GetList 获取字典项列表
func (s *sDictData) GetList(ctx context.Context, in dto.DictDataGetListInput) (out []*entity.SysDictData, err error) {
	if err = dao.SysDictData.Ctx(ctx).
		Fields("sys_dict_data.*").
		LeftJoin("sys_dict", "sys_dict.id=sys_dict_data.dict_id").
		Where("sys_dict.code = ?", strings.TrimSpace(in.DictCode)).
		OrderAsc("sys_dict_data.sort").
		Scan(&out); err != nil {
		return
	}

	return
}
