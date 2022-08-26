package cmd

import (
	"admin/internal/model/entity"
	"admin/internal/service"
	"context"
	"github.com/gogf/gf/v2/util/gconv"
)

func permInit(ctx context.Context) {
	roleIdStr := gconv.String(adminRoleId)
	service.Enforcer().Clear(0, roleIdStr)
	for i, obj := range apiList {
		in := &entity.CasbinRule{
			Id:    uint64(i + 1),
			Ptype: "p",
			V0:    roleIdStr,
			V1:    obj.Path,
			V2:    obj.Method,
		}
		result, err := service.CasbinRule().Save(ctx, in)
		if err != nil {
			logger.Fatal(ctx, err)
		}
		logger.Debugf(ctx, "result:%v", result)
	}
	logger.Debugf(ctx, "api init")
}
