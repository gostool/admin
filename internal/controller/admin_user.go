package controller

import (
	v1 "admin/api/v1"
	"admin/internal/common"
	"admin/internal/consts"
	"admin/internal/model"
	"admin/internal/model/serializer"
	"admin/internal/service"
	"context"
)

var (
	AdminUser = cAdminUser{}
)

type cAdminUser struct{}

func (c *cAdminUser) Profile(ctx context.Context, req *v1.AdminUserWebReq) (res *v1.AdminUserWebRes, err error) {
	userId := common.GetVarFromCtx(ctx, consts.CtxUserId).Int()
	userName := common.GetVarFromCtx(ctx, consts.CtxUserName).String()
	roleId := common.GetVarFromCtx(ctx, consts.CtxUserRoleId).Int()
	role, err := service.Role().Detail(ctx, model.RoleDetailInput{
		Id: roleId,
	})
	if err != nil {
		logger.Fatal(ctx, err)
		return res, err
	}
	res = &v1.AdminUserWebRes{
		Id:       userId,
		Passport: userName,
		RoleId:   roleId,
		RoleMap: map[int]*serializer.Role{
			roleId: role,
		},
	}
	return res, nil
}
