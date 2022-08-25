package middleware

import (
	"admin/internal/common"
	"admin/internal/consts"
	"github.com/gogf/gf/v2/net/ghttp"
)

// Casbin
// RBAC 访问控制
func Casbin(r *ghttp.Request) {
	method := r.Request.Method        // 获取请求方法
	url := r.Request.URL.RequestURI() // 获取请求的URI
	ctx := r.GetCtx()
	roleId := common.GetVarFromCtx(ctx, consts.CtxUserRoleId).Int()
	logger.Debugf(ctx, "roleId:%v method:%v url:%v is not admin\n", roleId, method, url)
	//data := gmap.Map{}
	//e, err := Casbin.GetEnforcer()
	//if err != nil {
	//	logger.Errorf(ctx, "GetEnforcer err:%v\n", err)
	//	JsonExit(r, consts.ErrCodeErr, consts.ErrMsgErr, data)
	//}
	//success, err := e.Enforce(gconv.String(roleId), url, method)
	//if err != nil {
	//	logger.Errorf(ctx, "Enforce err:%v\n", err)
	//	JsonExit(r, consts.ErrCodeErr, consts.ErrMsgErr, data)
	//}
	//if success {
	//	r.Middleware.Next()
	//	return
	//} else {
	//	err = errors.New("权限不足")
	//	logger.Errorf(ctx, "roleId:%v method:%v url:%v err:%v\n", roleId, method, url, err)
	//	JsonExit(r, consts.ErrCodeBad, err.Error(), data)
	//}
}
