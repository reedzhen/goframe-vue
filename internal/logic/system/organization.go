package system

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"goframe-vben/internal/codes"
	"goframe-vben/internal/dao"
	"goframe-vben/internal/library/tree"
	"goframe-vben/internal/model/do"
	"goframe-vben/internal/model/dto"
	"goframe-vben/internal/model/entity"
	"goframe-vben/internal/service"
	"strings"
)

type sOrganization struct{}

func init() {
	service.RegisterOrganization(NewOrganization())
}

func NewOrganization() *sOrganization {
	return &sOrganization{}
}

// Create 新建组织
func (s *sOrganization) Create(ctx context.Context, in dto.OrganizationCreateInput) (out int64, err error) {
	db := dao.SysOrganization.Ctx(ctx)

	// 生成组织编号
	in.Code = s.genOrgCode(ctx)

	// 生成树
	if in.Level, in.Tree, err = tree.GenerateChild(db, in.ParentId); err != nil {
		return
	}

	return db.Data(in).InsertAndGetId()
}

// GenOrderSn 生成组织编号
func (s *sOrganization) genOrgCode(ctx context.Context) string {
	code := fmt.Sprintf("%v%v", gtime.Now().Format("ymd"), grand.Digits(6))
	count, _ := dao.SysOrganization.Ctx(ctx).Where(dao.SysOrganization.Columns().Code, code).Count()
	if count > 0 {
		return s.genOrgCode(ctx)
	}
	return code
}

// Update 编辑组织
func (s *sOrganization) Update(ctx context.Context, in dto.OrganizationUpdateInput) error {
	if in.ParentId == in.Id {
		return gerror.NewCode(codes.CodeOrgParentSelf)
	}

	// 获取组织明细
	org, err := s.ValidateExists(ctx, in.Id)
	if err != nil {
		return err
	}

	// 判断 targetId 是否是 id 的下级节点。
	isChild, err := tree.IsDescendant(dao.SysOrganization.Ctx(ctx), org.Id, org.Tree, in.ParentId)
	if err != nil {
		return err
	}
	if isChild {
		return gerror.NewCode(codes.CodeOrgChildAsParent)
	}

	updateData := do.SysOrganization{
		ParentId:  in.ParentId,
		Code:      in.Code,
		Name:      in.Name,
		FullName:  in.FullName,
		Type:      in.Type,
		Status:    in.Status,
		LinkId:    in.LinkId,
		LinkMan:   in.LinkMan,
		LinkPhone: in.LinkPhone,
		Remark:    in.Remark,
		UpdatedBy: in.UpdatedBy,
	}

	return dao.SysOrganization.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if org.ParentId != in.ParentId {
			// 更新当前组织的tree和level
			level, treePath, err := tree.GenerateChild(dao.SysOrganization.Ctx(ctx), in.ParentId)
			if err != nil {
				return err
			}
			// 更新子级组织的tree和level
			if err := tree.RebuildDescendants(dao.SysOrganization.Ctx(ctx), org.Id, level, treePath); err != nil {
				return err
			}

			updateData.Level = level
			updateData.Tree = treePath
		}

		if _, err := dao.SysOrganization.Ctx(ctx).WherePri(in.Id).Data(updateData).Update(); err != nil {
			return err
		}

		return nil
	})
}

// Delete 删除组织
func (s *sOrganization) Delete(ctx context.Context, id int64) error {

	childCnt, err := dao.SysOrganization.Ctx(ctx).Where(dao.SysOrganization.Columns().ParentId, id).Count()
	if err != nil {
		return err
	}
	if childCnt > 0 {
		return gerror.NewCode(codes.CodeOrgHasChild)
	}

	cnt, err := dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().OrganizationId, id).Count()
	if err != nil {
		return err
	}
	if cnt > 0 {
		return gerror.NewCode(codes.CodeOrgInUse)
	}

	// 删除组织
	if _, err = dao.SysOrganization.Ctx(ctx).WherePri(id).Delete(); err != nil {
		return err
	}

	return nil
}

// GetList 获取组织列表
func (s *sOrganization) GetList(ctx context.Context, in dto.OrganizationGetListInput) ([]*entity.SysOrganization, error) {
	var (
		out = make([]*entity.SysOrganization, 0) // 返回值
		db  = dao.SysOrganization.Ctx(ctx)
	)

	if in.Name != "" {
		db = db.WhereLike("name", "%"+strings.TrimSpace(in.Name)+"%")
	}
	if err := db.Scan(&out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetInfo 详情
func (s *sOrganization) GetInfo(ctx context.Context, id int64) (out *entity.SysOrganization, err error) {
	if err = dao.SysOrganization.Ctx(ctx).WherePri(id).Scan(&out); err != nil {
		return
	}
	return
}

// ValidateExists 获取并校验详情
func (s *sOrganization) ValidateExists(ctx context.Context, id int64) (out *entity.SysOrganization, err error) {
	if err = dao.SysOrganization.Ctx(ctx).WherePri(id).Scan(&out); err != nil {
		return
	}
	if out == nil {
		return nil, gerror.NewCode(codes.CodeOrgNotFound)
	}
	return
}

// GetCountByTenantId 获取数量
func (s *sOrganization) GetCountByTenantId(ctx context.Context, tenantId int64) int {
	cnt, err := dao.SysOrganization.Ctx(ctx).Where("tenant_id", tenantId).Count()
	if err != nil {
		return 0
	}
	return cnt
}
