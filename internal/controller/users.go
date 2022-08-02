package controller

import (
	v1 "admin/api/v1"
	"admin/internal/model"
	"admin/internal/service"
	"context"
)

var (
	User = cUser{}
)

type cUser struct{}

func (c *cUser) Login(ctx context.Context, req *v1.UserReq) (res *v1.UserRes, err error) {
	err = service.User().Login(ctx, model.UserLoginInput{
		Passport: "12",
		Password: "34",
	})
	return nil, err
}
