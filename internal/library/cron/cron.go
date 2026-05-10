package cron

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/robfig/cron/v3"
	"goframe-vben/internal/consts"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/model/entity"
	"goframe-vben/internal/service"
	"strings"
	"time"
)

// cronRes 调用接口返回结果
type cronRes struct {
	Code    int         `json:"code"`    // 错误码((0:成功, 1:失败, >1:错误码))
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据
}

// getCronName 生成Cron名称
func getCronName(cronId int64) string {
	return fmt.Sprintf("cron@%d", cronId)
}

// StartOne 启动单个任务(修改单个cron时，需重新启动)
func StartOne(in *entity.SysCron) (err error) {
	if in == nil {
		return
	}

	c := gcron.Search(getCronName(in.Id))
	if c != nil {
		c.Start()
	}
	// 找不到就添加新的cron
	return StartAll([]*entity.SysCron{in})
}

// StartAll 启动所有任务
func StartAll(list []*entity.SysCron) (err error) {
	ctx := gctx.New()
	for _, c := range list {
		cronName := getCronName(c.Id)
		if gcron.Search(cronName) != nil {
			continue
		}

		// linux crontab pattern 转框架 6段式
		pattern := c.Pattern
		if !strings.Contains(c.Pattern, "@") {
			pattern = fmt.Sprintf("# %s", c.Pattern)
		}

		var t *gcron.Entry
		switch c.Policy {
		case consts.CronPolicyParallel:
			t, err = gcron.Add(ctx, pattern, GenExecuteFun(c), cronName)
		case consts.CronPolicySingle:
			t, err = gcron.AddSingleton(ctx, pattern, GenExecuteFun(c), cronName)
		case consts.CronPolicyOnce:
			t, err = gcron.AddOnce(ctx, pattern, GenExecuteFun(c), cronName)
		case consts.CronPolicyTimes:
			if c.Count < 1 {
				return gerror.Newf("定时任务[%s]执行次数不能小于1", c.Title)
			}
			t, err = gcron.AddTimes(ctx, pattern, c.Count, GenExecuteFun(c), cronName)
		default:
			return gerror.Newf("当前策略不支持, %v", c.Policy)
		}
		if err != nil {
			return err
		}
		if t == nil {
			return gerror.New("启动任务失败")
		}
	}

	return nil
}

// Execute 立即执行一次某个任务
func Execute(ctx context.Context, in *entity.SysCron) error {
	st := gtime.Now()
	cronRecordDo := dto.CronRecordCreateInput{
		CronId: in.Id,
		Status: consts.CronRecordStatusSuccess,
		Remark: "",
	}

	if err := g.Try(ctx, func(ctx context.Context) {
		// 获取数据库里请求头转成map类型
		h, err := gjson.LoadContent(in.ApiHeader.MustToJson())
		if err != nil {
			panic(err)
		}
		headerMap := make(map[string]string)
		if !h.IsNil() {
			if err := h.Scan(&headerMap); err != nil {
				panic(err)
			}
		}

		p, err := gjson.LoadContent(in.ApiParam.MustToJson())
		if err != nil {
			panic(err)
		}

		// 请求接口
		resp, err := g.Client().ContentJson().Timeout(10*time.Second).SetHeaderMap(headerMap).Post(ctx, in.ApiUrl, p.Map())
		if err != nil {
			panic(err)
		}
		b := resp.ReadAll()
		var res *cronRes
		if err := gjson.DecodeTo(b, &res); err != nil {
			panic(err)
		}
		if res == nil {
			panic(gerror.New("接口调用失败"))
		}
		if res.Code > 0 {
			panic(gerror.New(res.Message))
		}
		cronRecordDo.Output = gjson.New(res).MustToJsonString()
	}); err != nil {
		cronRecordDo.Status = consts.CronRecordStatusFailure
		cronRecordDo.Output = err.Error()
		Log().Error(ctx, "定时任务执行失败1", err)
	}

	// 1.记录 cron 日志
	milliseconds := gtime.Now().Sub(st).Milliseconds() // 计算耗时
	cronRecordDo.SpendTime = int(milliseconds)
	if err := service.CronRecord().Create(ctx, cronRecordDo); err != nil {
		Log().Errorf(ctx, "定时任务记录失败, err:%s", err.Error())
		return err
	}

	// 2.更新 cron 下次执行时间
	nextTime, err := GetCronNextTime(in.Pattern, time.Now())
	if err != nil {
		Log().Errorf(ctx, "定时任务获取下一次执行时时间失败, err:%s", err.Error())
		return err
	}
	// 编辑最新执行时间
	if err := service.Cron().UpdateLastRunAt(ctx, dto.CronUpdateRunAtInput{
		Id:        in.Id,
		Policy:    in.Policy,
		NextRunAt: gtime.NewFromTime(nextTime),
		LastRunAt: gtime.Now(),
	}); err != nil {
		return err
	}
	Log().Debugf(ctx, "⏱️ 定时任务[%s]执行完毕, 耗时:%vms", in.Title, milliseconds)

	return nil
}

// GenExecuteFun 生成执行过程
func GenExecuteFun(in *entity.SysCron) func(ctx context.Context) {
	return func(ctx context.Context) {
		if err := Execute(ctx, in); err != nil {
			Log().Errorf(ctx, "定时任务执行失败 %+v", err)
		}
	}
}

// ResetStart 重置任务
func ResetStart(in *entity.SysCron) (err error) {
	if err = Stop(in); err != nil {
		return
	}
	if err = RemoveOne(in); err != nil {
		return
	}
	return StartOne(in)
}

// Stop 停止单个任务
func Stop(in *entity.SysCron) (err error) {
	if in == nil {
		return
	}

	c := gcron.Search(getCronName(in.Id))
	if c == nil {
		return
	}
	c.Stop()
	return
}

// RemoveOne 删除任务
func RemoveOne(in *entity.SysCron) (err error) {
	if in == nil {
		return
	}
	for _, v := range gcron.Entries() {
		if v.Name == getCronName(in.Id) {
			gcron.Remove(v.Name)
			break
		}
	}
	return
}

// RemoveAll 移除所有任务
func RemoveAll() {
	for _, v := range gcron.Entries() {
		gcron.Remove(v.Name)
		Log().Debugf(context.Background(), "❌ [cron] remove %s", v.Name)
	}
}

// PrintCronStatus 打印所有定时任务
func PrintCronStatus() {
	for k, v := range gcron.Entries() {
		s := ""
		switch v.Status() {
		case gtimer.StatusReady: // Job or Timer is ready for running.
			s = "准备开始"
		case gtimer.StatusRunning: // Job or Timer is already running.
			s = "已经开始"
		case gtimer.StatusStopped: // Job or Timer is stopped.
			s = "暂停"
		case gtimer.StatusClosed: // Job or Timer is closed and waiting to be deleted.
			s = "已关闭等待删除"
		}
		fmt.Println(fmt.Sprintf("⏱️序号:%d", k+1), v.Name, v.RegisterTime.Format("2006-01-02 15:04:05"), fmt.Sprintf("状态:%s", s))
	}
}

// GetCronNextTime 获取下一次执行时间,不支持秒
func GetCronNextTime(cronStr string, t time.Time) (nextTime time.Time, err error) {
	p := cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	s, err := p.Parse(cronStr)
	if err != nil {
		return
	}
	nextTime = s.Next(t)
	return
}

// Log Cron日志
func Log() *glog.Logger {
	return g.Log(consts.LoggerGroupCron)
}
