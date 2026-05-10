package system

import (
	"context"
	"goframe-vben/internal/dao"
	"goframe-vben/internal/library/query"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/model/entity"
	"goframe-vben/internal/service"
)

// sCronRecord 定时任务日志
type sCronRecord struct{}

func init() {
	service.RegisterCronRecord(NewCronRecord())
}

func NewCronRecord() *sCronRecord {
	return &sCronRecord{}
}

// Create 新增
func (s *sCronRecord) Create(ctx context.Context, in dto.CronRecordCreateInput) (err error) {
	_, err = dao.SysCronRecord.Ctx(ctx).Data(in).Insert()
	return
}

// Page 分页
func (s *sCronRecord) Page(ctx context.Context, in dto.CronRecordPageInput) (out *query.Result, err error) {
	var items []*entity.SysCronRecord
	return query.Page(dao.SysCronRecord.Ctx(ctx), &in, &items)
}

// DeleteByCronId 删除
func (s *sCronRecord) DeleteByCronId(ctx context.Context, cronId int64) error {
	if _, err := dao.SysCronRecord.Ctx(ctx).Where("cron_id", cronId).Delete(); err != nil {
		return err
	}

	return nil
}
