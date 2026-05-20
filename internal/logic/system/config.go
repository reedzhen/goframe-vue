package system

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmode"
	"goframe-vben/internal/consts"
	"goframe-vben/internal/dao"
	"goframe-vben/internal/library/cache"
	"goframe-vben/internal/library/contexts"
	"goframe-vben/internal/library/query"
	"goframe-vben/internal/model/do"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/model/entity"
	"goframe-vben/internal/service"
)

const (
	configStatusEnabled  = 1
	configStatusDisabled = 2
	configSystemYes      = 1
	configSystemNo       = 2
	defaultValueType     = "string"
	defaultInputType     = 1
)

// sConfig 系统配置中心
type sConfig struct{}

func init() {
	service.RegisterConfig(NewConfig())
}

func NewConfig() *sConfig {
	return &sConfig{}
}

// getCacheKey 获取配置缓存 key。
func (s *sConfig) getCacheKey(ctx context.Context, moduleCode string) string {
	//mode := g.Cfg().MustGet(ctx, "tenant.mode").String()
	//if mode == gftenant.ModeNone {
	//	return fmt.Sprintf(consts.CacheConfigKey, moduleCode)
	//}
	//return fmt.Sprintf(consts.CacheTenantConfigKey, gftenant.GetTenant(ctx), moduleCode)
	return fmt.Sprintf(consts.CacheConfigKey, moduleCode)
}

// getValueCache 获取模块配置值缓存。
func (s *sConfig) getValueCache(ctx context.Context, moduleCode string) (out *dto.ConfigGetValuesOutput, err error) {
	cacheKey := s.getCacheKey(ctx, moduleCode)
	cacheV, err := cache.Instance().Get(ctx, cacheKey)
	if err != nil {
		return
	}
	if cacheV.IsNil() {
		return nil, nil
	}
	err = cacheV.Scan(&out)
	return
}

// setValueCache 设置模块配置值缓存。
func (s *sConfig) setValueCache(ctx context.Context, moduleCode string, data *dto.ConfigGetValuesOutput) error {
	cacheKey := s.getCacheKey(ctx, moduleCode)
	//ttl := consts.CacheTenantConfigKeyTTL * time.Second
	//mode := g.Cfg().MustGet(ctx, "tenant.mode").String()
	//if mode == gftenant.ModeNone {
	//	ttl = consts.CacheConfigKeyTTL * time.Second
	//}
	ttl := consts.CacheConfigKeyTTL * time.Second
	return cache.Instance().Set(ctx, cacheKey, data, ttl)
}

// deleteValueCache 删除模块配置值缓存。
func (s *sConfig) deleteValueCache(ctx context.Context, moduleCode string) error {
	cacheKey := s.getCacheKey(ctx, moduleCode)
	_, err := cache.Instance().Remove(ctx, cacheKey)
	return err
}

// ModuleList 获取配置模块列表。
func (s *sConfig) ModuleList(ctx context.Context, in dto.ConfigModuleListInput) (out []*entity.SysConfigModule, err error) {
	db := dao.SysConfigModule.Ctx(ctx).OrderAsc(dao.SysConfigModule.Columns().Sort).OrderAsc(dao.SysConfigModule.Columns().Id)
	if keywords := strings.TrimSpace(in.Keywords); keywords != "" {
		like := "%" + keywords + "%"
		db = db.Where("(name LIKE ? OR code LIKE ?)", like, like)
	}
	if in.Status != nil {
		db = db.Where(dao.SysConfigModule.Columns().Status, *in.Status)
	}
	err = db.Scan(&out)
	return
}

// ModuleCreate 新增配置模块。
func (s *sConfig) ModuleCreate(ctx context.Context, in dto.ConfigModuleCreateInput) error {
	code := strings.TrimSpace(in.Code)
	if code == "" {
		return gerror.NewCode(gcode.CodeValidationFailed, "模块编码不能为空")
	}
	if exists, err := s.moduleCodeExists(ctx, code, 0); err != nil {
		return err
	} else if exists {
		return gerror.NewCode(gcode.CodeValidationFailed, "模块编码已存在")
	}
	status := normalizeStatus(in.Status)
	_, err := dao.SysConfigModule.Ctx(ctx).Data(do.SysConfigModule{
		Code:        code,
		Name:        strings.TrimSpace(in.Name),
		Description: strings.TrimSpace(in.Description),
		Sort:        in.Sort,
		Status:      status,
		CreatedBy:   contexts.GetUserId(ctx),
	}).Insert()
	return err
}

