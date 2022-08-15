package model

// UserLoginInput 用户登录
type UserLoginInput struct {
	Name     string // 账号
	Password string // 密码(明文)
}

type UserLoginWebInput struct {
	Name      string // 账号
	Password  string // 密码(明文)
	Captcha   string
	CaptchaId string
}

// UserCreateInput UserRegisterInput 创建用户
type UserCreateInput struct {
	Name     string // 账号
	Password string // 密码(明文)
	Nickname string // 昵称
	RoleId   int
}

type UserDetailInput struct {
	Id int64
}

type UserDeleteInput struct {
	Id int64
}

type UserUpdateInput struct {
	Id int64
}
