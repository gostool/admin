package controller

import (
	v1 "admin/api/v1"
	"admin/internal/service"
	"context"
)

var (
	Tools = cTools{}
)

type cTools struct{}

func (t *cTools) Captcha(ctx context.Context, req *v1.ToolCaptchaReq) (res *v1.ToolCaptchaRes, err error) {
	id, b64s, err := service.Tools().Captcha(ctx)
	if err != nil {
		return res, err
	}
	res = &v1.ToolCaptchaRes{
		Id:   id,
		Path: b64s,
	}
	return res, err
}
