package controller

import (
	v1 "admin/api/v1"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Menu = cMenu{}
)

type cMenu struct{}

func (c *cMenu) List(ctx context.Context, req *v1.MenuListReq) (res *v1.MenuListRes, err error) {
	g.Log().Debugf(ctx, `receive say: %+v`, req)
	return
}
