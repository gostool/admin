package role_menu

import (
	"admin/internal/consts"
	"admin/internal/dao"
	"admin/internal/model"
	"admin/internal/model/entity"
	"admin/internal/model/serializer"
	"admin/internal/service"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"sync"
)

type sRoleMenu struct {
}

var logger *glog.Logger

type txBulkFunc func(ctx context.Context, tx *gdb.TX, roleId int, idSet *gset.IntSet) (err error)

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

func (s *sRoleMenu) SortField() string {
	return "sort asc"
}

func (s *sRoleMenu) RoleMenuIdList(ctx context.Context, roleId int) (idList []int, err error) {
	query := g.Map{
		"role_id":    roleId,
		"is_deleted": consts.CREATED,
	}
	objList := ([]serializer.RoleMenu)(nil)
	err = dao.RoleMenu.Ctx(ctx).Fields(model.RoleMenuFields).Where(query).Scan(&objList)
	if err != nil {
		return idList, err
	}
	idList = make([]int, 0, len(objList))
	for _, obj := range objList {
		idList = append(idList, obj.MenuId)
	}
	return idList, nil
}

func (s *sRoleMenu) GetRoleMenuList(ctx context.Context, roleId int) (objList []*serializer.MenuDetail, err error) {
	// 1.获取当前角色的菜单列表
	menuIdList, err := s.RoleMenuIdList(ctx, roleId)
	if err != nil {
		return objList, err
	}
	// 2.获取所有菜单
	query := g.Map{
		"id":         menuIdList,
		"is_deleted": consts.CREATED,
	}
	objList = ([]*serializer.MenuDetail)(nil)
	err = dao.Menu.Ctx(ctx).Fields(model.MenuFields).Order(s.SortField()).Where(query).Scan(&objList)
	if err != nil {
		return objList, err
	}
	return objList, nil
}

// GetTreeMap
// 菜单树
// 根据父节点聚合
func (s *sRoleMenu) GetTreeMap(ctx context.Context, roleId int) (treeMap map[int][]*serializer.MenuDetail, err error) {
	objList, err := s.GetRoleMenuList(ctx, roleId)
	if err != nil {
		return nil, err
	}
	// 1.父节点聚合
	treeMap = make(map[int][]*serializer.MenuDetail)
	for _, v := range objList {
		treeMap[v.ParentId] = append(treeMap[v.ParentId], v)
	}
	return treeMap, err
}

