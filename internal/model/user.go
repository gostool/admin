package model

// UserLoginInput 用户登录
type UserLoginInput struct {
	Name     string // 账号
	Password string // 密码(明文)
}

// UserRegisterInput 创建用户
type UserRegisterInput struct {
	Name     string // 账号
	Password string // 密码(明文)
	Nickname string // 昵称
}
