// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"admin/internal/model"
	"admin/internal/model/entity"
	"admin/internal/model/serializer"
	"context"
	"database/sql"

	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/frame/g"
)

type IRoleMenu interface {
	Count(ctx context.Context) (data int, err error)
	List(ctx context.Context, in model.RoleMenuListInput) (items []*serializer.RoleMenu, err error)
	SortField() string
	RoleMenuIdList(ctx context.Context, roleId int) (idList []int, err error)
	GetRoleMenuList(ctx context.Context, roleId int) (objList []*serializer.MenuDetail, err error)
	GetTreeMap(ctx context.Context, roleId int) (treeMap map[int][]*serializer.MenuDetail, err error)
	GetTreeByRoleId(ctx context.Context, roleId int) (items []*serializer.MenuDetail, err error)
	InsertAndGetId(ctx context.Context, data g.Map) (id int64, err error)
	Create(ctx context.Context, in *model.RoleMenuCreateInput) (id int64, err error)
	Update(ctx context.Context, in model.RoleMenuUpdateInput) (row int64, err error)
	Detail(ctx context.Context, in model.RoleMenuDetailInput) (data *serializer.RoleMenu, err error)
	Delete(ctx context.Context, in model.RoleMenuDeleteInput) (result sql.Result, err error)
	SafeDelete(ctx context.Context, r *model.OrmDeleteInput) (row int64, err error)
	Save(ctx context.Context, in *entity.RoleMenu) (result sql.Result, err error)
	TxBulkCreateRoleMenu(ctx context.Context, roleId int, deleteIdSet, insertIdSet *gset.IntSet) (err error)
	AsyncTxBulkCreateRoleMenu(ctx context.Context, roleId int, deleteIdSet, insertIdSet *gset.IntSet) (err error)
	GetMenuIdSet(ctx context.Context, roleId int) (idSet *gset.IntSet, err error)
	BulkCreateRoleMenu(ctx context.Context, r *model.PermMenuReq) (err error)
}

var localRoleMenu IRoleMenu

func RoleMenu() IRoleMenu {
	if localRoleMenu == nil {
		panic("implement not found for interface IRoleMenu, forgot register?")
	}
	return localRoleMenu
}

func RegisterRoleMenu(i IRoleMenu) {
	localRoleMenu = i
}
