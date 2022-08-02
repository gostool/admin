// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"admin/internal/model"
	"context"
)

type IUser interface {
	Login(ctx context.Context, in model.UserLoginInput) (err error)
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