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

type IUser interface {
	FindOne(ctx context.Context, query *g.Map) (data *serializer.User, err error)
	Find(ctx context.Context, pk int64) (user *serializer.User, err error)
	LoginWeb(ctx context.Context, in model.UserLoginWebInput) (data *serializer.User, err error)
	Login(ctx context.Context, in model.UserLoginInput) (uid int64, err error)
	Register(ctx context.Context, in model.UserCreateInput) (uid int64, err error)
	Count(ctx context.Context) (data int, err error)
	Create(ctx context.Context, in model.UserCreateInput) (uid int64, err error)
	Save(ctx context.Context, passport, password, nickname string, roleId int, roleIds string) (result sql.Result, err error)
	Patch(ctx context.Context, in model.UserPatchInput) (uid int64, err error)
	Update(ctx context.Context, in model.UserUpdateInput) (uid int64, err error)
	List(ctx context.Context, in model.UserListInput) (items []*serializer.UserInfo, err error)
	Detail(ctx context.Context, in model.UserDetailInput) (data *serializer.User, err error)
	Delete(ctx context.Context, in model.UserDeleteInput) (result sql.Result, err error)
	SafeDelete(ctx context.Context, r *model.OrmDeleteInput) (row int64, err error)
	Logout(ctx context.Context) error
}

var localUser IUser

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
