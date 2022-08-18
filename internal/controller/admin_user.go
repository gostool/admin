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

// Detail 获取用户详情
func (c *cAdminUser) Detail(ctx context.Context, req *v1.AdminUserDetailReq) (res *v1.AdminUserDetailRes, err error) {
	userId := common.GetVarFromCtx(ctx, consts.CtxUserId).Int()
	passport := common.GetVarFromCtx(ctx, consts.CtxUserPassport).String()
	userNickname := common.GetVarFromCtx(ctx, consts.CtxUserNickName).String()
	roleId := common.GetVarFromCtx(ctx, consts.CtxUserRoleId).Int()
	role, err := service.Role().Detail(ctx, model.RoleDetailInput{
		Id: roleId,
	})
	if err != nil {
		logger.Fatal(ctx, err)
		return res, err
	}
	res = &v1.AdminUserDetailRes{
		UserDetail: &serializer.UserDetail{
			Id:       userId,
			Passport: passport,
			Nickname: userNickname,
			RoleId:   roleId,
			RoleMap: map[int]*serializer.Role{
				roleId: role,
			},
		},
	}
	return res, nil
}

func (c *cAdminUser) List(ctx context.Context, req *v1.AdminUserListReq) (res *v1.AdminUserListRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	in := model.UserListInput{
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	items, err := service.User().List(ctx, in)
	if err != nil {
		return res, err
	}
	cnt, err := service.User().Count(ctx)
	if err != nil {
		return res, err
	}
	res = &v1.AdminUserListRes{
		Count: cnt,
		Items: items,
	}
	return res, nil
}

func (c *cAdminUser) Update(ctx context.Context, req *v1.AdminUserUpdateReq) (res *v1.OrmUpdateRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	in := model.UserUpdateInput{
		Id: req.Id,
		UserAttr: model.UserAttr{
			Passport: req.Passport,
			Password: req.Password,
			RoleId:   req.RoleId,
			Nickname: req.Nickname,
		},
	}
	_, err = service.User().Update(ctx, in)
	if err != nil {
		return res, err
	}
	res = &v1.OrmUpdateRes{}
	return res, nil
}

func (c *cAdminUser) Delete(ctx context.Context, req *v1.AdminUserDeleteReq) (res *v1.OrmDeleteRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	_, err = service.User().SafeDelete(ctx, &model.OrmDeleteInput{
		Id: req.Id,
	})
	if err != nil {
		return res, err
	}
	res = &v1.OrmDeleteRes{}
	return res, nil
}

func (c *cAdminUser) Create(ctx context.Context, req *v1.AdminUserRegisterReq) (res *v1.AdminUserCreateRes, err error) {
	logger.Debugf(ctx, `receive say: %+v`, req)
	in := model.UserCreateInput{
		RoleIds: req.ToRoleIds(),
		UserAttr: model.UserAttr{
			Passport: req.Passport,
			Password: req.Password,
			RoleId:   req.RoleId,
			Nickname: req.Nickname,
		},
	}
	id, err := service.User().Register(ctx, in)
	if err != nil {
		return res, err
	}
	res = &v1.AdminUserCreateRes{}
	res.Id = int(id)
	return res, nil
}

func (c *cAdminUser) Patch(ctx context.Context, req *v1.AdminUserPatchReq) (res *v1.OrmUpdateRes, err error) {
	in := model.UserPatchInput{
		Id:      req.Id,
		RoleIds: req.ToRoleIds(),
		UserAttr: model.UserAttr{
			Passport: req.Passport,
			Password: req.Password,
			RoleId:   req.RoleId,
			Nickname: req.Nickname,
		},
	}
	_, err = service.User().Patch(ctx, in)
	if err != nil {
		return res, err
	}
	res = &v1.OrmUpdateRes{}
	return res, nil
}
