// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package service

import (
	"admin/internal/model"
	"context"
)

type IStatus interface {
	GetServerInfo(ctx context.Context) (data *model.Server, err error)
	InitOs(ctx context.Context) model.Os
	InitCpu(ctx context.Context) (model.Cpu, error)
	InitRAM(ctx context.Context) (model.Rrm, error)
	InitDisk(ctx context.Context) (model.Disk, error)
}

var localStatus IStatus

func Status() IStatus {
	if localStatus == nil {
		panic("implement not found for interface IStatus, forgot register?")
	}
	return localStatus
}

func RegisterStatus(i IStatus) {
	localStatus = i
}