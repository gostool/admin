package controller

import (
	v1 "admin/api/v1"
	"admin/internal/model"
	"admin/internal/service"
	"context"
)

var (
	Role = cRole{}
)

type cRole struct{}

func (c *cRole) List(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	items, err := service.Role().List(ctx, model.RoleListInput{
		Page:     req.PageReq.Page,
		PageSize: req.PageReq.PageSize,
	})
	if err != nil {
		return res, err
	}
	cnt, err := service.Role().Count(ctx)
	if err != nil {
		return res, err
	}
	res = &v1.RoleListRes{
		Count: cnt,
		Items: items,
	}
	return res, nil
}

func (c *cRole) Detail(ctx context.Context, req *v1.RoleDetailReq) (res *v1.RoleDetailRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	data, err := service.Role().Detail(ctx, model.RoleDetailInput{
		Id: req.Id,
	})
	if err != nil {
		return res, err
	}

	res = &v1.RoleDetailRes{
		Role: data,
	}
	return res, nil
}

func (c *cRole) Update(ctx context.Context, req *v1.RoleUpdateReq) (res *v1.OrmUpdateRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	in := model.RoleUpdateInput{
		Id: req.Id,
		RoleAttr: model.RoleAttr{
			Name:   req.Name,
			Router: req.Router,
			Pid:    req.Pid,
		},
	}
	_, err = service.Role().Update(ctx, in)
	if err != nil {
		return res, err
	}
	res = &v1.OrmUpdateRes{}
	return res, nil
}

func (c *cRole) Delete(ctx context.Context, req *v1.RoleDeleteReq) (res *v1.OrmDeleteRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	_, err = service.Role().SafeDelete(ctx, &model.OrmDeleteInput{
		Id: req.Id,
	})
	if err != nil {
		return res, err
	}
	res = &v1.OrmDeleteRes{}
	return res, nil
}
