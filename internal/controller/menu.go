package controller

import (
	v1 "admin/api/v1"
	"admin/internal/model"
	"admin/internal/service"
	"context"
)

var (
	Menu = cMenu{}
)

type cMenu struct{}

func (c *cMenu) List(ctx context.Context, req *v1.MenuListReq) (res *v1.MenuListRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	items, err := service.Menu().List(ctx, model.MenuListInput{
		Page:     req.PageReq.Page,
		PageSize: req.PageReq.PageSize,
	})
	if err != nil {
		return res, err
	}
	cnt, err := service.Menu().Count(ctx)
	if err != nil {
		return res, err
	}
	res = &v1.MenuListRes{
		Count: cnt,
		Items: items,
	}
	return res, nil
}

func (c *cMenu) Detail(ctx context.Context, req *v1.MenuDetailReq) (res *v1.MenuDetailRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	data, err := service.Menu().Detail(ctx, model.MenuDetailInput{
		Id: req.Id,
	})
	if err != nil {
		return res, err
	}

	res = &v1.MenuDetailRes{
		Menu: data,
	}
	return res, nil
}

func (c *cMenu) Update(ctx context.Context, req *v1.MenuUpdateReq) (res *v1.OrmUpdateRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	in := model.MenuUpdateInput{
		Id: req.Id,
		MenuAttr: model.MenuAttr{
			Title:       req.Title,
			Icon:        req.Icon,
			KeepAlive:   req.KeepAlive,
			DefaultMenu: req.DefaultMenu,
			Hidden:      req.Hidden,
			Sort:        req.Sort,
			ParentId:    req.ParentId,
			Name:        req.Name,
			Path:        req.Path,
			Component:   req.Component,
		},
	}
	_, err = service.Menu().Update(ctx, in)
	if err != nil {
		return res, err
	}
	res = &v1.OrmUpdateRes{}
	return res, nil
}

func (c *cMenu) Delete(ctx context.Context, req *v1.MenuDeleteReq) (res *v1.OrmDeleteRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	_, err = service.Menu().SafeDelete(ctx, &model.OrmDeleteInput{
		Id: req.Id,
	})
	if err != nil {
		return res, err
	}
	res = &v1.OrmDeleteRes{}
	return res, nil
}

func (c *cMenu) Create(ctx context.Context, req *v1.MenuCreateReq) (res *v1.MenuCreateRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	in := model.MenuCreateInput{
		MenuAttr: model.MenuAttr{
			Title:       req.Title,
			Icon:        req.Icon,
			KeepAlive:   req.KeepAlive,
			DefaultMenu: req.DefaultMenu,
			Hidden:      req.Hidden,
			Sort:        req.Sort,
			ParentId:    req.ParentId,
			Name:        req.Name,
			Path:        req.Path,
			Component:   req.Component,
		},
	}
	id, err := service.Menu().Create(ctx, &in)
	if err != nil {
		return res, err
	}
	res = &v1.MenuCreateRes{}
	res.Id = int(id)
	return res, nil
}
