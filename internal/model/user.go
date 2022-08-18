package model

import (
	"admin/internal/consts"
	"github.com/gogf/gf/v2/frame/g"
)

// UserLoginInput 用户登录
type UserLoginInput struct {
	Passport string // 账号
	Password string // 密码(明文)
}

type UserLoginWebInput struct {
	Passport  string // 账号
	Password  string // 密码(明文)
	Captcha   string
	CaptchaId string
}

type UserListInput struct {
	Page     int
	PageSize int
}

// UserCreateInput UserRegisterInput 创建用户
type UserCreateInput struct {
	UserAttr
}

type UserDetailInput struct {
	Id int
}

type UserDeleteInput struct {
	Id int
}

type UserUpdateInput struct {
	Id int
	UserAttr
}

type UserAttr struct {
	Passport string // 账号
	Password string // 密码(明文)
	Nickname string // 昵称
	RoleId   int
}

func (r *UserUpdateInput) ToWhereMap() (data g.Map) {
	data = g.Map{
		"id":         r.Id,
		"is_deleted": consts.CREATED,
	}
	return data
}

func (r *UserUpdateInput) ToMap() (data g.Map) {
	data = g.Map{}
	data["passport"] = r.Passport
	data["password"] = r.Password
	data["nickname"] = r.Nickname
	data["role_id"] = r.RoleId
	return data
}
