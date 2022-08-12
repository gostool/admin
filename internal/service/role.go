// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"admin/internal/model"
	"admin/internal/model/entity"
	"context"
	"database/sql"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type IRole interface {
	Count(ctx context.Context) (data int, err error)
	List(ctx context.Context, in model.RoleListInput) (data gdb.Result, err error)
	InsertAndGetId(ctx context.Context, data g.Map) (id int64, err error)
	Create(ctx context.Context, in *model.RoleCreateInput) (id int64, err error)
	Update(ctx context.Context, in model.RoleUpdateInput) (id int64, err error)
	Detail(ctx context.Context, in model.RoleDetailInput) (data *entity.Role, err error)
	Delete(ctx context.Context, in model.RoleDeleteInput) (result sql.Result, err error)
}

var localRole IRole

func Role() IRole {
	if localRole == nil {
		panic("implement not found for interface IRole, forgot register?")
	}
	return localRole
}

func RegisterRole(i IRole) {
	localRole = i
}
