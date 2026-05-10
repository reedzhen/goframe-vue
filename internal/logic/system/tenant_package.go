package system

import (
	"context"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-vben/api"
	"goframe-vben/internal/consts"
	"goframe-vben/internal/dao"
	"goframe-vben/internal/library/query"
	"goframe-vben/internal/model/do"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/model/entity"
	"goframe-vben/internal/service"
)

// sTenantPackage 租户套餐
type sTenantPackage struct{}

func init() {
	service.RegisterTenantPackage(NewTenantPackage())
}

func NewTenantPackage() *sTenantPackage {
	return &sTenantPackage{}
}

// GetPage 获取租户套餐分页
func (s *sTenantPackage) GetPage(ctx context.Context, in dto.TenantPackagePageInput) (res *query.Result, err error) {
	items := make([]*entity.SysTenantPackage, 0)
	return query.Page(dao.SysTenantPackage.Ctx(ctx), &in, &items)
}

// GetList 获取租户套餐列表
func (s *sTenantPackage) GetList(ctx context.Context, in dto.TenantPackageGetListInput) ([]*entity.SysTenantPackage, error) {
	var (
		db  = dao.SysTenantPackage.Ctx(ctx).Where("status", consts.TenantPackageStatusOk)
		out = make([]*entity.SysTenantPackage, 0)
	)

	if in.Name != "" {
		db = db.Where("name", in.Name)
	}
	if in.Status > 0 {
		db = db.Where("status", in.Status)
	}
	if err := db.Scan(&out); err != nil {
		return nil, err
	}

	return out, nil
}

// Create 新建租户套餐
func (s *sTenantPackage) Create(ctx context.Context, in dto.TenantPackageCreateInput) (err error) {
	_, err = dao.SysTenantPackage.Ctx(ctx).Data(in).Insert()
	return
}

// Update 修改租户套餐
func (s *sTenantPackage) Update(ctx context.Context, in dto.TenantPackageUpdateInput) (err error) {
	param := gconv.Map(in)
	_, err = dao.SysTenantPackage.Ctx(ctx).WherePri(in.Id).FieldsEx(dao.SysTenantPackage.Columns().Id).Data(param).Update()
	return
}

// Delete 删除租户套餐
func (s *sTenantPackage) Delete(ctx context.Context, id int64) error {
	cnt, err := dao.SysTenant.Ctx(ctx).Where(dao.SysTenant.Columns().PackageId, id).Count()
	if err != nil {
		return err
	}
	if cnt > 0 {
		return gerror.New("该租户套餐已被使用，无法删除")
	}

	if _, err := dao.SysTenantPackage.Ctx(ctx).Where(dao.SysTenantPackage.Columns().Id, id).Delete(); err != nil {
		return err
	}
	return nil
}

// GetInfo 获取租户套餐详情
func (s *sTenantPackage) GetInfo(ctx context.Context, id int64) (out *entity.SysTenantPackage, err error) {
	if err = dao.SysTenantPackage.Ctx(ctx).WherePri(id).Scan(&out); err != nil {
		return
	}
	if out == nil {
		return nil, gerror.New("租户套餐不存在")
	}
	return
}

// GetMenuList 获取所有菜单,套餐已包含则checked==true
func (s *sTenantPackage) GetMenuList(ctx context.Context, packageId int64) ([]*api.MenuItem, error) {
	// 获取所有菜单
	menuList, err := service.Menu().GetList(ctx, dto.MenuGetListInput{})
	if err != nil {
		return nil, err
	}
	if len(menuList) == 0 {
		return nil, nil
	}

	// 获取套餐明细
	menuIds := make([]int64, 0)
	if packageId > 0 {
		packageInfo, err := s.GetInfo(ctx, packageId)
		if err != nil {
			return nil, err
		}
		menuIds = gconv.Int64s(gstr.Split(packageInfo.MenuIds, ","))
	}

	// 挂上Checked标记
	menus := make([]*api.MenuItem, 0)
	for _, v := range menuList {
		temp := api.MenuItem{
			Id:      v.Id,
			Checked: false,
			Title:   v.Title,
			//Icon:      v.Icon,
			Path:      v.Path,
			Component: v.Component,
			//Hide:      v.Hide,
			Sort:     v.Sort,
			ParentId: v.ParentId,
		}
		for _, menuId := range menuIds {
			if v.Id == menuId {
				temp.Checked = true
				break
			}
		}
		menus = append(menus, &temp)
	}
	return menus, nil
}

