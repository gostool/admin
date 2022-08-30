package log

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
	"github.com/gogf/gf/v2/util/gconv"
)

type sLog struct {
}

var logger *glog.Logger

func init() {
	logger = g.Log(consts.LoggerDebug)
	instance := New()
	service.RegisterLog(instance)
}

func New() *sLog {
	return &sLog{}
}

func (s *sLog) Count(ctx context.Context) (data int, err error) {
	query := g.Map{
		"is_deleted": consts.CREATED,
	}
	return dao.OperationLog.Ctx(ctx).Count(query)
}

func (s *sLog) Search(ctx context.Context, in model.LogSearchInput) (items []*serializer.Log, err error) {
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
	if (in.Status > 0) && (in.Status < 600) {
		query["status"] = in.Status
	}
	return s.search(ctx, in.Page, in.PageSize, in.OrmSortType.Type, query)
}

func (s *sLog) search(ctx context.Context, page, limit int, orderType int, query g.Map) (dataList []*serializer.Log, err error) {
	path, ok := query["path"]
	if ok {
		// 针对path 字段，like 查询
		delete(query, "path")
	}
	m := dao.OperationLog.Ctx(ctx).Fields(model.LogFields).Page(page, limit).Where(query)
	if path != nil {
		m = m.WhereLike("path", gconv.String(path)+"%")
	}
	err = m.Order(consts.OrderFiledByType(orderType)).Scan(&dataList)
	if err != nil {
		return dataList, err
	}
	return dataList, err
}

func (s *sLog) List(ctx context.Context, in model.LogListInput) (items []*serializer.Log, err error) {
	query := g.Map{
		"is_deleted": consts.CREATED,
	}
	err = dao.OperationLog.Ctx(ctx).Fields(model.LogFields).Page(in.Page, in.PageSize).Where(query).Scan(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *sLog) InsertAndGetId(ctx context.Context, data g.Map) (id int64, err error) {
	id, err = dao.OperationLog.Ctx(ctx).InsertAndGetId(data)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (s *sLog) Create(ctx context.Context, in *model.LogCreateInput) (id int64, err error) {
	return s.InsertAndGetId(ctx, in.ToMap())
}

func (s *sLog) update(ctx context.Context, query, data g.Map) (row int64, err error) {
	logger.Debugf(ctx, "data:%v\n", data)
	result, err := dao.OperationLog.Ctx(ctx).Where(query).Update(data)
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
func (s *sLog) Update(ctx context.Context, in model.LogUpdateInput) (row int64, err error) {
	return s.update(ctx, in.ToWhereMap(), in.ToMap())
}

// Detail 执行详情
func (s *sLog) Detail(ctx context.Context, in model.LogDetailInput) (data *serializer.Log, err error) {
	query := g.Map{
		"id":         in.Id,
		"is_deleted": consts.CREATED,
	}
	err = dao.OperationLog.Ctx(ctx).Unscoped().Fields(model.LogFields).Where(query).Scan(&data)
	if err != nil {
		return data, err
	}
	if data == nil {
		return data, consts.ErrNotExit
	}
	return data, nil
}

// Delete 执行删除
func (s *sLog) Delete(ctx context.Context, in model.LogDeleteInput) (result sql.Result, err error) {
	query := g.Map{
		"id":         in.Id,
		"is_deleted": consts.CREATED,
	}
	result, err = dao.OperationLog.Ctx(ctx).Unscoped().Where(query).Delete()
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *sLog) SafeDelete(ctx context.Context, r *model.OrmDeleteInput) (row int64, err error) {
	row, err = s.update(ctx, r.ToWhereMap(), r.ToMap())
	if err != nil {
		logger.Error(ctx, err)
		return 0, consts.ErrDel
	}
	return row, nil
}
