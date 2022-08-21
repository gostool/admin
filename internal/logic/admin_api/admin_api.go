package admin_api

import (
	"admin/internal/consts"
	"admin/internal/dao"
	"admin/internal/model"
	"admin/internal/model/entity"
	"admin/internal/model/serializer"
	"admin/internal/service"
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAdminApi struct {
}

var logger *glog.Logger

func init() {
	instance := New()
	logger = g.Log(consts.LoggerDebug)
	service.RegisterAdminApi(instance)
}

func New() *sAdminApi {
	return &sAdminApi{}
}

func (s *sAdminApi) Count(ctx context.Context) (data int, err error) {
	query := g.Map{
		"is_deleted": consts.CREATED,
	}
	return dao.Api.Ctx(ctx).Count(query)
}

func (s *sAdminApi) Search(ctx context.Context, in model.AdminApiSearchInput) (items []*serializer.Api, err error) {
	query := g.Map{
		"is_deleted": consts.CREATED,
	}
	if in.Method != "" {
		query["method"] = in.Method
	}
	if gconv.String(in.Path) != "" {
		// 前缀查询
		query["path"] = in.Path
	}
	if in.Group != "" {
		query["group"] = in.Group
	}
	return s.search(ctx, in.Page, in.PageSize, in.OrderKey, in.Desc, query)
}

func (s *sAdminApi) search(ctx context.Context, page, limit int, orderKey string, desc bool, query g.Map) (dataList []*serializer.Api, err error) {
	path, ok := query["path"]
	if ok {
		// 针对path 字段，like 查询
		delete(query, "path")
	}
	m := dao.Api.Ctx(ctx).Fields(model.ApiFields).Page(page, limit).Where(query)
	if path != "" {
		m = m.WhereLike("path", gconv.String(path)+"%")
	}
	if desc {
		m = m.OrderDesc(orderKey)
	} else {
		m = m.OrderAsc(orderKey)
	}
	err = m.Scan(&dataList)
	if err != nil {
		return dataList, err
	}
	return dataList, err
}

func (s *sAdminApi) List(ctx context.Context, in model.AdminApiListInput) (items []*serializer.Api, err error) {
	query := g.Map{
		"is_deleted": consts.CREATED,
	}
	err = dao.Api.Ctx(ctx).Fields(model.ApiFields).Page(in.Page, in.PageSize).Where(query).Scan(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *sAdminApi) InsertAndGetId(ctx context.Context, data g.Map) (id int64, err error) {
	id, err = dao.Api.Ctx(ctx).InsertAndGetId(data)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (s *sAdminApi) Create(ctx context.Context, in *model.AdminApiCreateInput) (id int64, err error) {
	return s.InsertAndGetId(ctx, in.ToMap())
}

func (s *sAdminApi) update(ctx context.Context, query, data g.Map) (row int64, err error) {
	logger.Debugf(ctx, "data:%v\n", data)
	result, err := dao.Api.Ctx(ctx).Where(query).Update(data)
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
func (s *sAdminApi) Update(ctx context.Context, in model.AdminApiUpdateInput) (row int64, err error) {
	return s.update(ctx, in.ToWhereMap(), in.ToMap())
}

// Detail 执行详情
func (s *sAdminApi) Detail(ctx context.Context, in model.AdminApiDetailInput) (data *serializer.Api, err error) {
	query := g.Map{
		"id":         in.Id,
		"is_deleted": consts.CREATED,
	}
	err = dao.Api.Ctx(ctx).Unscoped().Fields(model.ApiFields).Where(query).Scan(&data)
	if err != nil {
		return data, err
	}
	if data == nil {
		return data, consts.ErrNotExit
	}
	return data, nil
}

// Delete 执行删除
func (s *sAdminApi) Delete(ctx context.Context, in model.AdminApiDeleteInput) (result sql.Result, err error) {
	query := g.Map{
		"id":         in.Id,
		"is_deleted": consts.CREATED,
	}
	result, err = dao.Api.Ctx(ctx).Unscoped().Where(query).Delete()
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *sAdminApi) SafeDelete(ctx context.Context, r *model.OrmDeleteInput) (row int64, err error) {
	row, err = s.update(ctx, r.ToWhereMap(), r.ToMap())
	if err != nil {
		logger.Error(ctx, err)
		return 0, consts.ErrDel
	}
	return row, nil
}

func (s *sAdminApi) Save(ctx context.Context, in *entity.Api) (result sql.Result, err error) {
	return dao.Api.Ctx(ctx).Save(in)
}
