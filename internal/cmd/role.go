package cmd

import (
	"admin/internal/consts"
	"admin/internal/model/serializer"
	"admin/internal/service"
	"context"
)

var RoleList = []*serializer.Role{
	&serializer.Role{
		Id:        1,
		Name:      "admin",
		Router:    "dashboard",
		IsDeleted: consts.CREATED,
	},
	&serializer.Role{
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
