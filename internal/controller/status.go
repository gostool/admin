package controller

import (
	v1 "admin/api/v1"
	"admin/internal/service"
	"context"
)

var (
	Status = cStatus{}
)

type cStatus struct{}

func (c *cStatus) List(ctx context.Context, req *v1.StatusReq) (res *v1.StatusRes, err error) {
	data, err := service.Status().GetServerInfo()
	if err != nil {
		return res, err
	}
	res = &v1.StatusRes{
		Server: data,
	}
	return res, nil
}
