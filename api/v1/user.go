package v1

import (
	"admin/internal/model/serializer"
	"github.com/gogf/gf/v2/frame/g"
)

type UserReq struct {
	g.Meta   `path:"/user/login" method:"post" tags:"UserService" summary:"login an account"`
	Passport string `v:"required|length:5,16#请输入用户名|用户名称长度应当在:5到:16之间" dc:"Your name" json:"passport"`
	Password string `v:"required|length:5,16#请输入确认密码|密码长度应当在:5到:16之间" dc:"Your password" json:"password"`
}

type UserRegisterReq struct {
	g.Meta   `path:"/user/register" method:"post" tags:"UserService" summary:"register an account"`
	Name     string `v:"required|length:5,16#请输入用户名|用户名称长度应当在:5到:16之间" dc:"Your name" json:"name"`
	Password string `v:"required|length:5,16#请输入确认密码|密码长度应当在:5到:16之间" dc:"Your password" json:"password"`
	Nickname string `v:"required|length:5,16#请输入昵称|昵称长度应当在:5到:16之间" dc:"Your nickanme" json:"nickname"`
}

type UserWebReq struct {
	g.Meta    `path:"/user/webLogin" method:"post" tags:"UserService" summary:"login an account"`
	Passport  string `v:"required|length:5,16#请输入用户名|用户名称长度应当在:5到:16之间" dc:"Your name" json:"passport"`
	Password  string `v:"required|length:5,16#请输入确认密码|密码长度应当在:5到:16之间" dc:"Your password" json:"password"`
	Captcha   string `v:"required|length:4,4#请输入4位验证码" dc:"your captcha" json:"captcha"`
	CaptchaId string `v:"required|length:1,40#请输入captchaId" dc:"your captchaId" json:"captchaId"`
}

type UserWebRes struct {
	Id      int                      `json:"id"`
	RoleId  int                      `json:"roleId"`
	RoleMap map[int]*serializer.Role `json:"roleMap"`
	Token   string                   `json:"token"`
}

//--------------------------------------------------------------------------------------

type UserRes struct {
	Id    int    `json:"id"`
	Role  int    `json:"role"`
	Token string `json:"token"`
}
