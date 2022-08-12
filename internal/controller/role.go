package controller

import (
	v1 "admin/api/v1"
	"admin/internal/model"
	"admin/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Role = cRole{}
)

type cRole struct{}

func (c *cRole) List(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	g.Log().Debugf(ctx, `receive say: %+v`, req)
	data, err := service.Role().List(ctx, model.RoleListInput{
		Page:     req.PageReq.Page,
		PageSize: req.PageReq.PageSize,
	})
	if err != nil {
		return res, err
	}
	res = &v1.RoleListRes{
		Data: data,
	}
	return res, nil
}

func (c *cRole) Detail(ctx context.Context, req *v1.RoleDetailReq) (res *v1.RoleDetailRes, err error) {
	g.Log().Debugf(ctx, `receive say: %+v`, req)
	data, err := service.Role().Detail(ctx, model.RoleDetailInput{
		Id: req.Id,
	})
	if err != nil {
		return res, err
	}
	res = &v1.RoleDetailRes{
		Data: data,
	}
	return res, nil
}

func (c *cRole) Update(ctx context.Context, req *v1.RoleUpdateReq) (res *v1.RoleUpdateRes, err error) {
	g.Log().Debugf(ctx, `receive say: %+v`, req)
	id, err := service.Role().Update(ctx, model.RoleUpdateInput{
		Id: req.Id,
		RoleAttr: model.RoleAttr{
			Name:   req.Name,
			Router: req.Router,
			Pid:    req.Pid,
		},
	})
	if err != nil {
		return res, err
	}
	res = &v1.RoleUpdateRes{
		Id: string(id),
	}
	return res, nil
}

func (c *cRole) Delete(ctx context.Context, req *v1.RoleDeleteReq) (res *v1.RoleDeleteRes, err error) {
	g.Log().Debugf(ctx, `receive say: %+v`, req)
	data, err := service.Role().Delete(ctx, model.RoleDeleteInput{
		Id: req.Id,
	})
	if err != nil {
		return res, err
	}
	res = &v1.RoleDeleteRes{
		Data: data,
	}
	return res, nil
}
