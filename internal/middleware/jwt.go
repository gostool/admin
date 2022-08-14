package middleware

import (
	jwt "admin/internal/logic/token"
	"admin/internal/logic/user"
	"admin/library/response"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

// JwtAuth 检查jwt token是否有效
func JwtAuth(r *ghttp.Request) {
	ctx := r.GetCtx()
	token := r.GetHeader("token")
	if token == "" {
		response.JsonErrExit(ctx, r, response.ERR_CODE_TOKEN, "token 不能为空")
	}
	uid, err := jwt.New().CheckToken(ctx, token)
	if err != nil {
		response.JsonErrExit(ctx, r, response.ERR_CODE_TOKEN, err.Error())
	}
	// 必须在数据库验证存在用户
	user, err := user.New().Find(ctx, gconv.Int64(uid))
	if err != nil {
		response.JsonErrExit(ctx, r, response.ERR_CODE_TOKEN, err.Error())
	} else {
		//认证成功后，配置参数
		r.SetParam("uid", user.Id)
		r.SetParam("roleId", user.RoleId)
		r.SetParam("name", user.Name)
		r.SetParam("data", user.ToData())
		r.SetParam("token", token)
		r.Middleware.Next()
	}
}
