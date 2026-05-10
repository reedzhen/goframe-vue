package system

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/os/gtime"
	"goframe-vben/internal/dao"
	"goframe-vben/internal/library/query"
	"goframe-vben/internal/model/do"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/model/entity"
	"goframe-vben/internal/service"
)

type sOperationRecord struct{}

func init() {
	service.RegisterOperationRecord(NewOperationRecord())
}

func NewOperationRecord() *sOperationRecord {
	return &sOperationRecord{}
}

// Page 操作日志分页
func (s *sOperationRecord) Page(ctx context.Context, in dto.OperationRecordPageInput) (out *query.Result, err error) {
	list := make([]*entity.SysOperationRecord, 0)
	out, err = query.Page(dao.SysOperationRecord.Ctx(ctx).OrderDesc("id"), &in, &list)
	return
}

// AsyncCreate 异步调用操作日志
func (s *sOperationRecord) AsyncCreate(ctx context.Context, in dto.OperationRecordCreateInput) {
	if err := grpool.AddWithRecover(ctx, func(ctx context.Context) {
		if err := s.Create(ctx, in); err != nil {
			g.Log().Error(ctx, err)
		}
	}, func(ctx context.Context, err error) {
		g.Log().Error(ctx, err)
	}); err != nil {
		g.Log().Error(ctx, err)
	}
	return
}

// Create 新建操作日志
func (s *sOperationRecord) Create(ctx context.Context, in dto.OperationRecordCreateInput) (err error) {
	_, err = dao.SysOperationRecord.Ctx(ctx).Data(do.SysOperationRecord{
		UserId:     in.UserId,
		Username:   in.Username,
		Nickname:   in.Nickname,
		Url:        in.Url,
		Method:     in.Method,
		Module:     in.Module,
		Summary:    in.Summary,
		Param:      in.Param,
		JsonResult: in.JsonResult,
		ErrorMsg:   in.ErrorMsg,
		SpendTime:  in.SpendTime,
		TraceId:    in.TraceId,
		Status:     in.Status,
		Platform:   in.Platform,
		UserAgent:  in.UserAgent,
		Ip:         in.Ip,
		Remark:     in.Remark,
	}).Insert()
	return
}

// Delete 删除超过n天的日志
func (s *sOperationRecord) Delete(ctx context.Context, day int) (err error) {
	_, err = dao.SysOperationRecord.Ctx(ctx).WhereLT("created_at", gtime.Now().AddDate(0, 0, -day)).Delete()
	return
}
