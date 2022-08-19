package controller

import (
	v1 "admin/api/v1"
	"admin/internal/model"
	"admin/internal/service"
	"context"
)

var (
	AdminApi = cAdminApi{}
)

type cAdminApi struct{}

func (c *cAdminApi) List(ctx context.Context, req *v1.AdminApiListReq) (res *v1.AdminApiListRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	in := model.AdminApiListInput{
		Page:     req.PageReq.Page,
		PageSize: req.PageReq.PageSize,
	}
	items, err := service.AdminApi().List(ctx, in)
	if err != nil {
		return res, err
	}
	cnt, err := service.AdminApi().Count(ctx)
	if err != nil {
		return res, err
	}
	res = &v1.AdminApiListRes{
		Count: cnt,
		Items: items,
	}
	return res, nil
}

func (c *cAdminApi) Detail(ctx context.Context, req *v1.AdminApiDetailReq) (res *v1.AdminApiDetailRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	data, err := service.AdminApi().Detail(ctx, model.AdminApiDetailInput{
		Id: req.Id,
	})
	if err != nil {
		return res, err
	}

	res = &v1.AdminApiDetailRes{
		Api: data,
	}
	return res, nil
}

func (c *cAdminApi) Update(ctx context.Context, req *v1.AdminApiUpdateReq) (res *v1.OrmUpdateRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	in := model.AdminApiUpdateInput{
		Id: req.Id,
		AdminApiAttr: model.AdminApiAttr{
			Path:   req.Path,
			Method: req.Method,
			Group:  req.Group,
			Detail: req.Detail,
		},
	}
	_, err = service.AdminApi().Update(ctx, in)
	if err != nil {
		return res, err
	}
	res = &v1.OrmUpdateRes{}
	return res, nil
}

func (c *cAdminApi) Delete(ctx context.Context, req *v1.AdminApiDeleteReq) (res *v1.OrmDeleteRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	_, err = service.AdminApi().SafeDelete(ctx, &model.OrmDeleteInput{
		Id: req.Id,
	})
	if err != nil {
		return res, err
	}
	res = &v1.OrmDeleteRes{}
	return res, nil
}

func (c *cAdminApi) Create(ctx context.Context, req *v1.AdminApiCreateReq) (res *v1.OrmIdRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	in := model.AdminApiCreateInput{
		AdminApiAttr: model.AdminApiAttr{
			Path:   req.Path,
			Method: req.Method,
			Group:  req.Group,
			Detail: req.Detail,
		},
	}
	id, err := service.AdminApi().Create(ctx, &in)
	if err != nil {
		return res, err
	}
	res = &v1.OrmIdRes{}
	res.Id = int(id)
	return res, nil
}