func (s *sRoleMenu) GetTreeByRoleId(ctx context.Context, roleId int) (items []*serializer.MenuDetail, err error) {
	menuTree, err := s.GetTreeMap(ctx, roleId)
	root := menuTree[0]
	// 遍历一级菜单
	items = root
	for i := 0; i < len(root); i++ {
		menu := items[i]
		if menu == nil {
			continue
		}

		childList, ok := menuTree[menu.Id]
		if !ok {
			// 当前的一级菜单没有child
			continue
		}
		if menu.Children == nil {
			n := len(childList)
			menu.Children = make([]*serializer.MenuDetail, n, n)
		}
		copy(menu.Children, childList)
		logger.Debugf(ctx, "menu:%v Children:%v\n", menu.Title, len(menu.Children))
	}
	return items, err
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

func (s *sRoleMenu) Save(ctx context.Context, in *entity.RoleMenu) (result sql.Result, err error) {
	return dao.RoleMenu.Ctx(ctx).Save(in)
}

func txBulkDelete(ctx context.Context, tx *gdb.TX, roleId int, idSet *gset.IntSet) (err error) {
	if idSet == nil {
		return nil
	}
	if idSet.Size() <= 0 {
		return nil
	}
	query := g.Map{
		"role_id": roleId,
		"menu_id": idSet.Slice(),
	}
	_, err = dao.RoleMenu.Ctx(ctx).Unscoped().TX(tx).Where(query).Delete()
	return err
}

func txBulkInsert(ctx context.Context, tx *gdb.TX, roleId int, idSet *gset.IntSet) (err error) {
	if idSet == nil {
		return nil
	}
	if idSet.Size() <= 0 {
		return nil
	}
	data := make([]g.Map, 0, idSet.Size())
	idSet.Iterator(func(v int) bool {
		item := g.Map{}
		item["menu_id"] = v
		item["role_id"] = roleId
		item["is_deleted"] = consts.CREATED
		data = append(data, item)
		return true
	})
	_, err = dao.RoleMenu.Ctx(ctx).TX(tx).Data(data).Insert()
	return err
}

func (s *sRoleMenu) TxBulkCreateRoleMenu(ctx context.Context, roleId int, deleteIdSet, insertIdSet *gset.IntSet) (err error) {
	tx, err := g.DB().Begin(ctx)
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	err = txBulkDelete(ctx, tx, roleId, deleteIdSet)
	if err != nil {
		return err
	}
	err = txBulkInsert(ctx, tx, roleId, insertIdSet)
	if err != nil {
		return err
	}
	return nil
}

// makeMenuRoleDBBulkResponse
// 获取所有goroutine的结果.
// 上传数据
func makeMenuRoleDBBulkResponse(ch <-chan g.Map) (err error) {
	for response := range ch {
		if gconv.Int(response["code"]) != consts.PoolWorkOk {
			msg := gconv.String(response["msg"])
			return errors.New(msg)
		}
	}
	return nil
}

func JobRoleMenuBulk(ctx context.Context, f txBulkFunc, tx *gdb.TX, roleId int, idSet *gset.IntSet, ch chan<- g.Map) {
	data := g.Map{
		"code": consts.PoolWorkOk,
		"msg":  "",
		"type": consts.JobTypeRoleMenu,
		"data": nil,
	}
	defer func() {
		ch <- data
		if err := recover(); err != nil {
			logger.Error(ctx, err)
		}
	}()
	err := f(ctx, tx, roleId, idSet)
	if err != nil {
		data["code"] = consts.PoolWorkErr
		data["msg"] = err.Error()
	}
}

func (s *sRoleMenu) AsyncTxBulkCreateRoleMenu(ctx context.Context, roleId int, deleteIdSet, insertIdSet *gset.IntSet) (err error) {
	tx, err := g.DB().Begin(ctx)
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan g.Map, 2)
	service.Pool().Add(ctx, func(ctx context.Context) {
		defer wg.Done()
		JobRoleMenuBulk(ctx, txBulkDelete, tx, roleId, deleteIdSet, ch)
	})
	service.Pool().Add(ctx, func(ctx context.Context) {
		defer wg.Done()
		JobRoleMenuBulk(ctx, txBulkInsert, tx, roleId, insertIdSet, ch)
	})
	service.Pool().Show(ctx)
	wg.Wait()
	service.Pool().Show(ctx)
	close(ch)
	err = makeMenuRoleDBBulkResponse(ch)
	if err != nil {
		return err
	}
	return nil
}

func (s *sRoleMenu) GetMenuIdSet(ctx context.Context, roleId int) (idSet *gset.IntSet, err error) {
	objIdList, err := s.RoleMenuIdList(ctx, roleId)
	if err != nil {
		return idSet, err
	}
	idSet = gset.NewIntSetFrom(objIdList)
	return idSet, nil
}

func (s *sRoleMenu) BulkCreateRoleMenu(ctx context.Context, r *model.PermMenuReq) (err error) {
	oldIdSet, err := s.GetMenuIdSet(ctx, r.RoleId)
	if err != nil {
		return err
	}
	newIdSet := gset.NewIntSet()
	for _, v := range r.MenuIdList {
		newIdSet.Add(v)
	}
	logger.Infof(ctx, "RoleId:%v\n", r.RoleId)
	logger.Infof(ctx, "newIdSet:%v\n", newIdSet)
	logger.Infof(ctx, "oldIdSet:%v\n", oldIdSet)
	deleteIdSet := oldIdSet.Diff(newIdSet)
	insertIdSet := newIdSet.Diff(oldIdSet)
	logger.Infof(ctx, "deleteIdSet:%v\n", deleteIdSet)
	logger.Infof(ctx, "insertIdSet:%v\n", insertIdSet)
	if err != nil {
		return err
	}
	//return s.TxBulkCreateRoleMenu(ctx, r.RoleId, deleteIdSet, insertIdSet)
	return s.AsyncTxBulkCreateRoleMenu(ctx, r.RoleId, deleteIdSet, insertIdSet)
}
