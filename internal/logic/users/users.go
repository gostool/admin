package users

import (
	"admin/internal/model"
	"admin/internal/service"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
)

type sUser struct {
	avatarUploadPath      string // 头像上传路径
	avatarUploadUrlPrefix string // 头像上传对应的URL前缀
}

func init() {
	user := New()
	fmt.Println(user)
	service.RegisterUser(user)
}

func New() *sUser {
	return &sUser{}
}

// 执行登录
func (s *sUser) Login(ctx context.Context, in model.UserLoginInput) (err error) {
	//userEntity, err := s.GetUserByPassportAndPassword(
	//	ctx,
	//	in.Passport,
	//	s.EncryptPassword(in.Passport, in.Password),
	//)
	g.Log().Info(ctx, "logic login")
	if err != nil {
		return err
	}
	//if userEntity == nil {
	//	return gerror.New(`账号或密码错误`)
	//}
	//if err := service.Session().SetUser(ctx, userEntity); err != nil {
	//	return err
	//}
	// 自动更新上线
	//service.BizCtx().SetUser(ctx, &model.ContextUser{
	//	Id:       userEntity.Id,
	//	Passport: userEntity.Passport,
	//	Nickname: userEntity.Nickname,
	//	Avatar:   userEntity.Avatar,
	//})
	return nil
}

// 注销
func (s *sUser) Logout(ctx context.Context) error {
	//return service.Session().RemoveUser(ctx)
	return nil
}
