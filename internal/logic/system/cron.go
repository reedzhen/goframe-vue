package system

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-vben/internal/codes"
	"goframe-vben/internal/consts"
	"goframe-vben/internal/dao"
	"goframe-vben/internal/library/cron"
	"goframe-vben/internal/library/query"
	"goframe-vben/internal/model/do"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/model/entity"
	"goframe-vben/internal/service"
)

// sCron 定时任务
type sCron struct{}

func init() {
	service.RegisterCron(NewCron())
}

func NewCron() *sCron {
	return &sCron{}
}

// Page 分页
func (s *sCron) Page(ctx context.Context, in dto.CronPageInput) (out *query.Result, err error) {
	var items []*entity.SysCron
	return query.Page(dao.SysCron.Ctx(ctx), &in, &items)
}

// Create 新增
func (s *sCron) Create(ctx context.Context, in dto.CronCreateInput) (err error) {
	if in.ApiParam == "" {
		in.ApiParam = "{}"
	}
	if in.ApiHeader == "" {
		in.ApiHeader = "{}"
	}
	_, err = dao.SysCron.Ctx(ctx).Data(in).Insert()
	return
}

// Update 编辑
func (s *sCron) Update(ctx context.Context, in dto.CronUpdateInput) error {
	c, err := s.GetInfo(ctx, in.Id)
	if err != nil {
		return err
	}

	var cronDo do.SysCron
	if err = gconv.Struct(in, &cronDo); err != nil {
		return err
	}
	if _, err = dao.SysCron.Ctx(ctx).WherePri(in.Id).FieldsEx(dao.SysCron.Columns().Id).Data(cronDo).Update(); err != nil {
		return err
	}

	if c.Status == string(consts.CronStatusActive) {
		// 获取编辑后的数据
		if err = gconv.Struct(in, &c); err != nil {
			return err
		}
		// 重新启动，比如修改了策略
		if err = cron.ResetStart(c); err != nil {
			return err
		}
		// 控制台打印
		cron.PrintCronStatus()
	}

	return nil
}

// UpdateLastRunAt 编辑最新执行时间
func (s *sCron) UpdateLastRunAt(ctx context.Context, in dto.CronUpdateRunAtInput) error {
	cronDo := do.SysCron{
		NextRunAt: in.NextRunAt,
		LastRunAt: in.LastRunAt,
	}
	// 单次和多次执行的任务更新状态
	if in.Policy == consts.CronPolicyOnce || in.Policy == consts.CronPolicyTimes {
		cronDo.Status = consts.CronStatusInactive
	}
	if _, err := dao.SysCron.Ctx(ctx).WherePri(in.Id).Data(cronDo).Update(); err != nil {
		return err
	}
	return nil
}

// ExecuteOnce 在线执行一次
func (s *sCron) ExecuteOnce(ctx context.Context, cronId int64) error {
	c, err := s.GetInfo(ctx, cronId)
	if err != nil {
		return err
	}

	return cron.Execute(gctx.New(), c)
}

// Delete 删除
func (s *sCron) Delete(ctx context.Context, id int64) error {
	c, err := s.GetInfo(ctx, id)
	if err != nil {
		return err
	}

	if err := dao.SysCron.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if _, err := dao.SysCron.Ctx(ctx).Where(dao.SysCron.Columns().Id, id).Delete(); err != nil {
			return err
		}

		if err := service.CronRecord().DeleteByCronId(ctx, id); err != nil {
			return err
		}

		if err := cron.RemoveOne(c); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

// GetInfo 详情
func (s *sCron) GetInfo(ctx context.Context, id int64) (out *entity.SysCron, err error) {
	if err = dao.SysCron.Ctx(ctx).WherePri(id).Scan(&out); err != nil {
		return
	}
	if out == nil {
		return nil, gerror.NewCode(codes.CodeCronNotFound)
	}
	return
}

// GetList 列表
func (s *sCron) GetList(ctx context.Context, in dto.CronGetListInput) (out []*entity.SysCron, err error) {
	db := dao.SysCron.Ctx(ctx)

	if in.Title != "" {
		db = db.Where("title", in.Title)
	}
	if in.Tag != "" {
		db = db.Where("tag", in.Tag)
	}
	err = db.Scan(&out)
	return
}

// Start 启动所有cron
func (s *sCron) Start(ctx context.Context) {
	defer func() {
		cron.PrintCronStatus()
	}()

	// 获取所有可用的任务
	var list []*entity.SysCron
	if err := dao.SysCron.Ctx(ctx).Where("status", consts.CronStatusActive).Order("sort asc,id desc").Scan(&list); err != nil {
		return
	}
	if len(list) == 0 {
		cron.Log().Debug(ctx, "❗️暂无需要执行的 Cron")
		return
	}

	if err := cron.StartAll(list); err != nil {
		cron.Log().Error(ctx, err)
		return
	}

	cron.Log().Debug(ctx, "🎉 [Cron] start...")
}

// Remove 移除所有cron
func (s *sCron) Remove() {
	cron.RemoveAll()
}

// ChangeStatus 修改状态
func (s *sCron) ChangeStatus(ctx context.Context, in dto.CronChangeStatusInput) error {
	defer func() {
		cron.PrintCronStatus()
	}()

	c, err := s.GetInfo(ctx, in.Id)
	if err != nil {
		return err
	}

	if _, err = dao.SysCron.Ctx(ctx).WherePri(in.Id).Data(do.SysCron{
		Status:    in.Status,
		UpdatedBy: in.UpdatedBy,
	}).Update(); err != nil {
		return err
	}

	// 如果将状态置为进行中
	if in.Status == consts.CronStatusActive {
		return cron.ResetStart(c)
	}

	// 将状态置为已结束
	return cron.Stop(c)
}
