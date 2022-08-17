package controller

import (
	v1 "admin/api/v1"
	"admin/internal/consts"
	"admin/internal/model"
	"admin/internal/model/serializer"
	"admin/internal/service"
	"github.com/gogf/gf/v2/net/ghttp"
)

var (
	AdminUser = cAdminUser{}
)

type cAdminUser struct{}

func (c *cAdminUser) Profile(r *ghttp.Request) {
	ctx := r.GetCtx()
	userId := r.GetCtxVar(consts.CtxUserId).Int()
	userName := r.GetCtxVar(consts.CtxUserName).String()
	roleId := r.GetCtxVar(consts.CtxUserRoleId).Int()
	role, err := service.Role().Detail(ctx, model.RoleDetailInput{
		Id: roleId,
	})
	if err != nil {
		logger.Fatal(ctx, err)
		r.SetError(err)
		return
	}
	res := &v1.AdminUserWebRes{
		Id:       userId,
		Passport: userName,
		RoleId:   roleId,
		RoleMap: map[int]*serializer.Role{
			roleId: role,
		},
	}
	logger.Debugf(ctx, "role:%v\n", res)
	r.SetCtxVar(consts.CtxResponseData, res)
	return
}
