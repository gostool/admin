package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserReq struct {
	g.Meta   `path:"/login" method:"post" tags:"UserService" summary:"login an account"`
	Passport string `v:"required|length:5,16#请输入用户名|用户名称长度应当在:5到:16之间" dc:"Your name" json:"passport"`
	Password string `v:"required|length:5,16#请输入确认密码|密码长度应当在:5到:16之间" dc:"Your password" json:"password"`
}

type UserRegisterReq struct {
	g.Meta   `path:"/register" method:"post" tags:"UserService" summary:"register an account"`
	Name     string `v:"required|length:5,16#请输入用户名|用户名称长度应当在:5到:16之间" dc:"Your name" json:"name"`
	Password string `v:"required|length:5,16#请输入确认密码|密码长度应当在:5到:16之间" dc:"Your password" json:"password"`
	Nickname string `v:"required|length:5,16#请输入昵称|昵称长度应当在:5到:16之间" dc:"Your nickanme" json:"nickname"`
}

//--------------------------------------------------------------------------------------

type UserRes struct {
	Id    int64  `json:"id""`
	Role  int64  `json:"role"`
	Token string `json:"tokens"`
}
