package middleware

import (
	"admin/internal/common"
	"admin/internal/consts"
	"admin/internal/service"
	"errors"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

// Casbin
// RBAC 访问控制
func Casbin(r *ghttp.Request) {
	method := r.Request.Method // 获取请求方法
	url := r.Request.URL.Path  // 获取请求的URI
	ctx := r.GetCtx()
	roleId := common.GetVarFromCtx(ctx, consts.CtxUserRoleId).Int()
	logger.Debugf(ctx, "roleId:%v method:%v url:%v is not admin\n", roleId, method, url)
	data := gmap.Map{}
	e := service.Enforcer().Enforcer(ctx)
	success, err := e.Enforce(gconv.String(roleId), url, method)
	if err != nil {
		logger.Errorf(ctx, "Enforce err:%v\n", err)
		JsonExit(r, consts.ErrCodeErr, consts.ErrMsgErr, data)
	}
	if success {
		r.Middleware.Next()
		return
	} else {
		err = errors.New("权限不足")
		logger.Errorf(ctx, "roleId:%v method:%v url:%v err:%v\n", roleId, method, url, err)
		JsonExit(r, consts.ErrCodeBad, err.Error(), data)
	}
}
