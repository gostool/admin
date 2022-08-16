package role_menu

import (
	"admin/internal/consts"
	"admin/internal/dao"
	"admin/internal/model"
	"admin/internal/model/serializer"
	"admin/internal/service"
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

type sRoleMenu struct {
}

var logger *glog.Logger

func init() {
	logger = g.Log(consts.LoggerDebug)
	instance := New()
	service.RegisterRoleMenu(instance)
}

func New() *sRoleMenu {
	return &sRoleMenu{}
}

func (s *sRoleMenu) Count(ctx context.Context) (data int, err error) {
	query := g.Map{
		"is_deleted": consts.CREATED,
	}
	return dao.RoleMenu.Ctx(ctx).Count(query)
}

func (s *sRoleMenu) List(ctx context.Context, in model.RoleMenuListInput) (items []*serializer.RoleMenu, err error) {
	query := g.Map{
		"is_deleted": consts.CREATED,
	}
	err = dao.RoleMenu.Ctx(ctx).Fields(model.RoleMenuFields).Page(in.Page, in.PageSize).Where(query).Scan(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *sRoleMenu) InsertAndGetId(ctx context.Context, data g.Map) (id int64, err error) {
	id, err = dao.RoleMenu.Ctx(ctx).InsertAndGetId(data)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (s *sRoleMenu) Create(ctx context.Context, in *model.RoleMenuCreateInput) (id int64, err error) {
	return s.InsertAndGetId(ctx, in.ToMap())
}

func (s *sRoleMenu) update(ctx context.Context, query, data g.Map) (row int64, err error) {
	logger.Debugf(ctx, "data:%v\n", data)
	result, err := dao.RoleMenu.Ctx(ctx).Where(query).Update(data)
	if err != nil {
		return 0, err
	}
	row, err = result.RowsAffected()
	if err != nil {
		return row, err
	}
	if row == 0 {
		return 0, consts.ErrUpdate
	}
	return row, err
}

// Update 执行更新
func (s *sRoleMenu) Update(ctx context.Context, in model.RoleMenuUpdateInput) (row int64, err error) {
	return s.update(ctx, in.ToWhereMap(), in.ToMap())
}

// Detail 执行详情
func (s *sRoleMenu) Detail(ctx context.Context, in model.RoleMenuDetailInput) (data *serializer.RoleMenu, err error) {
	query := g.Map{
		"id":         in.Id,
		"is_deleted": consts.CREATED,
	}
	err = dao.RoleMenu.Ctx(ctx).Unscoped().Fields(model.RoleMenuFields).Where(query).Scan(&data)
	if err != nil {
		return data, err
	}
	if data == nil {
		return data, consts.ErrNotExit
	}
	return data, nil
}

// Delete 执行删除
func (s *sRoleMenu) Delete(ctx context.Context, in model.RoleMenuDeleteInput) (result sql.Result, err error) {
	query := g.Map{
		"id":         in.Id,
		"is_deleted": consts.CREATED,
	}
	result, err = dao.RoleMenu.Ctx(ctx).Unscoped().Where(query).Delete()
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *sRoleMenu) SafeDelete(ctx context.Context, r *model.OrmDeleteInput) (row int64, err error) {
	row, err = s.update(ctx, r.ToWhereMap(), r.ToMap())
	if err != nil {
		logger.Error(ctx, err)
		return 0, consts.ErrDel
	}
	return row, nil
}

func (s *sRoleMenu) Save(ctx context.Context, in *serializer.RoleMenu) (result sql.Result, err error) {
	return dao.RoleMenu.Ctx(ctx).Save(in)
}
