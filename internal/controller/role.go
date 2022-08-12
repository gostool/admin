package controller

import (
	v1 "admin/api/v1"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Role = cRole{}
)

type cRole struct{}

func (c *cRole) List(ctx context.Context, req *v1.RoleListReq) (res *v1.RoleListRes, err error) {
	g.Log().Debugf(ctx, `receive say: %+v`, req)
	return
}

func (c *cRole) Detail(ctx context.Context, req *v1.RoleDetailReq) (res *v1.RoleDetailRes, err error) {
	g.Log().Debugf(ctx, `receive say: %+v`, req)
	return
}

func (c *cRole) Update(ctx context.Context, req *v1.RoleUpdateReq) (res *v1.RoleUpdateRes, err error) {
	g.Log().Debugf(ctx, `receive say: %+v`, req)
	return
}

func (c *cRole) Delete(ctx context.Context, req *v1.RoleDeleteReq) (res *v1.RoleDeleteRes, err error) {
	g.Log().Debugf(ctx, `receive say: %+v`, req)
	return
}
