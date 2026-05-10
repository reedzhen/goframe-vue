package system

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"goframe-vben/internal/consts"
	"goframe-vben/internal/dao"
	"goframe-vben/internal/library/cache"
	"goframe-vben/internal/library/gvorm"
	"goframe-vben/internal/library/query"
	"goframe-vben/internal/model/do"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/model/entity"
	"goframe-vben/internal/service"
	"net/url"
	"strings"
	"time"
)

// sTenant 租户
type sTenant struct{}

func init() {
	service.RegisterTenant(NewTenant())
}

func NewTenant() *sTenant {
	return &sTenant{}
}

// GetPage 获取租户分页
func (s *sTenant) GetPage(ctx context.Context, in dto.TenantPageInput) (res *query.Result, err error) {
	items := make([]*entity.SysTenant, 0)
	return query.Page(dao.SysTenant.Ctx(ctx), &in, &items)
}

// GetList 获取租户列表
func (s *sTenant) GetList(ctx context.Context, in dto.TenantGetListInput) ([]*entity.SysTenant, error) {
	var (
		db  = dao.SysTenant.Ctx(ctx)
		out = make([]*entity.SysTenant, 0)
	)

	if in.TenantName != "" {
		db = db.Where("tenant_name", in.TenantName)
	}
	if in.LinkMan != "" {
		db = db.Where("link_man", in.LinkMan)
	}
	if in.LinkPhone != "" {
		db = db.Where("link_phone", in.LinkPhone)
	}
	if in.PackageId > 0 {
		db = db.Where("package_id", in.PackageId)
	}
	if err := db.Scan(&out); err != nil {
		return nil, err
	}

	return out, nil
}

// Create 新增租户
func (s *sTenant) Create(ctx context.Context, in dto.TenantCreateInput) error {
	// 校验租户名称是否重复
	if err := gvorm.CheckFieldExist(dao.SysTenant.Ctx(ctx), dao.SysTenant.Columns().TenantName, in.TenantName); err != nil {
		return gerror.New("租户已存在")
	}

	// 校验租户域名是否重复
	if err := gvorm.CheckFieldExist(dao.SysTenant.Ctx(ctx), dao.SysTenant.Columns().Website, in.Website); err != nil {
		return gerror.New("域名已存在")
	}

	// 校验套餐被禁用
	tPackage, err := service.TenantPackage().GetInfo(ctx, in.PackageId)
	if err != nil {
		return err
	}
	if tPackage.Status == consts.TenantPackageStatusDisable {
		return gerror.New("套餐已被禁用")
	}
	if tPackage.MenuIds == "" {
		return gerror.New("套餐权限为空")
	}

	return dao.SysTenant.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 创建租户
		newTenantId, err := dao.SysTenant.Ctx(ctx).Data(in).InsertAndGetId()
		if err != nil {
			return err
		}

		// 创建角色
		newRoleId, err := service.Role().Create(ctx, dto.RoleCreateInput{
			//TenantId: newTenantId,
			Name: "租户管理员",
			//Code:      "tenant_admin", // 暂时不写，不然新增租户role code 会重复
			Remark:    "系统自动生成",
			Status:    int(consts.RoleStatusOk),
			CreatedBy: in.CreatedBy,
		})
		if err != nil {
			return err
		}

		// 给角色分配套餐权限
		menuIds := gconv.Int64s(gstr.Split(tPackage.MenuIds, ","))
		if err := service.Role().SaveRoleMenuList(ctx, dto.RoleSaveMenuInput{
			//TenantId: newTenantId,
			RoleId:  newRoleId,
			MenuIds: menuIds,
		}); err != nil {
			return err
		}

		// 创建用户，并分配角色
		pass := in.LinkPhone[len(in.LinkPhone)-6:]
		newUserId, err := service.User().Create(ctx, dto.UserCreateInput{
			ParentId:  0,
			Password:  pass,
			TenantId:  newTenantId,
			CreatedBy: in.CreatedBy,
			UserCreateUpdateBase: dto.UserCreateUpdateBase{
				Username:       grand.S(10),
				Nickname:       in.LinkMan,
				Phone:          in.LinkPhone,
				RoleId:         newRoleId,
				OrganizationId: 0,
			},
		})
		if err != nil {
			return err
		}

		// 修改租户的管理员
		if _, err := dao.SysTenant.Ctx(ctx).WherePri(newTenantId).Data(do.SysTenant{LinkId: newUserId}).Update(); err != nil {
			return err
		}

		return nil
	})
}