// ModuleUpdate 编辑配置模块。
func (s *sConfig) ModuleUpdate(ctx context.Context, in dto.ConfigModuleUpdateInput) error {
	oldModule, err := s.validateModule(ctx, in.Id)
	if err != nil {
		return err
	}
	code := strings.TrimSpace(in.Code)
	if code == "" {
		return gerror.NewCode(gcode.CodeValidationFailed, "模块编码不能为空")
	}
	if exists, err := s.moduleCodeExists(ctx, code, in.Id); err != nil {
		return err
	} else if exists {
		return gerror.NewCode(gcode.CodeValidationFailed, "模块编码已存在")
	}
	_, err = dao.SysConfigModule.Ctx(ctx).WherePri(in.Id).Data(do.SysConfigModule{
		Code:        code,
		Name:        strings.TrimSpace(in.Name),
		Description: strings.TrimSpace(in.Description),
		Sort:        in.Sort,
		Status:      normalizeStatus(in.Status),
		UpdatedBy:   contexts.GetUserId(ctx),
	}).Update()
	if err == nil {
		if oldModule.Code != code {
			_ = s.deleteValueCache(ctx, oldModule.Code)
		}
		_ = s.deleteValueCache(ctx, code)
	}
	return err
}

// ModuleDelete 删除配置模块。
func (s *sConfig) ModuleDelete(ctx context.Context, id int64) error {
	module, err := s.validateModule(ctx, id)
	if err != nil {
		return err
	}
	count, err := dao.SysConfigItem.Ctx(ctx).Where(dao.SysConfigItem.Columns().ModuleId, id).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.NewCode(gcode.CodeValidationFailed, "模块下存在配置项，不能删除")
	}
	if _, err = dao.SysConfigModule.Ctx(ctx).WherePri(id).Delete(); err != nil {
		return err
	}
	_ = s.deleteValueCache(ctx, module.Code)
	return nil
}

// ItemPage 获取配置项分页。
func (s *sConfig) ItemPage(ctx context.Context, in dto.ConfigItemPageInput) (out *query.Result, err error) {
	var items []*entity.SysConfigItem
	db := dao.SysConfigItem.Ctx(ctx).OrderAsc(dao.SysConfigItem.Columns().Sort).OrderAsc(dao.SysConfigItem.Columns().Id)
	return query.Page(db, &in, &items)
}

// ItemCreate 新增配置项。
func (s *sConfig) ItemCreate(ctx context.Context, in dto.ConfigItemCreateInput) error {
	module, err := s.validateModule(ctx, in.ModuleId)
	if err != nil {
		return err
	}
	key := strings.TrimSpace(in.ConfigKey)
	if key == "" {
		return gerror.NewCode(gcode.CodeValidationFailed, "配置键名不能为空")
	}
	if exists, err := s.itemKeyExists(ctx, key, 0); err != nil {
		return err
	} else if exists {
		return gerror.NewCode(gcode.CodeValidationFailed, "配置键名已存在")
	}
	_, err = dao.SysConfigItem.Ctx(ctx).Data(do.SysConfigItem{
		ModuleId:     in.ModuleId,
		Name:         strings.TrimSpace(in.Name),
		ConfigKey:    key,
		ConfigValue:  in.ConfigValue,
		DefaultValue: in.DefaultValue,
		ValueType:    normalizeValueType(in.ValueType),
		InputType:    normalizeInputType(in.InputType),
		InputParams:  strings.TrimSpace(in.InputParams),
		Description:  strings.TrimSpace(in.Description),
		Sort:         in.Sort,
		Status:       normalizeStatus(in.Status),
		IsSystem:     normalizeIsSystem(in.IsSystem),
		CreatedBy:    contexts.GetUserId(ctx),
	}).Insert()
	if err == nil {
		_ = s.deleteValueCache(ctx, module.Code)
	}
	return err
}

