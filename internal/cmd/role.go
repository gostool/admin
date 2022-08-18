package cmd

import (
	"admin/internal/consts"
	"admin/internal/model/entity"
	"admin/internal/service"
	"context"
)

var RoleList = []*entity.Role{
	{
		Id:        1,
		Name:      "admin",
		Router:    "dashboard",
		IsDeleted: consts.CREATED,
	},
	{
		Id:        2,
		Name:      "guest",
		Router:    "dashboard",
		IsDeleted: consts.CREATED,
	},
}

func RoleInit(ctx context.Context) {
	for _, role := range RoleList {
		result, err := service.Role().Save(ctx, role)
		if err != nil {
			logger.Fatal(ctx, err)
		}
		logger.Debugf(ctx, "result:%v", result)
	}
	logger.Debugf(ctx, "Role init")
}