// Update 编辑租户
func (s *sTenant) Update(ctx context.Context, in dto.TenantUpdateInput) error {
	// 获取租户
	tenant, err := s.GetInfo(ctx, in.Id)
	if err != nil {
		return err
	}

	param := gconv.Map(in)
	if _, err := dao.SysTenant.Ctx(ctx).WherePri(in.Id).FieldsEx(dao.SysTenant.Columns().Id).Data(param).Update(); err != nil {
		return err
	}

	// 删除租户redis缓存
	if err := s.DeleteTenantCache(ctx, tenant); err != nil {
		return err
	}

	return nil
}

// Delete 删除租户
func (s *sTenant) Delete(ctx context.Context, id int64) error {
	tenant, err := s.GetInfo(ctx, id)
	if err != nil {
		return err
	}

	if cnt := service.Organization().GetCountByTenantId(ctx, id); cnt > 0 {
		return gerror.New("该租户下存在机构，删除驳回")
	}

	if err := dao.SysTenant.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除角色and角色菜单绑定
		if err := service.Role().DeleteByTenantId(ctx, id); err != nil {
			return err
		}
		// 删除用户
		if err := service.User().DeleteByTenantId(ctx, id); err != nil {
			return err
		}
		// 删除租户
		if _, err := dao.SysTenant.Ctx(ctx).Where(dao.SysTenant.Columns().Id, id).Delete(); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}

	// 删除租户redis缓存
	if err := s.DeleteTenantCache(ctx, tenant); err != nil {
		return err
	}

	return nil
}

// DeleteTenantCache 删除租户redis缓存
func (s *sTenant) DeleteTenantCache(ctx context.Context, tenant *entity.SysTenant) error {
	if _, err := cache.Instance().Remove(ctx, fmt.Sprintf("SelectCache:"+consts.CacheTenantDomainKey, tenant.Website)); err != nil {
		return err
	}
	if _, err := cache.Instance().Remove(ctx, fmt.Sprintf("SelectCache:"+consts.CacheTenantIdKey, tenant.Id)); err != nil {
		return err
	}
	return nil
}

// GetInfo 获取租户详情
func (s *sTenant) GetInfo(ctx context.Context, id int64) (out *entity.SysTenant, err error) {
	if err = dao.SysTenant.Ctx(ctx).WherePri(id).Scan(&out); err != nil {
		return
	}
	if out == nil {
		return nil, gerror.New("租户不存在")
	}
	return
}

// GetInfoWithCache 获取租户详情
func (s *sTenant) GetInfoWithCache(r *ghttp.Request) (out *entity.SysTenant, err error) {
	ctx := r.GetCtx()
	tenantId := gconv.Int64(r.GetHeader(consts.TenantHeaderKey))

	if tenantId > 0 {
		// 通过租户编号获取并缓存租户信息
		if err = dao.SysTenant.Ctx(ctx).Cache(gdb.CacheOption{
			Duration: 12 * time.Hour,
			Name:     fmt.Sprintf(consts.CacheTenantIdKey, tenantId),
		}).WherePri(tenantId).Scan(&out); err != nil {
			return
		}
	} else {
		// 解析请求的域名
		host := s.getRequestHost(r)
		if err = dao.SysTenant.Ctx(ctx).Cache(gdb.CacheOption{
			Duration: 12 * time.Hour,
			Name:     fmt.Sprintf(consts.CacheTenantDomainKey, host),
		}).Where("website", strings.TrimSpace(host)).Scan(&out); err != nil {
			return
		}
	}

	return
}

// getRequestHost 解析请求的域名
func (s *sTenant) getRequestHost(r *ghttp.Request) string {
	referer, err := url.Parse(r.Request.Referer())
	if err != nil {
		return ""
	}
	return referer.Host
}
