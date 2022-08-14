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

type ILog interface {
	Count(ctx context.Context) (data int, err error)
	List(ctx context.Context, in model.LogListInput) (items []*serializer.Log, err error)
	InsertAndGetId(ctx context.Context, data g.Map) (id int64, err error)
	Create(ctx context.Context, in *model.LogCreateInput) (id int64, err error)
	Update(ctx context.Context, in model.LogUpdateInput) (row int64, err error)
	Detail(ctx context.Context, in model.LogDetailInput) (data *serializer.Log, err error)
	Delete(ctx context.Context, in model.LogDeleteInput) (result sql.Result, err error)
	SafeDelete(ctx context.Context, r *model.OrmDeleteInput) (row int64, err error)
}

var localLog ILog

func Log() ILog {
	if localLog == nil {
		panic("implement not found for interface ILog, forgot register?")
	}
	return localLog
}

func RegisterLog(i ILog) {
	localLog = i
}