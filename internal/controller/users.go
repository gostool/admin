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
	uid, err := service.User().Login(ctx, model.UserLoginInput{
		Name:     req.Name,
		Password: req.Password,
	})
	if err != nil {
		return res, err
	}
	res = &v1.UserRes{
		Id:    uid,
		Role:  1,
		Token: "faf",
	}
	return res, err
}

func (c *cUser) Register(ctx context.Context, req *v1.UserRegisterReq) (res *v1.UserRes, err error) {
	uid, err := service.User().Register(ctx, model.UserCreateInput{
		Name:     req.Name,
		Password: req.Password,
		Nickname: req.Nickname,
	})
	if err != nil {
		return res, err
	}
	res = &v1.UserRes{
		Id:    uid,
		Role:  1,
		Token: "faf",
	}
	return res, err
}
