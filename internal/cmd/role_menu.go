package cmd

import (
	"admin/internal/model"
	"admin/internal/model/serializer"
	"admin/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func menuRoleInit(ctx context.Context) error {
	in := model.RoleDetailInput{
		Id: 1,
	}
	role, err := service.Role().Detail(ctx, in)
	if err != nil {
		logger.Debug(ctx, "no such role id")
		panic("please do this first: ./main role")
	}
	g.Dump(role)
	menuIdList := GetMenuIdList()
	for _, menuId := range menuIdList {
		roleMenu := &serializer.RoleMenu{
			Id:        menuId,
			RoleId:    role.Id,
			MenuId:    menuId,
			IsDeleted: 0,
		}
		r, err := service.RoleMenu().Save(ctx, roleMenu)
		if err != nil {
			logger.Fatal(ctx, err)
		}
		logger.Debugf(ctx, "r:%v", r)
	}
	return nil
}

func GetMenuIdList() []int {
	menuIdList := make([]int, 0, len(menuList))
	for _, menu := range menuList {
		menuIdList = append(menuIdList, menu.Id)
	}
	return menuIdList
}
