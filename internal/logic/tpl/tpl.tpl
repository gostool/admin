package tpl

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

type sTpl struct {
}

var logger *glog.Logger

func init() {
	instance := New()
	logger = g.Log(consts.LoggerDebug)
	service.RegisterTpl(instance)
}

func New() *sTpl {
	return &sTpl{}
}

func (s *sTpl) Count(ctx context.Context) (data int, err error) {
	query := g.Map{
		"is_deleted": consts.CREATED,
	}
	return dao.Tpl.Ctx(ctx).Count(query)
}

func (s *sTpl) List(ctx context.Context, in model.TplListInput) (items []*serializer.Tpl, err error) {
	query := g.Map{
		"is_deleted": consts.CREATED,
	}
	err = dao.Tpl.Ctx(ctx).Fields(model.TplFields).Page(in.Page, in.PageSize).Where(query).Scan(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *sTpl) InsertAndGetId(ctx context.Context, data g.Map) (id int64, err error) {
	id, err = dao.Tpl.Ctx(ctx).InsertAndGetId(data)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (s *sTpl) Create(ctx context.Context, in *model.TplCreateInput) (id int64, err error) {
	return s.InsertAndGetId(ctx, in.ToMap())
}

func (s *sTpl) update(ctx context.Context, query, data g.Map) (row int64, err error) {
	logger.Debugf(ctx, "data:%v\n", data)
	result, err := dao.Tpl.Ctx(ctx).Where(query).Update(data)
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
func (s *sTpl) Update(ctx context.Context, in model.TplUpdateInput) (row int64, err error) {
	return s.update(ctx, in.ToWhereMap(), in.ToMap())
}

// Detail 执行详情
func (s *sTpl) Detail(ctx context.Context, in model.TplDetailInput) (data *serializer.Tpl, err error) {
	query := g.Map{
		"id":         in.Id,
		"is_deleted": consts.CREATED,
	}
	err = dao.Tpl.Ctx(ctx).Unscoped().Fields(model.TplFields).Where(query).Scan(&data)
	if err != nil {
		return data, err
	}
	if data == nil {
		return data, consts.ErrNotExit
	}
	return data, nil
}

// Delete 执行删除
func (s *sTpl) Delete(ctx context.Context, in model.TplDeleteInput) (result sql.Result, err error) {
	query := g.Map{
		"id":         in.Id,
		"is_deleted": consts.CREATED,
	}
	result, err = dao.Tpl.Ctx(ctx).Unscoped().Where(query).Delete()
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *sTpl) SafeDelete(ctx context.Context, r *model.OrmDeleteInput) (row int64, err error) {
	row, err = s.update(ctx, r.ToWhereMap(), r.ToMap())
	if err != nil {
		logger.Error(ctx, err)
		return 0, consts.ErrDel
	}
	return row, nil
}

func (s *sTpl) Save(ctx context.Context, in *entity.Tpl) (result sql.Result, err error) {
	return dao.Tpl.Ctx(ctx).Save(in)
}

func (s *sTpl) SortField() string {
	return "sort asc"
}