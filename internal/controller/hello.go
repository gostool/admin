package controller

import (
	"admin/api/v1"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	res = &v1.HelloRes{
		Reply: fmt.Sprintf(`Hi %s`, req.Name),
	}
	err = gerror.New("未知错误，Hello api")
	return
}