// ItemUpdate 编辑配置项。
func (s *sConfig) ItemUpdate(ctx context.Context, in dto.ConfigItemUpdateInput) error {
	if _, err := s.validateItem(ctx, in.Id); err != nil {
		return err
	}
	module, err := s.validateModule(ctx, in.ModuleId)
	if err != nil {
		return err
	}
	key := strings.TrimSpace(in.ConfigKey)
	if key == "" {
		return gerror.NewCode(gcode.CodeValidationFailed, "配置键名不能为空")
	}
	if exists, err := s.itemKeyExists(ctx, key, in.Id); err != nil {
		return err
	} else if exists {
		return gerror.NewCode(gcode.CodeValidationFailed, "配置键名已存在")
	}
	_, err = dao.SysConfigItem.Ctx(ctx).WherePri(in.Id).Data(do.SysConfigItem{
		ModuleId:     in.ModuleId,
		Name:         strings.TrimSpace(in.Name),
		ConfigKey:    key,
		ConfigValue:  in.ConfigValue,
		DefaultValue: in.DefaultValue,
		ValueType:    normalizeValueType(in.ValueType),
		InputType:    normalizeInputType(in.InputType),
		InputParams:  strings.TrimSpace(in.InputParams),
		Description:  strings.TrimSpace(in.Description),
		Sort:         in.Sort,
		Status:       normalizeStatus(in.Status),
		IsSystem:     normalizeIsSystem(in.IsSystem),
		UpdatedBy:    contexts.GetUserId(ctx),
	}).Update()
	if err == nil {
		_ = s.deleteValueCache(ctx, module.Code)
	}
	return err
}

// ItemDelete 删除配置项。
func (s *sConfig) ItemDelete(ctx context.Context, id int64) error {
	item, err := s.validateItem(ctx, id)
	if err != nil {
		return err
	}
	module, err := s.validateModule(ctx, item.ModuleId)
	if err != nil {
		return err
	}
	if _, err = dao.SysConfigItem.Ctx(ctx).WherePri(id).Delete(); err != nil {
		return err
	}
	_ = s.deleteValueCache(ctx, module.Code)
	return nil
}

// GetValuesByModuleCode 根据模块编码获取配置值。
func (s *sConfig) GetValuesByModuleCode(ctx context.Context, moduleCode string) (out *dto.ConfigGetValuesOutput, err error) {
	moduleCode = strings.TrimSpace(moduleCode)
	if moduleCode == "" {
		return nil, gerror.NewCode(gcode.CodeValidationFailed, "模块编码不能为空")
	}
	if gmode.Mode() != gmode.DEVELOP {
		out, err = s.getValueCache(ctx, moduleCode)
		if err != nil || out != nil {
			return
		}
	}
	module, err := s.getModuleByCode(ctx, moduleCode)
	if err != nil || module == nil {
		return out, err
	}
	var items []*entity.SysConfigItem
	err = dao.SysConfigItem.Ctx(ctx).
		Where(dao.SysConfigItem.Columns().ModuleId, module.Id).
		Where(dao.SysConfigItem.Columns().Status, configStatusEnabled).
		OrderAsc(dao.SysConfigItem.Columns().Sort).
		OrderAsc(dao.SysConfigItem.Columns().Id).
		Scan(&items)
	if err != nil {
		return
	}
	out = &dto.ConfigGetValuesOutput{
		ModuleCode: moduleCode,
		Data:       make(map[string]any, len(items)),
	}
	for _, item := range items {
		out.Data[item.ConfigKey] = castConfigValue(item)
	}
	if gmode.Mode() != gmode.DEVELOP {
		err = s.setValueCache(ctx, moduleCode, out)
	}
	return
}

// SaveValuesByModuleCode 保存模块配置值。
func (s *sConfig) SaveValuesByModuleCode(ctx context.Context, in dto.ConfigSaveValuesInput) error {
	module, err := s.getModuleByCode(ctx, in.ModuleCode)
	if err != nil {
		return err
	}
	if module == nil {
		return gerror.NewCode(gcode.CodeValidationFailed, "配置模块不存在")
	}
	if err = dao.SysConfigItem.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		for _, value := range in.Values {
			key := strings.TrimSpace(value.ConfigKey)
			if key == "" {
				continue
			}
			if _, err = dao.SysConfigItem.Ctx(ctx).
				Where(dao.SysConfigItem.Columns().ModuleId, module.Id).
				Where(dao.SysConfigItem.Columns().ConfigKey, key).
				Data(do.SysConfigItem{
					ConfigValue: value.ConfigValue,
					UpdatedBy:   contexts.GetUserId(ctx),
				}).
				Update(); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return err
	}
	return s.deleteValueCache(ctx, module.Code)
}

