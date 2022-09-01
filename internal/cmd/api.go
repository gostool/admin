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
		Path:      "/api/role/list",
		Group:     "role",
		Method:    "GET",
		Detail:    "角色列表",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/role/create",
		Group:     "role",
		Method:    "POST",
		Detail:    "角色创建",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/role/update",
		Group:     "role",
		Method:    "POST",
		Detail:    "角色编辑",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/role/delete",
		Group:     "role",
		Method:    "POST",
		Detail:    "角色删除",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/role/detail",
		Group:     "role",
		Method:    "POST",
		Detail:    "角色详情",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/role/tree",
		Group:     "role",
		Method:    "GET",
		Detail:    "tree详情",
		IsDeleted: consts.CREATED,
	},

	// menu
	{
		Path:      "/api/menu/list",
		Group:     "menu",
		Method:    "GET",
		Detail:    "菜单列表:当前用户角色的所有菜单，用于生成动态路由",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/menu/tree",
		Group:     "menu",
		Method:    "GET",
		Detail:    "菜单tree",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/menu/create",
		Group:     "menu",
		Method:    "POST",
		Detail:    "菜单创建",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/menu/update",
		Group:     "menu",
		Method:    "POST",
		Detail:    "菜单编辑",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/menu/delete",
		Group:     "menu",
		Method:    "POST",
		Detail:    "菜单删除",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/menu/detail",
		Group:     "menu",
		Method:    "POST",
		Detail:    "菜单详情",
		IsDeleted: consts.CREATED,
	},

	// log
	{
		Path:      "/api/log/list",
		Group:     "log",
		Method:    "GET",
		Detail:    "系统日志",
		IsDeleted: consts.CREATED,
	},

	// api
	{
		Path:      "/api/admin_api/list",
		Group:     "api",
		Method:    "GET",
		Detail:    "api列表",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/admin_api/search",
		Group:     "api",
		Method:    "GET",
		Detail:    "api search",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/admin_api/create",
		Group:     "api",
		Method:    "POST",
		Detail:    "api创建",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/admin_api/update",
		Group:     "api",
		Method:    "POST",
		Detail:    "api更新",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/admin_api/detail",
		Group:     "api",
		Method:    "POST",
		Detail:    "api详情",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/admin_api/delete",
		Group:     "api",
		Method:    "POST",
		Detail:    "api删除",
		IsDeleted: consts.CREATED,
	},
	// user
	{
		Path:      "/api/admin_user/delete",
		Group:     "user",
		Method:    "POST",
		Detail:    "user删除",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/admin_user/create",
		Group:     "user",
		Method:    "POST",
		Detail:    "user创建",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/admin_user/detail",
		Group:     "user",
		Method:    "GET",
		Detail:    "user详情",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/admin_user/list",
		Group:     "user",
		Method:    "GET",
		Detail:    "user列表",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/admin_user/update",
		Group:     "user",
		Method:    "POST",
		Detail:    "user更新",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/admin_user/patch",
		Group:     "user",
		Method:    "POST",
		Detail:    "user patch",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/admin_casbin/list",
		Group:     "casbin",
		Method:    "GET",
		Detail:    "casbin列表",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/admin_casbin/update",
		Group:     "casbin",
		Method:    "POST",
		Detail:    "casbin更新",
		IsDeleted: consts.CREATED,
	},
	// role_menu
	{
		Path:      "/api/role_menu/tree",
		Group:     "role_menu",
		Method:    "GET",
		Detail:    "role_menu tree",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/role_menu/list",
		Group:     "role_menu",
		Method:    "GET",
		Detail:    "role_menu列表",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/role_menu/detail",
		Group:     "role_menu",
		Method:    "POST",
		Detail:    "role_menu详情",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/role_menu/update",
		Group:     "role_menu",
		Method:    "POST",
		Detail:    "role_menu更新",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/role_menu/delete",
		Group:     "role_menu",
		Method:    "POST",
		Detail:    "role_menu删除",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/role_menu/create",
		Group:     "role_menu",
		Method:    "POST",
		Detail:    "role_menu创建",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/role_menu/tree",
		Group:     "role_menu",
		Method:    "POST",
		Detail:    "role_menu treeById",
		IsDeleted: consts.CREATED,
	},
	{
		Path:      "/api/status/list",
		Group:     "status",
		Method:    "GET",
		Detail:    "server status",
		IsDeleted: consts.CREATED,
	},
}

func apiInit(ctx context.Context) {
	for i, obj := range apiList {
		obj.Id = i + 1
		result, err := service.AdminApi().Save(ctx, obj)
		if err != nil {
			logger.Fatal(ctx, err)
		}
		logger.Debugf(ctx, "result:%v", result)
	}
	logger.Debugf(ctx, "api init")
}
