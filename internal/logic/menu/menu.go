package menu

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

type sMenu struct {
}

var logger *glog.Logger

func init() {
	logger = g.Log(consts.LoggerDebug)
	instance := New()
	service.RegisterMenu(instance)
}

func New() *sMenu {
	return &sMenu{}
}

func (s *sMenu) Count(ctx context.Context) (data int, err error) {
	query := g.Map{
		"is_deleted": consts.CREATED,
	}
	return dao.Menu.Ctx(ctx).Count(query)
}

func (s *sMenu) List(ctx context.Context, in model.MenuListInput) (items []*serializer.Menu, err error) {
	query := g.Map{
		"is_deleted": consts.CREATED,
	}
	err = dao.Menu.Ctx(ctx).Fields(model.MenuFields).Page(in.Page, in.PageSize).Where(query).Scan(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *sMenu) InsertAndGetId(ctx context.Context, data g.Map) (id int64, err error) {
	id, err = dao.Menu.Ctx(ctx).InsertAndGetId(data)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (s *sMenu) Create(ctx context.Context, in *model.MenuCreateInput) (id int64, err error) {
	return s.InsertAndGetId(ctx, in.ToMap())
}

func (s *sMenu) update(ctx context.Context, query, data g.Map) (row int64, err error) {
	logger.Debugf(ctx, "data:%v\n", data)
	result, err := dao.Menu.Ctx(ctx).Where(query).Update(data)
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
func (s *sMenu) Update(ctx context.Context, in model.MenuUpdateInput) (row int64, err error) {
	return s.update(ctx, in.ToWhereMap(), in.ToMap())
}

// Detail 执行详情
func (s *sMenu) Detail(ctx context.Context, in model.MenuDetailInput) (data *serializer.Menu, err error) {
	query := g.Map{
		"id":         in.Id,
		"is_deleted": consts.CREATED,
	}
	err = dao.Menu.Ctx(ctx).Unscoped().Fields(model.MenuFields).Where(query).Scan(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// Delete 执行删除
func (s *sMenu) Delete(ctx context.Context, in model.MenuDeleteInput) (result sql.Result, err error) {
	query := g.Map{
		"id":         in.Id,
		"is_deleted": consts.CREATED,
	}
	result, err = dao.Menu.Ctx(ctx).Unscoped().Where(query).Delete()
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *sMenu) SafeDelete(ctx context.Context, r *model.OrmDeleteInput) (row int64, err error) {
	row, err = s.update(ctx, r.ToWhereMap(), r.ToMap())
	if err != nil {
		logger.Error(ctx, err)
		return 0, consts.ErrDel
	}
	return row, nil
}
