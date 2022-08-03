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
		Name:     "12",
		Password: "34",
	})
	res = &v1.UserRes{
		Id:    1,
		Role:  1,
		Token: "faf",
	}
	return res, err
}
