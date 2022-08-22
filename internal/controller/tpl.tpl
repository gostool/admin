package controller

import (
	v1 "admin/api/v1"
	"admin/internal/model"
	"admin/internal/service"
	"context"
)

var (
	Tpl = cTpl{}
)

type cTpl struct{}

func (c *cTpl) List(ctx context.Context, req *v1.TplListReq) (res *v1.TplListRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	attrReqSearch := model.OpTplAttrSearch{
		Status: req.Status,
		Method: req.Method,
		Path:   req.Path,
	}
	ormSortType := model.OrmSortType{
		Type: req.Type,
	}
	in := model.TplSearchInput{
		Page:     req.PageReq.Page,
		PageSize: req.PageReq.PageSize,
	}
    in.OpTplAttrSearch = attrReqSearch
	in.OrmSortType = ormSortType
	items, err := service.Tpl().Search(ctx, in)
	if err != nil {
		return res, err
	}
	cnt, err := service.Tpl().Count(ctx)
	if err != nil {
		return res, err
	}
	res = &v1.TplListRes{
		Count: cnt,
		Items: items,
	}
	return res, nil
}

func (c *cTpl) Detail(ctx context.Context, req *v1.TplDetailReq) (res *v1.TplDetailRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	data, err := service.Tpl().Detail(ctx, model.TplDetailInput{
		Id: req.Id,
	})
	if err != nil {
		return res, err
	}

	res = &v1.TplDetailRes{
		Tpl: data,
	}
	return res, nil
}

func (c *cTpl) Update(ctx context.Context, req *v1.TplUpdateReq) (res *v1.OrmUpdateRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	in := model.TplUpdateInput{
		Id: req.Id,
		TplAttr: model.TplAttr{
			UserId:   req.UserId,
			Status:   req.Status,
			Extra:    req.Extra,
			Ip:       req.Ip,
			Path:     req.Path,
			Method:   req.Method,
			Request:  req.Request,
			Response: req.Response,
			Agent:    req.Agent,
			Latency:  req.Latency,
			ErrMsg:   req.ErrMsg,
			UserName: req.UserName,
		},
	}
	_, err = service.Tpl().Update(ctx, in)
	if err != nil {
		return res, err
	}
	res = &v1.OrmUpdateRes{}
	return res, nil
}

func (c *cTpl) Delete(ctx context.Context, req *v1.TplDeleteReq) (res *v1.OrmDeleteRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	_, err = service.Tpl().SafeDelete(ctx, &model.OrmDeleteInput{
		Id: req.Id,
	})
	if err != nil {
		return res, err
	}
	res = &v1.OrmDeleteRes{}
	return res, nil
}

func (c *cTpl) Create(ctx context.Context, req *v1.TplCreateReq) (res *v1.TplCreateRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	in := model.TplCreateInput{
		TplAttr: model.TplAttr{
			UserId:   req.UserId,
			Status:   req.Status,
			Extra:    req.Extra,
			Ip:       req.Ip,
			Path:     req.Path,
			Method:   req.Method,
			Request:  req.Request,
			Response: req.Response,
			Agent:    req.Agent,
			Latency:  req.Latency,
			ErrMsg:   req.ErrMsg,
			UserName: req.UserName,
		},
	}
	id, err := service.Tpl().Create(ctx, &in)
	if err != nil {
		return res, err
	}
	res = &v1.TplCreateRes{}
	res.Id = int(id)
	return res, nil
}

func (c *cTpl) Tree(ctx context.Context, req *v1.TplTreeReq) (res *v1.TplTreeRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	items, err := service.Tpl().GetTree(ctx, model.TplListInput{
		Page:     req.PageReq.Page,
		PageSize: req.PageReq.PageSize,
	})
	if err != nil {
		return nil, err
	}
	cnt, err := service.Tpl().Count(ctx)
	if err != nil {
		return nil, err
	}
	res = &v1.TplTreeRes{
		Count: cnt,
		Items: items,
	}
	return res, nil
}