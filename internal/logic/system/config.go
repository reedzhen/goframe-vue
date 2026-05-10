package system

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmode"
	"goframe-vben/internal/codes"

	"goframe-vben/internal/consts"
	"goframe-vben/internal/dao"
	"goframe-vben/internal/library/cache"
	"goframe-vben/internal/library/gftenant"
	"goframe-vben/internal/model/do"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/model/entity"
	"goframe-vben/internal/service"
	"strings"
	"time"
)

// sConfig 贴牌配置
type sConfig struct {
}

func init() {
	service.RegisterConfig(NewConfig())
}

func NewConfig() *sConfig {
	return &sConfig{}
}

// getCacheKey 获取缓存key
func (s *sConfig) getCacheKey(ctx context.Context, group string) string {
	mode := g.Cfg().MustGet(ctx, "tenant.mode").String()
	if mode == gftenant.ModeNone {
		return fmt.Sprintf(consts.CacheConfigKey, group)
	}
	return fmt.Sprintf(consts.CacheTenantConfigKey, gftenant.GetTenant(ctx), group)
}

// getGroupCache 获取缓存配置
func (s *sConfig) getGroupCache(ctx context.Context, group string) (out []*entity.SysConfig, err error) {
	if group == "" {
		return nil, gerror.NewCode(codes.CodeConfigGroupEmpty)
	}

	cacheKey := s.getCacheKey(ctx, group)
	cacheV, err := cache.Instance().Get(ctx, cacheKey)
	if err != nil {
		return
	}
	if err = cacheV.Scan(&out); err != nil {
		return
	}
	return
}

// setGroupCache 设置缓存配置
func (s *sConfig) setGroupCache(ctx context.Context, group string, confList []*entity.SysConfig) (err error) {
	cacheKey := s.getCacheKey(ctx, group)

	ttl := consts.CacheTenantConfigKeyTTL * time.Second
	mode := g.Cfg().MustGet(ctx, "tenant.mode").String()
	if mode == gftenant.ModeNone {
		ttl = consts.CacheConfigKeyTTL * time.Second
	}
	return cache.Instance().Set(ctx, cacheKey, confList, ttl)
}

// deleteGroupCache 删除缓存
func (s *sConfig) deleteGroupCache(ctx context.Context, group string) (err error) {
	cacheKey := s.getCacheKey(ctx, group)
	if _, err = cache.Instance().Remove(ctx, cacheKey); err != nil {
		return
	}
	return
}

// GetConfigByGroup 根据分组名称获取配置列表
func (s *sConfig) GetConfigByGroup(ctx context.Context, group string) (out *dto.ConfigGetListOutput, err error) {
	out = &dto.ConfigGetListOutput{
		Group: group,
		Data:  make(g.MapStrAny),
	}

	confList := make([]*entity.SysConfig, 0)
	// 开发环境不走缓存
	if gmode.Mode() != gmode.DEVELOP {
		confList, err = s.getGroupCache(ctx, group)
		if err != nil {
			return nil, err
		}
	}

	if len(confList) == 0 {
		// 获取配置列表
		if err = dao.SysConfig.Ctx(ctx).Where("group", group).Scan(&confList); err != nil {
			return
		}
		if len(confList) == 0 {
			return
		}

		// 设置分组缓存
		if err = s.setGroupCache(ctx, group, confList); err != nil {
			return
		}
	}

	// 类型装换
	for _, v := range confList {
		val := strings.TrimSpace(v.Value)
		if val == "" {
			val = v.DefaultValue
		}
		switch v.Type {
		case "string":
			out.Data[v.Key] = gconv.String(val)
		case "int":
			out.Data[v.Key] = gconv.Int(val)
		case "int64":
			out.Data[v.Key] = gconv.Int64(val)
		case "[]string":
			out.Data[v.Key] = gconv.Strings(val)
		case "[]int":
			out.Data[v.Key] = gconv.Ints(val)
		case "bool":
			out.Data[v.Key] = gconv.Bool(val)
		default:
			out.Data[v.Key] = gconv.String(val)
		}
	}

	return
}

// UpdateByGroup 更新指定分组的配置
func (s *sConfig) UpdateByGroup(ctx context.Context, in dto.ConfigUpdateInput) error {
	for k, v := range in.Data {
		// 更新
		if _, err := dao.SysConfig.Ctx(ctx).Where("`key` = ?", k).Data(do.SysConfig{Key: k, Value: v, Type: "string", Group: in.Group}).OnDuplicate("value", "updated_at").Save(); err != nil {
			return err
		}
	}

	// 删除缓存
	if err := s.deleteGroupCache(ctx, in.Group); err != nil {
		g.Log().Error(ctx, err)
	}

	return nil
}

// GetUpload 获取上传配置
func (s *sConfig) GetUpload(ctx context.Context) (out *dto.ConfigUploadOutput, err error) {
	res, err := s.GetConfigByGroup(ctx, "upload")
	if err != nil || res == nil {
		return
	}

	err = gconv.Scan(res.Data, &out)
	return
}

//// GetPay 获取支付配置
//func (s *sConfig) GetPay(ctx context.Context) (conf *dto.ConfigPayOutput, err error) {
//	res, err := s.GetConfigByGroup(ctx, "pay")
//	if err != nil || res == nil {
//		return
//	}
//
//	err = gconv.Scan(res.Data, &conf)
//	return
//}
//
//// GetBasic 获取基本配置
//func (s *sConfig) GetBasic(ctx context.Context) (conf *dto.ConfigBasicOutput, err error) {
//	res, err := s.GetConfigByGroup(ctx, "basic")
//	if err != nil || res == nil {
//		return
//	}
//
//	err = gconv.Scan(res.Data, &conf)
//	return
//}
//
//// UploadLogo 上传图片
//func (s *sConfig) UploadLogo(ctx context.Context, in dto.ConfigUploadLogoInput) (string, error) {
//	tenantId := gftenant.GetTenant(ctx)
//	// 执行文件上传
//	result, err := service.File().Upload(ctx, model.FileUploadInput{
//		File:       in.File,
//		Dir:        fmt.Sprintf("T%d/%s", tenantId, in.Dir),
//		RandomName: true,
//		//Name:       fmt.Sprintf("logo_%d", tenantId),
//	})
//	if err != nil {
//		return "", gerror.NewCode(gcode.CodeValidationFailed, "上传图片失败失败")
//	}
//
//	return result, nil
//}
//
//func (s *sConfig) getInfoByKey(key string, list []*entity.SysConfig) *entity.SysConfig {
//	if len(list) == 0 {
//		return nil
//	}
//
//	for _, v := range list {
//		if key == v.Key {
//			return v
//		}
//	}
//	return nil
//}
