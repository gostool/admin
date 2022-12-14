package cmd

import (
	"admin/internal/consts"
	"admin/internal/model"
	"admin/internal/model/entity"
	"admin/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

func menuRoleInit(ctx context.Context) error {
	in := model.RoleDetailInput{
		Id: 1, // 1 is role admin
	}
	role, err := service.Role().Detail(ctx, in)
	if err != nil {
		logger.Debug(ctx, "no such role id")
		panic("please do this first: ./main role")
	}
	g.Dump(role)
	if role.Name != "admin" {
		panic("role.id=1 must be admin")
	}
	menuIdList := GetMenuIdList()
	for _, menuId := range menuIdList {
		roleMenu := &entity.RoleMenu{
			Id:        menuId,
			RoleId:    role.Id,
			MenuId:    menuId,
			IsDeleted: consts.CREATED,
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
