package controller

import (
	v1 "admin/api/v1"
	"admin/internal/model"
	"admin/internal/service"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	AdminCasbin = cAdminCasbin{}
)

type cAdminCasbin struct{}

func (c *cAdminCasbin) List(ctx context.Context, req *v1.AdminCasbinListReq) (res *v1.AdminCasbinListRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	in := model.EnforcerListInput{
		RoleId: req.RoleId,
	}
	items, err := service.Enforcer().List(ctx, in)
	if err != nil {
		return res, err
	}
	res = &v1.AdminCasbinListRes{
		Count: -1,
		Items: items,
	}
	return res, nil
}

func (c *cAdminCasbin) Update(ctx context.Context, req *v1.AdminCasbinUpdateReq) (res *v1.AdminCasbinUpdateRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	in := model.EnforcerUpdateInput{
		RoleId: req.RoleId,
	}
	infoList := make([]model.AdminCasbinAttr, len(req.ApiInfoList))
	err = gconv.Struct(req.ApiInfoList, &infoList)
	if err != nil {
		return nil, err
	}
	err = service.Enforcer().Update(ctx, in)
	if err != nil {
		return res, err
	}

	res = &v1.AdminCasbinUpdateRes{}
	return res, nil
}
