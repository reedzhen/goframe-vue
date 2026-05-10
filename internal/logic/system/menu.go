package system

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmode"
	"goframe-vben/internal/codes"
	"goframe-vben/internal/consts"
	"goframe-vben/internal/dao"
	"goframe-vben/internal/library/cache"
	"goframe-vben/internal/library/contexts"
	"goframe-vben/internal/model/do"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/model/entity"
	"goframe-vben/internal/service"
	"strings"
	"time"
)

type sMenu struct{}

func init() {
	service.RegisterMenu(NewMenu())
}

func NewMenu() *sMenu {
	return &sMenu{}
}

// GetList 获取菜单列表
func (s *sMenu) GetList(ctx context.Context, in dto.MenuGetListInput) (out []*entity.SysMenu, err error) {
	var (
		db = dao.SysMenu.Ctx(ctx)
	)
	if in.MenuType > 0 {
		db = db.Where("menuType", in.MenuType)
	}
	if in.ParentId > 0 {
		db = db.Where("parentId", in.ParentId)
	}
	if err = db.OrderAsc("sort").Scan(&out); err != nil {
		return
	}

	return
}

// Create 新建菜单
func (s *sMenu) Create(ctx context.Context, in dto.MenuCreateInput) (out int64, err error) {
	//title := in.MenuMeta.Get("title").String()
	data := do.SysMenu{
		ParentId:  in.ParentId,
		Title:     in.Title,
		Path:      in.Path,
		Component: in.Component,
		MenuType:  in.MenuType,
		Sort:      in.Sort,
		Authority: in.Authority,
		Icon:      in.Icon,
		Hide:      in.Hide,
		MenuMeta:  in.MenuMeta,
		CreatedBy: in.CreatedBy,
		ApiPath:   in.ApiPath,
	}

	return dao.SysMenu.Ctx(ctx).Data(data).OmitEmptyData().InsertAndGetId()
}

// ValidateExists 获取并校验详情
func (s *sMenu) ValidateExists(ctx context.Context, id int64) (out *entity.SysMenu, err error) {
	if err = dao.SysMenu.Ctx(ctx).WherePri(id).Scan(&out); err != nil {
		return
	}
	if out == nil {
		return nil, gerror.NewCode(codes.CodeMenuNotFound)
	}
	return
}

// Update 编辑菜单
func (s *sMenu) Update(ctx context.Context, in dto.MenuUpdateInput) (err error) {
	// 获取旧的菜单数据
	oldMenu, err := s.ValidateExists(ctx, in.Id)
	if err != nil {
		return err
	}

	return dao.SysMenu.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		param := gconv.Map(in)
		if _, err = dao.SysMenu.Ctx(ctx).WherePri(in.Id).FieldsEx(dao.SysMenu.Columns().Id).Data(param).Update(); err != nil {
			return err
		}

		// 如果隐藏状态发生变化，并且新状态有效，那么级联更新所有的子节点
		if in.Hide != 0 && oldMenu.Hide != in.Hide {
			childIds := s.getAllChildIds(ctx, in.Id)
			if len(childIds) > 0 {
				if _, err = dao.SysMenu.Ctx(ctx).WhereIn(dao.SysMenu.Columns().Id, childIds).Data(do.SysMenu{Hide: in.Hide}).Update(); err != nil {
					return err
				}
			}
		}
		return nil
	})

}

// getAllChildIds 递归获取所有子孙菜单ID
func (s *sMenu) getAllChildIds(ctx context.Context, parentId int64) []int64 {
	var allMenus []*entity.SysMenu
	_ = dao.SysMenu.Ctx(ctx).Fields(dao.SysMenu.Columns().Id, dao.SysMenu.Columns().ParentId).Scan(&allMenus)

	var childIds []int64
	var findChildren func(pId int64)
	findChildren = func(pId int64) {
		for _, m := range allMenus {
			if m.ParentId == pId {
				childIds = append(childIds, m.Id)
				findChildren(m.Id)
			}
		}
	}
	findChildren(parentId)
	return childIds
}

// Delete 删除菜单
func (s *sMenu) Delete(ctx context.Context, id int64) (err error) {
	cnt, err := dao.SysMenu.Ctx(ctx).Where("parent_id", id).Count()
	if err != nil {
		return err
	}
	if cnt > 0 {
		return gerror.NewCode(codes.CodeMenuHasChild)
	}

	err = dao.SysUser.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 删除菜单
		if _, err = dao.SysMenu.Ctx(ctx).Where(dao.SysMenu.Columns().Id, id).Delete(); err != nil {
			return err
		}
		// 删除菜单和角色的关系
		if _, err = dao.SysRoleMenu.Ctx(ctx).Where(dao.SysRoleMenu.Columns().MenuId, id).Delete(); err != nil {
			return err
		}
		return nil
	})

	return
}

