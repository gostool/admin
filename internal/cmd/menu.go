package cmd

import (
	"admin/internal/consts"
	"admin/internal/model/entity"
	"admin/internal/service"
	"context"
)

var menuList = []*entity.Menu{
	{
		Id:          1,
		Path:        "dashboard",
		Name:        "dashboard",
		Hidden:      0,
		Component:   "view/dashboard/index.vue",
		Sort:        1,
		KeepAlive:   0,
		DefaultMenu: 0,
		Title:       "仪表盘",
		Icon:        "PieChart",
		IsDeleted:   consts.CREATED,
	},
	{
		Id:          2,
		Path:        "about",
		Name:        "about",
		Hidden:      0,
		Component:   "view/about/index.vue",
		Sort:        7,
		KeepAlive:   0,
		DefaultMenu: 0,
		Title:       "关于我们",
		Icon:        "InfoFilled",
		IsDeleted:   consts.CREATED,
	},
	{
		Id:          3,
		Path:        "admin",
		Name:        "superAdmin",
		Hidden:      0,
		Component:   "view/superAdmin/index.vue",
		Sort:        1,
		KeepAlive:   0,
		DefaultMenu: 0,
		Title:       "超级管理员",
		Icon:        "User",
		IsDeleted:   consts.CREATED,
	},
	{
		Id:          4,
		Path:        "authority",
		Name:        "authority",
		Hidden:      0,
		ParentId:    3,
		Component:   "view/superAdmin/authority/authority.vue",
		Sort:        1,
		KeepAlive:   0,
		DefaultMenu: 0,
		Title:       "角色管理",
		Icon:        "Avatar",
		IsDeleted:   consts.CREATED,
	},
	{
		Id:          5,
		Path:        "menu",
		Name:        "menu",
		Hidden:      0,
		Component:   "view/superAdmin/menu/menu.vue",
		ParentId:    3,
		Sort:        2,
		KeepAlive:   1,
		DefaultMenu: 0,
		Title:       "菜单管理",
		Icon:        "Menu",
		IsDeleted:   consts.CREATED,
	},
	{
		Id:          6,
		Path:        "api",
		Name:        "api",
		Hidden:      0,
		ParentId:    3,
		Component:   "view/superAdmin/api/api.vue",
		Sort:        3,
		KeepAlive:   1,
		DefaultMenu: 0,
		Title:       "api管理",
		Icon:        "Monitor",
		IsDeleted:   consts.CREATED,
	},
	{
		Id:          7,
		Path:        "user",
		Name:        "user",
		ParentId:    3,
		Hidden:      0,
		Component:   "view/superAdmin/user/user.vue",
		Sort:        4,
		KeepAlive:   0,
		DefaultMenu: 0,
		Title:       "用户管理",
		Icon:        "User",
		IsDeleted:   consts.CREATED,
	},
	{
		Id:          8,
		Path:        "person",
		Name:        "person",
		ParentId:    0,
		Hidden:      0,
		Component:   "view/person/person.vue",
		Sort:        4,
		KeepAlive:   0,
		DefaultMenu: 0,
		Title:       "个人信息",
		Icon:        "User",
		IsDeleted:   consts.CREATED,
	},
	{
		Id:          9,
		Path:        "example",
		Name:        "example",
		ParentId:    0,
		Hidden:      0,
		Component:   "view/example/index.vue",
		Sort:        6,
		KeepAlive:   0,
		DefaultMenu: 0,
		Title:       "示例文件",
		Icon:        "Management",
		IsDeleted:   consts.CREATED,
	},
	{
		Id:          10,
		Path:        "excel",
		Name:        "excel",
		ParentId:    9,
		Hidden:      0,
		Component:   "view/example/excel/excel.vue",
		Sort:        4,
		KeepAlive:   0,
		DefaultMenu: 0,
		Title:       "excel导入导出",
		Icon:        "Files",
		IsDeleted:   consts.CREATED,
	},
	{
		Id:          11,
		Path:        "upload",
		Name:        "upload",
		ParentId:    9,
		Hidden:      0,
		Component:   "view/example/upload/upload.vue",
		Sort:        5,
		KeepAlive:   0,
		DefaultMenu: 0,
		Title:       "媒体库（上传下载）",
		Icon:        "UploadFilled",
		IsDeleted:   consts.CREATED,
	},
	{
		Id:          12,
		Path:        "breakpoint",
		Name:        "breakpoint",
		ParentId:    9,
		Hidden:      0,
		Component:   "view/example/breakpoint/breakpoint.vue",
		Sort:        6,
		KeepAlive:   0,
		DefaultMenu: 0,
		Title:       "断点续传",
		Icon:        "UploadFilled",
		IsDeleted:   consts.CREATED,
	},
	{
		Id:          13,
		Path:        "customer",
		Name:        "customer",
		ParentId:    9,
		Hidden:      0,
		Component:   "view/example/customer/customer.vue",
		Sort:        7,
		KeepAlive:   0,
		DefaultMenu: 0,
		Title:       "客户列表（资源示例）",
		Icon:        "User",
		IsDeleted:   consts.CREATED,
	},
	{
		Id:          14,
		Path:        "systemTools",
		Name:        "systemTools",
		ParentId:    0,
		Hidden:      0,
		Component:   "view/systemTools/index.vue",
		Sort:        5,
		KeepAlive:   0,
		DefaultMenu: 0,
		Title:       "系统工具",
		Icon:        "Setting",
		IsDeleted:   consts.CREATED,
	},
	{
		Id:          15,
		Path:        "autoCode",
		Name:        "autoCode",
		Hidden:      0,
		ParentId:    14,
		Component:   "view/systemTools/autoCode/index.vue",
		Sort:        1,
		KeepAlive:   1,
		DefaultMenu: 0,
		Title:       "代码生成器",
		Icon:        "Cpu",
		IsDeleted:   consts.CREATED,
	},
	{
		Id:          16,
		Path:        "formCreate",
		Name:        "formCreate",
		Hidden:      0,
		ParentId:    14,
		Component:   "view/systemTools/formCreate/index.vue",
		Sort:        2,
		KeepAlive:   1,
		DefaultMenu: 0,
		Title:       "表单生成器",
		Icon:        "MagicStick",
		IsDeleted:   consts.CREATED,
	},
	{
		Id:          17,
		Path:        "system",
		Name:        "system",
		Hidden:      0,
		ParentId:    14,
		Component:   "view/systemTools/system/system.vue",
		Sort:        3,
		KeepAlive:   1,
		DefaultMenu: 0,
		Title:       "系统配置",
		Icon:        "Setting",
		IsDeleted:   consts.CREATED,
	},

	{
		Id:          18,
		Path:        "dictionary",
		Name:        "dictionary",
		ParentId:    3,
		Hidden:      0,
		Component:   "view/superAdmin/dictionary/sysDictionary.vue",
		Sort:        5,
		KeepAlive:   1,
		DefaultMenu: 0,
		Title:       "字典管理",
		Icon:        "Notebook",
		IsDeleted:   consts.CREATED,
	},

	{
		Id:          19,
		Path:        "dictionaryDetail/:id",
		Name:        "dictionaryDetail",
		ParentId:    3,
		Hidden:      1,
		Component:   "view/superAdmin/dictionary/sysDictionaryDetail.vue",
		Sort:        1,
		KeepAlive:   1,
		DefaultMenu: 0,
		Title:       "字典详情",
		Icon:        "ShoppingBag",
		IsDeleted:   consts.CREATED,
	},
	{
		Id:          20,
		Path:        "operation",
		Name:        "operation",
		ParentId:    3,
		Hidden:      0,
		Component:   "view/superAdmin/operation/sysOperationRecord.vue",
		Sort:        6,
		KeepAlive:   1,
		DefaultMenu: 0,
		Title:       "操作历史",
		Icon:        "Timer",
		IsDeleted:   consts.CREATED,
	},
	{
		Id:          21,
		Path:        "simpleUploader",
		Name:        "simpleUploader",
		ParentId:    9,
		Hidden:      0,
		Component:   "view/example/simpleUploader/simpleUploader",
		Sort:        6,
		KeepAlive:   1,
		DefaultMenu: 0,
		Title:       "断点续传（插件版）",
		Icon:        "UploadFilled",
		IsDeleted:   consts.CREATED,
	},
	{
		Id:          22,
		Path:        "state",
		Name:        "state",
		ParentId:    0,
		Hidden:      0,
		Component:   "view/system/state.vue",
		Sort:        6,
		KeepAlive:   0,
		DefaultMenu: 0,
		Title:       "服务器状态",
		Icon:        "Cloudy",
		IsDeleted:   consts.CREATED,
	},
}

func menuInit(ctx context.Context) {
	for _, menu := range menuList {
		result, err := service.Menu().Save(ctx, menu)
		if err != nil {
			logger.Fatal(ctx, err)
		}
		logger.Debugf(ctx, "result:%v", result)
	}
	logger.Debugf(ctx, "Menu init")
}
