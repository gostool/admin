package middleware

import (
	"admin/internal/consts"
	jwt "admin/internal/logic/token"
	"admin/internal/logic/user"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

// JwtAuth 检查jwt token是否有效
func JwtAuth(r *ghttp.Request) {
	ctx := r.GetCtx()
	token := r.GetHeader("token")
	data := gmap.Map{}
	if token == "" {
		JsonExit(r, consts.ErrCodeToken, "token 不能为空", data)
	}
	uid, err := jwt.New().CheckToken(ctx, token)
	if err != nil {
		JsonExit(r, consts.ErrCodeToken, "token验证失败", data)
	}
	// 必须在数据库验证存在用户
	user, err := user.New().Find(ctx, gconv.Int64(uid))
	if err != nil {
		JsonExit(r, consts.ErrCodeToken, "用户不存在", data)
	} else {
		//认证成功后，配置参数
		r.SetCtxVar(consts.CtxUserId, user.Id)
		r.SetCtxVar(consts.CtxUserRoleId, user.RoleId)
		r.SetCtxVar(consts.CtxUserPassport, user.Passport)
		r.SetCtxVar(consts.CtxUserNickName, user.Nickname)
		r.SetCtxVar(consts.CtxUserData, user.ToData())
		r.SetCtxVar(consts.CtxUserToken, token)
		r.Middleware.Next()
	}
}