// SyncTenantPermission 给使用此套餐的租户同步权限（套餐可能会加减权限）
func (s *sTenantPackage) SyncTenantPermission(ctx context.Context, id int64) (err error) {
	// 获取套餐信息
	packageInfo, err := s.GetInfo(ctx, id)
	if err != nil {
		return
	}

	// 待插入角色菜单
	insertMenus := make([]*do.SysRoleMenu, 0)
	// 待删除角色菜单
	deleteMenus := make([]*do.SysRoleMenu, 0)
	// 当前套餐的所有权限
	packageMenus := gset.NewIntSetFrom(gconv.Ints(gstr.Split(packageInfo.MenuIds, ",")))

	//// Diff: 差集，属于packageMenus且不属于roleMenus的元素为元素的集合。
	//roleMenus := gset.NewIntSetFrom([]int{1, 2, 3, 4})
	//packageMenus := gset.NewIntSetFrom([]int{1, 3, 4, 5})
	//
	//// 新增了那些权限 只同步给贴牌，不需要同步给贴牌用户
	//tmpMenuIds := packageMenus.Diff(roleMenus).Slice() // 5
	//g.Dump(tmpMenuIds)
	//// 减少了哪些权限 需要同步删除
	//tmpMenuIds2 := roleMenus.Diff(packageMenus).Slice() // 2
	//g.Dump(tmpMenuIds2)

	// 获取使用此套餐的贴牌列表
	tenantList, err := service.Tenant().GetList(ctx, dto.TenantGetListInput{PackageId: id})
	if err != nil {
		return
	}
	for _, v := range tenantList {
		// 获取租户主账号信息
		user, err1 := service.User().ValidateExists(ctx, v.LinkId)
		if err1 != nil {
			return
		}
		// 获取当前租户角色菜单列表
		roleMenuList, err1 := service.Role().GetRoleMenuListByCond(ctx, []int64{user.RoleId})
		if err1 != nil {
			return
		}
		roleMenus := gset.NewIntSetFrom(gconv.Ints(gdb.ListItemValuesUnique(roleMenuList, "MenuId")))
		// 新增了那些权限 只同步给贴牌，不需要同步给贴牌用户
		addMenuIds := packageMenus.Diff(roleMenus).Slice()
		// 减少了哪些权限 贴牌和贴牌用户需要同步删除
		delMenuIds := roleMenus.Diff(packageMenus).Slice()

		for _, menuId := range addMenuIds {
			insertMenus = append(insertMenus, &do.SysRoleMenu{
				TenantId: v.Id,
				RoleId:   user.RoleId,
				MenuId:   menuId,
			})
		}

		for _, menuId := range delMenuIds {
			deleteMenus = append(deleteMenus, &do.SysRoleMenu{
				TenantId: v.Id,
				//RoleId:   user.RoleId,
				MenuId: menuId,
			})
		}

	}

	// 事务
	return dao.SysTenantPackage.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 插入
		if len(insertMenus) > 0 {
			if _, err := dao.SysRoleMenu.Ctx(ctx).Data(insertMenus).Insert(); err != nil {
				return err
			}
		}
		// 删除
		if len(deleteMenus) > 0 {
			for _, v := range deleteMenus {
				if _, err := dao.SysRoleMenu.Ctx(ctx).Where("menu_id", v.MenuId).Where("tenant_id", v.TenantId).Delete(); err != nil {
					return err
				}
			}
		}
		return nil
	})
}
