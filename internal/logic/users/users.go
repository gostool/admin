package users

import (
	"admin/internal/consts"
	"admin/internal/dao"
	"admin/internal/model"
	"admin/internal/model/entity"
	"admin/internal/service"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
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
	password, err := gmd5.EncryptString(in.Password)
	if err != nil {
		return uid, err
	}
	query := g.Map{
		"name":       in.Name,
		"password":   password,
		"is_deleted": consts.CREATED,
	}
	var data *entity.Users
	err = dao.Users.Ctx(ctx).Unscoped().Fields(model.UserFields).Where(query).Scan(&data)
	if err != nil {
		return uid, err
	}
	if data == nil {
		return uid, errors.New("no data found")
	}
	uid = int64(data.Id)
	return uid, nil
}

// user reset api curd
// get /api/user/ 返回用户列表
// post /api/user/ 创建用户
// get /api/user/{pk}/ 获取pk用户详情
// put /api/user/pk/
//  delete /api/user/pk/ 删除确定用户
//  delete /api/user [] 删除确定用户

// Register 执行注册
func (s *sUser) Register(ctx context.Context, in model.UserCreateInput) (uid int64, err error) {
	password, err := gmd5.EncryptString(in.Password)
	if err != nil {
		return uid, err
	}
	in.Password = password
	uid, err = s.Create(ctx, in)
	return
}

func (s *sUser) Create(ctx context.Context, in model.UserCreateInput) (uid int64, err error) {
	data := g.Map{
		"name":       in.Name,
		"password":   in.Password,
		"nickname":   in.Nickname,
		"is_deleted": consts.CREATED,
	}
	r, err := dao.Users.Ctx(ctx).Data(data).Insert()
	if err != nil {
		return uid, nil
	}
	uid, err = r.LastInsertId()
	if err != nil {
		return uid, err
	}
	return uid, nil
}

// Update 执行更新
func (s *sUser) Update(ctx context.Context, in model.UserUpdateInput) (uid int64, err error) {
	return
}

// Detail 执行详情
func (s *sUser) Detail(ctx context.Context, in model.UserDetailInput) (data *entity.Users, err error) {
	query := g.Map{
		"id":         in.Id,
		"is_deleted": consts.CREATED,
	}
	err = dao.Users.Ctx(ctx).Unscoped().Fields(model.UserFields).Where(query).Scan(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// Delete 执行删除
func (s *sUser) Delete(ctx context.Context, in model.UserDeleteInput) (result sql.Result, err error) {
	query := g.Map{
		"id":         in.Id,
		"is_deleted": consts.CREATED,
	}
	result, err = dao.Users.Ctx(ctx).Unscoped().Where(query).Delete()
	if err != nil {
		return result, err
	}
	return result, nil
}

// Logout 注销
func (s *sUser) Logout(ctx context.Context) error {
	//return service.Session().RemoveUser(ctx)
	return nil
}
