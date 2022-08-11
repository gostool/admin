package controller

import (
	v1 "admin/api/v1"
	"admin/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Tools = cTools{}
)

type cTools struct{}

func (c *cTools) Captcha(ctx context.Context, req *v1.ToolsCaptchaReq) (res *v1.ToolsCaptchaRes, err error) {
	g.Log().Debugf(ctx, `receive say: %+v`, req)
	id, b64s, err := service.Tools().Captcha(ctx)
	if err != nil {
		return res, err
	}
	res = &v1.ToolsCaptchaRes{
		Id:   id,
		Path: b64s,
	}
	return
}
