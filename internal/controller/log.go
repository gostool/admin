package controller

import (
	v1 "admin/api/v1"
	"admin/internal/model"
	"admin/internal/service"
	"context"
)

var (
	Log = cLog{}
)

type cLog struct{}

func (c *cLog) List(ctx context.Context, req *v1.LogListReq) (res *v1.LogListRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	items, err := service.Log().List(ctx, model.LogListInput{
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
	res = &v1.LogListRes{
		Count: cnt,
		Items: items,
	}
	return res, nil
}

func (c *cLog) Detail(ctx context.Context, req *v1.LogDetailReq) (res *v1.LogDetailRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	data, err := service.Log().Detail(ctx, model.LogDetailInput{
		Id: req.Id,
	})
	if err != nil {
		return res, err
	}

	res = &v1.LogDetailRes{
		Log: data,
	}
	return res, nil
}

func (c *cLog) Update(ctx context.Context, req *v1.LogUpdateReq) (res *v1.OrmUpdateRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	in := model.LogUpdateInput{
		Id: req.Id,
		LogAttr: model.LogAttr{
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
	_, err = service.Log().Update(ctx, in)
	if err != nil {
		return res, err
	}
	res = &v1.OrmUpdateRes{}
	return res, nil
}

func (c *cLog) Delete(ctx context.Context, req *v1.LogDeleteReq) (res *v1.OrmDeleteRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	_, err = service.Log().SafeDelete(ctx, &model.OrmDeleteInput{
		Id: req.Id,
	})
	if err != nil {
		return res, err
	}
	res = &v1.OrmDeleteRes{}
	return res, nil
}

func (c *cLog) Create(ctx context.Context, req *v1.LogCreateReq) (res *v1.LogCreateRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	in := model.LogCreateInput{
		LogAttr: model.LogAttr{
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
	id, err := service.Log().Create(ctx, &in)
	if err != nil {
		return res, err
	}
	res = &v1.LogCreateRes{}
	res.Id = int(id)
	return res, nil
}
