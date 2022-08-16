// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"admin/internal/model"
	"admin/internal/model/serializer"
	"context"
	"database/sql"

	"github.com/gogf/gf/v2/frame/g"
)

type IMenu interface {
	Count(ctx context.Context) (data int, err error)
	List(ctx context.Context, in model.MenuListInput) (items []*serializer.Menu, err error)
	InsertAndGetId(ctx context.Context, data g.Map) (id int64, err error)
	Create(ctx context.Context, in *model.MenuCreateInput) (id int64, err error)
	Save(ctx context.Context, in *serializer.Menu) (result sql.Result, err error)
	Update(ctx context.Context, in model.MenuUpdateInput) (row int64, err error)
	Detail(ctx context.Context, in model.MenuDetailInput) (data *serializer.Menu, err error)
	Delete(ctx context.Context, in model.MenuDeleteInput) (result sql.Result, err error)
	SafeDelete(ctx context.Context, r *model.OrmDeleteInput) (row int64, err error)
}

var localMenu IMenu

func Menu() IMenu {
	if localMenu == nil {
		panic("implement not found for interface IMenu, forgot register?")
	}
	return localMenu
}

func RegisterMenu(i IMenu) {
	localMenu = i
}
