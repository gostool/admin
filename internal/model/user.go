package model

// UserLoginInput 用户登录
type UserLoginInput struct {
	Passport string // 账号
	Password string // 密码(明文)
}
