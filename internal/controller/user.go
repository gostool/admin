package controller

import (
	v1 "admin/api/v1"
	"admin/internal/model"
	"admin/internal/service"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	User = cUser{}
)

type cUser struct{}

func (c *cUser) LoginWeb(ctx context.Context, req *v1.UserWebReq) (res *v1.UserWebRes, err error) {
	udata, err := service.User().LoginWeb(ctx, model.UserLoginInput{
		Name:     req.Passport,
		Password: req.Password,
	})
	if err != nil {
		return res, err
	}
	token, err := service.Token().GenToken(ctx, gconv.String(udata.Id), 0)
	if err != nil {
		return res, err
	}
	rdata, err := service.Role().Detail(ctx, model.RoleDetailInput{
		Id: udata.RoleId,
	})
	if err != nil {
		return nil, err
	}
	res = &v1.UserWebRes{
		Id:      udata.Id,
		RoleId:  udata.RoleId,
		RoleMap: rdata,
		Token:   token,
	}
	return res, err
}

func (c *cUser) Login(ctx context.Context, req *v1.UserReq) (res *v1.UserRes, err error) {
	uid, err := service.User().Login(ctx, model.UserLoginInput{
		Name:     req.Passport,
		Password: req.Password,
	})
	if err != nil {
		return res, err
	}
	token, err := service.Token().GenToken(ctx, gconv.String(uid), 0)
	if err != nil {
		return res, err
	}
	res = &v1.UserRes{
		Id:    int(uid),
		Role:  1,
		Token: token,
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
		Id:    int(uid),
		Role:  1,
		Token: "faf",
	}
	return res, err
}