// GetUpload 获取上传配置。
func (s *sConfig) GetUpload(ctx context.Context) (out *dto.ConfigUploadOutput, err error) {
	res, err := s.GetValuesByModuleCode(ctx, "upload")
	if err != nil || res == nil {
		return
	}
	err = gconv.Scan(res.Data, &out)
	return
}

func (s *sConfig) getModuleByCode(ctx context.Context, code string) (out *entity.SysConfigModule, err error) {
	err = dao.SysConfigModule.Ctx(ctx).
		Where(dao.SysConfigModule.Columns().Code, strings.TrimSpace(code)).
		Where(dao.SysConfigModule.Columns().Status, configStatusEnabled).
		Scan(&out)
	return
}

func (s *sConfig) validateModule(ctx context.Context, id int64) (out *entity.SysConfigModule, err error) {
	if id <= 0 {
		return nil, gerror.NewCode(gcode.CodeValidationFailed, "配置模块不存在")
	}
	err = dao.SysConfigModule.Ctx(ctx).WherePri(id).Scan(&out)
	if err != nil {
		return
	}
	if out == nil {
		return nil, gerror.NewCode(gcode.CodeValidationFailed, "配置模块不存在")
	}
	return
}

func (s *sConfig) validateItem(ctx context.Context, id int64) (out *entity.SysConfigItem, err error) {
	if id <= 0 {
		return nil, gerror.NewCode(gcode.CodeValidationFailed, "配置项不存在")
	}
	err = dao.SysConfigItem.Ctx(ctx).WherePri(id).Scan(&out)
	if err != nil {
		return
	}
	if out == nil {
		return nil, gerror.NewCode(gcode.CodeValidationFailed, "配置项不存在")
	}
	return
}

func (s *sConfig) moduleCodeExists(ctx context.Context, code string, excludeId int64) (bool, error) {
	db := dao.SysConfigModule.Ctx(ctx).Where(dao.SysConfigModule.Columns().Code, code)
	if excludeId > 0 {
		db = db.WhereNot(dao.SysConfigModule.Columns().Id, excludeId)
	}
	count, err := db.Count()
	return count > 0, err
}

func (s *sConfig) itemKeyExists(ctx context.Context, key string, excludeId int64) (bool, error) {
	db := dao.SysConfigItem.Ctx(ctx).Where(dao.SysConfigItem.Columns().ConfigKey, key)
	if excludeId > 0 {
		db = db.WhereNot(dao.SysConfigItem.Columns().Id, excludeId)
	}
	count, err := db.Count()
	return count > 0, err
}

func castConfigValue(item *entity.SysConfigItem) any {
	val := strings.TrimSpace(item.ConfigValue)
	if val == "" {
		val = item.DefaultValue
	}
	switch strings.ToLower(strings.TrimSpace(item.ValueType)) {
	case "int":
		return gconv.Int(val)
	case "int64":
		return gconv.Int64(val)
	case "uint":
		return gconv.Uint(val)
	case "uint64":
		return gconv.Uint64(val)
	case "float", "float64":
		return gconv.Float64(val)
	case "bool":
		return gconv.Bool(val)
	case "[]string", "strings":
		return gconv.Strings(val)
	case "[]int", "ints":
		return gconv.Ints(val)
	default:
		return gconv.String(val)
	}
}

func normalizeStatus(status int) int {
	if status == configStatusDisabled {
		return configStatusDisabled
	}
	return configStatusEnabled
}

func normalizeIsSystem(value int) int {
	if value == configSystemYes {
		return configSystemYes
	}
	return configSystemNo
}

func normalizeValueType(valueType string) string {
	valueType = strings.ToLower(strings.TrimSpace(valueType))
	if valueType == "" {
		return defaultValueType
	}
	return valueType
}

func normalizeInputType(inputType int) int {
	if inputType < 1 || inputType > 6 {
		return defaultInputType
	}
	return inputType
}
