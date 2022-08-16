package controller

import (
	v1 "admin/api/v1"
	"admin/internal/model"
	"admin/internal/service"
	"context"
)

var (
	RoleMenu = cRoleMenu{}
)

type cRoleMenu struct{}

func (c *cRoleMenu) List(ctx context.Context, req *v1.RoleMenuListReq) (res *v1.RoleMenuListRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	items, err := service.RoleMenu().List(ctx, model.RoleMenuListInput{
		Page:     req.PageReq.Page,
		PageSize: req.PageReq.PageSize,
	})
	if err != nil {
		return res, err
	}
	cnt, err := service.RoleMenu().Count(ctx)
	if err != nil {
		return res, err
	}
	res = &v1.RoleMenuListRes{
		Count: cnt,
		Items: items,
	}
	return res, nil
}

func (c *cRoleMenu) Detail(ctx context.Context, req *v1.RoleMenuDetailReq) (res *v1.RoleMenuDetailRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	data, err := service.RoleMenu().Detail(ctx, model.RoleMenuDetailInput{
		Id: req.Id,
	})
	if err != nil {
		return res, err
	}

	res = &v1.RoleMenuDetailRes{
		RoleMenu: data,
	}
	return res, nil
}

func (c *cRoleMenu) Update(ctx context.Context, req *v1.RoleMenuUpdateReq) (res *v1.OrmUpdateRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	in := model.RoleMenuUpdateInput{
		Id: req.Id,
		RoleMenuAttr: model.RoleMenuAttr{
			RoleId: req.RoleId,
			MenuId: req.MenuId,
		},
	}
	_, err = service.RoleMenu().Update(ctx, in)
	if err != nil {
		return res, err
	}
	res = &v1.OrmUpdateRes{}
	return res, nil
}

func (c *cRoleMenu) Delete(ctx context.Context, req *v1.RoleMenuDeleteReq) (res *v1.OrmDeleteRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	_, err = service.RoleMenu().SafeDelete(ctx, &model.OrmDeleteInput{
		Id: req.Id,
	})
	if err != nil {
		return res, err
	}
	res = &v1.OrmDeleteRes{}
	return res, nil
}

func (c *cRoleMenu) Create(ctx context.Context, req *v1.RoleMenuCreateReq) (res *v1.RoleMenuCreateRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	in := model.RoleMenuCreateInput{
		RoleMenuAttr: model.RoleMenuAttr{
			RoleId: req.RoleId,
			MenuId: req.MenuId,
		},
	}
	id, err := service.RoleMenu().Create(ctx, &in)
	if err != nil {
		return res, err
	}
	res = &v1.RoleMenuCreateRes{}
	res.Id = int(id)
	return res, nil
}