// GetAuthorizedList 获取当前用户拥有的菜单列表
func (s *sMenu) GetAuthorizedList(ctx context.Context) ([]*entity.SysMenu, error) {
	var (
		curr = contexts.GetUser(ctx)
		out  = make([]*entity.SysMenu, 0)
	)

	if curr.IsAdmin { // 超管
		//.Where("find_in_set(?,show_type)", consts.MenuShowTypeSuper)
		if err := dao.SysMenu.Ctx(ctx).Order("sort asc").Scan(&out); err != nil {
			return nil, err
		}
	} else {
		// 获取角色所有菜单id
		menuIdsValue, _ := dao.SysRoleMenu.Ctx(ctx).Where(dao.SysRoleMenu.Columns().RoleId, curr.RoleId).Fields("DISTINCT menu_id").Array()
		// 获取菜单列表
		if err := dao.SysMenu.Ctx(ctx).WherePri(menuIdsValue).Order("sort asc").Scan(&out); err != nil {
			return nil, err
		}
	}

	return out, nil
}

// LoadRolePermissions 加载角色权限到缓存
func (s *sMenu) LoadRolePermissions(ctx context.Context, roleId int64) (err error) {
	apiPathList, err := dao.SysMenu.Ctx(ctx).As("m").Fields("m.api_path").
		LeftJoin("sys_role_menu srm", "m.id=srm.menu_id").
		Where("m.hide", consts.MenuIsHideNo).
		Where("m.menu_type", consts.MenuTypeButton).
		Where("srm.role_id", roleId).
		Where("JSON_LENGTH(m.api_path) > 0"). // 确保api_path字段不为空
		Array()
	if err != nil {
		return err
	}
	if len(apiPathList) == 0 {
		return gerror.NewCode(codes.CodeMenuNoApiAuth)
	}

	// 将权限列表转换为map，提高检索效率
	permMap := make(map[string]bool)
	for _, v := range apiPathList {
		paths := gconv.Strings(v)
		for _, p := range paths {
			if p != "" {
				permMap[p] = true
			}
		}
	}
	// 将权限map存入缓存，设置30分钟过期时间
	_ = cache.Instance().Set(ctx, fmt.Sprintf("perm:%d", roleId), permMap, 30*time.Minute)

	return
}

// ClearRolePermissionCache 清除角色权限缓存
func (s *sMenu) ClearRolePermissionCache(ctx context.Context, roleId int64) (err error) {
	// 删除缓存
	_, err = cache.Instance().Remove(ctx, fmt.Sprintf("perm:%d", roleId))
	return
}

// CheckRoleApiPermission 检查角色是否拥有指定接口权限
func (s *sMenu) CheckRoleApiPermission(ctx context.Context, roleId int64, route, method string) (bool, error) {
	return s.CheckRolePermissionFast(ctx, roleId, fmt.Sprintf("%s:%s", strings.ToLower(method), route))
}

// CheckRolePermissionFast 使用缓存检查用户权限
func (s *sMenu) CheckRolePermissionFast(ctx context.Context, roleId int64, authKey string) (bool, error) {
	// 开发环境走数据库查询，方便开发
	if !gmode.IsDevelop() {
		// 缓存键
		cacheKey := fmt.Sprintf("perm:%d", roleId)

		// 尝试从缓存获取用户权限map
		if v, _ := cache.Instance().Get(ctx, cacheKey); !v.IsEmpty() {
			permMap := v.MapStrAny()
			// 直接从权限map中检查，时间复杂度O(1)
			return gconv.Bool(permMap[authKey]), nil
		}

		// 缓存不存在，先加载权限
		if err := s.LoadRolePermissions(ctx, roleId); err != nil {
			return false, err
		}

		// 重新从缓存获取
		if v, _ := cache.Instance().Get(ctx, cacheKey); !v.IsEmpty() {
			permMap := v.MapStrAny()
			return gconv.Bool(permMap[authKey]), nil
		}
	}

	// 如果还是不存在，回退到数据库查询
	return s.CheckRolePermission(ctx, roleId, authKey)
}

// CheckRolePermission 检查角色是否拥有指定权限
func (s *sMenu) CheckRolePermission(ctx context.Context, roleId int64, authKey string) (bool, error) {
	//cacheKey := fmt.Sprintf("role_perm_%d_%s", roleId, authKey)
	apiPathList, err := dao.SysMenu.Ctx(ctx).As("m").Fields("m.api_path").
		LeftJoin("sys_role_menu srm", "m.id=srm.menu_id").
		//Where("m.api_path", authKey).
		Where("m.hide", consts.MenuIsHideNo).
		//Where("m.menu_type", consts.MenuTypeButton). // 也可以把接口权限挂到菜单上，例如登录日志，就一个单菜单
		Where("srm.role_id", roleId).
		Where("JSON_LENGTH(m.api_path) > 0"). // 确保api_path字段不为空
		Array()
	if err != nil {
		return false, err
	}

	for _, v := range apiPathList {
		paths := gconv.Strings(v)
		for _, p := range paths {
			if p != "" && p == authKey {
				return true, nil
			}
		}
	}

	return false, nil
}
