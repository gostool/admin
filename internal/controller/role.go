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
