package role

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
)

type sRole struct {
}

var logger *glog.Logger

func init() {
	instance := New()
	logger = g.Log(consts.LoggerDebug)
	service.RegisterRole(instance)
}

func New() *sRole {
	return &sRole{}
}

func (s *sRole) Count(ctx context.Context) (data int, err error) {
	query := g.Map{
		"is_deleted": consts.CREATED,
	}
	return dao.Role.Ctx(ctx).Count(query)
}

func (s *sRole) List(ctx context.Context, in model.RoleListInput) (items []*serializer.Role, err error) {
	query := g.Map{
		"is_deleted": consts.CREATED,
	}
	err = dao.Role.Ctx(ctx).Fields(model.RoleFields).Page(in.Page, in.PageSize).Where(query).Scan(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *sRole) InsertAndGetId(ctx context.Context, data g.Map) (id int64, err error) {
	id, err = dao.Role.Ctx(ctx).InsertAndGetId(data)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (s *sRole) Create(ctx context.Context, in *model.RoleCreateInput) (id int64, err error) {
	return s.InsertAndGetId(ctx, in.ToMap())
}

func (s *sRole) update(ctx context.Context, query, data g.Map) (row int64, err error) {
	logger.Debugf(ctx, "data:%v\n", data)
	result, err := dao.Role.Ctx(ctx).Where(query).Update(data)
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
func (s *sRole) Update(ctx context.Context, in model.RoleUpdateInput) (row int64, err error) {
	return s.update(ctx, in.ToWhereMap(), in.ToMap())
}

// Detail 执行详情
func (s *sRole) Detail(ctx context.Context, in model.RoleDetailInput) (data *serializer.Role, err error) {
	query := g.Map{
		"id":         in.Id,
		"is_deleted": consts.CREATED,
	}
	err = dao.Role.Ctx(ctx).Unscoped().Fields(model.RoleFields).Where(query).Scan(&data)
	if err != nil {
		return data, err
	}
	if data == nil {
		return data, consts.ErrNotExit
	}
	return data, nil
}

// Delete 执行删除
func (s *sRole) Delete(ctx context.Context, in model.RoleDeleteInput) (result sql.Result, err error) {
	query := g.Map{
		"id":         in.Id,
		"is_deleted": consts.CREATED,
	}
	result, err = dao.Role.Ctx(ctx).Unscoped().Where(query).Delete()
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *sRole) SafeDelete(ctx context.Context, r *model.OrmDeleteInput) (row int64, err error) {
	row, err = s.update(ctx, r.ToWhereMap(), r.ToMap())
	if err != nil {
		logger.Error(ctx, err)
		return 0, consts.ErrDel
	}
	return row, nil
}

func (s *sRole) Save(ctx context.Context, in *entity.Role) (result sql.Result, err error) {
	return dao.Role.Ctx(ctx).Save(in)
}

func (s *sRole) GetTree(ctx context.Context, in model.RoleListInput) (items []*serializer.RoleDetail, err error) {
	query := g.Map{
		"is_deleted": consts.CREATED,
	}
	var objList []*serializer.RoleDetail
	err = dao.Role.Ctx(ctx).Fields(model.RoleFields).Page(in.Page, in.PageSize).Where(query).Scan(&objList)
	if err != nil {
		return nil, err
	}
	treeMap := make(map[int][]*serializer.RoleDetail)
	for _, v := range objList {
		treeMap[v.Pid] = append(treeMap[v.Pid], v)
	}
	items = s.Bfs(treeMap)
	return items, nil
}

func (s *sRole) Bfs(treeMap map[int][]*serializer.RoleDetail) (items []*serializer.RoleDetail) {
	q := treeMap[0]
	items = q
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			node := q[i]
			if node == nil {
				continue
			}
			childList, ok := treeMap[node.Id]
			if !ok {
				// 当前菜单没有child
				continue
			}
			if node.Children == nil {
				n := len(childList)
				node.Children = make([]*serializer.RoleDetail, n, n)
			}
			copy(node.Children, childList)
			q = append(q, node.Children...)
		}
		q = q[size:]
	}
	return
}
