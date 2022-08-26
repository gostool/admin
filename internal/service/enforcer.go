// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"admin/internal/model"
	"admin/internal/model/serializer"
	"context"

	"github.com/casbin/casbin/v2"
)

type IEnforcer interface {
	Enforcer(ctx context.Context) (enforcer *casbin.SyncedEnforcer)
	List(ctx context.Context, in model.EnforcerListInput) (items []*serializer.CasbinApi, err error)
	Update(ctx context.Context, in model.EnforcerUpdateInput) (err error)
	Clear(v int, p ...string) bool
}

var localEnforcer IEnforcer

func Enforcer() IEnforcer {
	if localEnforcer == nil {
		panic("implement not found for interface IEnforcer, forgot register?")
	}
	return localEnforcer
}

func RegisterEnforcer(i IEnforcer) {
	localEnforcer = i
}
