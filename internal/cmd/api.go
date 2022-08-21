package cmd

import (
	"admin/internal/consts"
	"admin/internal/model/entity"
	"admin/internal/service"
	"context"
)

var apiList = []*entity.Api{
	//role
	{
		Id:        1,
		Path:      "/api/role/list",
		Group:     "role",
		Method:    "GET",
		Detail:    "角色列表",
		IsDeleted: consts.CREATED,
	},
	{
		Id:        2,
		Path:      "/api/role/create",
		Group:     "role",
		Method:    "POST",
		Detail:    "角色创建",
		IsDeleted: consts.CREATED,
	},
	{
		Id:        3,
		Path:      "/api/role/update",
		Group:     "role",
		Method:    "POST",
		Detail:    "角色编辑",
		IsDeleted: consts.CREATED,
	},
	{
		Id:        4,
		Path:      "/api/role/delete",
		Group:     "role",
		Method:    "POST",
		Detail:    "角色删除",
		IsDeleted: consts.CREATED,
	},
	{
		Id:        5,
		Path:      "/api/role/detail",
		Group:     "role",
		Method:    "POST",
		Detail:    "角色详情",
		IsDeleted: consts.CREATED,
	},

	// menu
	{
		Id:        6,
		Path:      "/api/menu/list",
		Group:     "menu",
		Method:    "GET",
		Detail:    "菜单列表:当前用户角色的所有菜单，用于生成动态路由",
		IsDeleted: consts.CREATED,
	},
	{
		Id:        7,
		Path:      "/api/menu/create",
		Group:     "menu",
		Method:    "POST",
		Detail:    "菜单创建",
		IsDeleted: consts.CREATED,
	},
	{
		Id:        8,
		Path:      "/api/menu/update",
		Group:     "menu",
		Method:    "POST",
		Detail:    "菜单编辑",
		IsDeleted: consts.CREATED,
	},
	{
		Id:        9,
		Path:      "/api/menu/delete",
		Group:     "menu",
		Method:    "POST",
		Detail:    "菜单删除",
		IsDeleted: consts.CREATED,
	},
	{
		Id:        10,
		Path:      "/api/menu/detail",
		Group:     "menu",
		Method:    "POST",
		Detail:    "菜单详情",
		IsDeleted: consts.CREATED,
	},

	// log
	{
		Id:        11,
		Path:      "/api/log/list",
		Group:     "log",
		Method:    "GET",
		Detail:    "系统日志",
		IsDeleted: consts.CREATED,
	},

	// api
	{
		Id:        12,
		Path:      "/api/admin_api/list",
		Group:     "api",
		Method:    "GET",
		Detail:    "api列表",
		IsDeleted: consts.CREATED,
	},
	{
		Id:        13,
		Path:      "/api/admin_api/create",
		Group:     "api",
		Method:    "POST",
		Detail:    "api创建",
		IsDeleted: consts.CREATED,
	},
	{
		Id:        14,
		Path:      "/api/admin_api/update",
		Group:     "api",
		Method:    "POST",
		Detail:    "api更新",
		IsDeleted: consts.CREATED,
	},
	{
		Id:        15,
		Path:      "/api/admin_api/detail",
		Group:     "api",
		Method:    "POST",
		Detail:    "api详情",
		IsDeleted: consts.CREATED,
	},
	{
		Id:        16,
		Path:      "/api/admin_api/delete",
		Group:     "api",
		Method:    "POST",
		Detail:    "api删除",
		IsDeleted: consts.CREATED,
	},
	// user
	{
		Id:        17,
		Path:      "/api/admin_user/delete",
		Group:     "user",
		Method:    "POST",
		Detail:    "user删除",
		IsDeleted: consts.CREATED,
	},
	{
		Id:        18,
		Path:      "/api/admin_user/create",
		Group:     "user",
		Method:    "POST",
		Detail:    "user创建",
		IsDeleted: consts.CREATED,
	},
	{
		Id:        19,
		Path:      "/api/admin_user/detail",
		Group:     "user",
		Method:    "GET",
		Detail:    "user详情",
		IsDeleted: consts.CREATED,
	},
	{
		Id:        20,
		Path:      "/api/admin_user/list",
		Group:     "user",
		Method:    "GET",
		Detail:    "user列表",
		IsDeleted: consts.CREATED,
	},
	{
		Id:        21,
		Path:      "/api/admin_user/update",
		Group:     "user",
		Method:    "POST",
		Detail:    "user更新",
		IsDeleted: consts.CREATED,
	},
}

func apiInit(ctx context.Context) {
	for _, obj := range apiList {
		result, err := service.AdminApi().Save(ctx, obj)
		if err != nil {
			logger.Fatal(ctx, err)
		}
		logger.Debugf(ctx, "result:%v", result)
	}
	logger.Debugf(ctx, "api init")
}
