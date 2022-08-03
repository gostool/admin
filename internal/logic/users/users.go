package users

import (
	"admin/internal/consts"
	"admin/internal/dao"
	"admin/internal/model"
	"admin/internal/model/entity"
	"admin/internal/service"
	"context"
	"errors"
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

// Login 执行登录
func (s *sUser) Login(ctx context.Context, in model.UserLoginInput) (uid int64, err error) {
	query := g.Map{
		"name":       in.Name,
		"password":   in.Password,
		"is_deleted": consts.CREATED,
	}
	var data *entity.Users
	err = dao.Users.Ctx(ctx).Unscoped().Where(query).Scan(&data)
	if err != nil {
		return uid, err
	}
	if data == nil {
		return uid, errors.New("no data found")
	}
	uid = int64(data.Id)
	return uid, nil
}

// Logout 注销
func (s *sUser) Logout(ctx context.Context) error {
	//return service.Session().RemoveUser(ctx)
	return nil
}
