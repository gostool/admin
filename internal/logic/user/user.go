package user

import (
	"admin/internal/consts"
	"admin/internal/dao"
	"admin/internal/model"
	"admin/internal/model/entity"
	"admin/internal/model/serializer"
	"admin/internal/service"
	"context"
	"database/sql"
	"errors"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

type sUser struct {
	avatarUploadPath      string // 头像上传路径
	avatarUploadUrlPrefix string // 头像上传对应的URL前缀
}

var logger *glog.Logger

func init() {
	logger = g.Log(consts.LoggerDebug)
	user := New()
	service.RegisterUser(user)
}

func New() *sUser {
	return &sUser{}
}

func (s *sUser) FindOne(ctx context.Context, query *g.Map) (data *serializer.User, err error) {
	data = (*serializer.User)(nil)
	// 找不到数据时，不会初始化data
	err = dao.User.Ctx(ctx).Fields(model.UserFields).Where(query).Scan(&data)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return data, errors.New("用户账号或密码错误")
	}
	return data, nil
}

func (s *sUser) Find(ctx context.Context, pk int64) (user *serializer.User, err error) {
	query := g.Map{
		"id":         pk,
		"is_deleted": consts.CREATED,
	}
	return s.FindOne(ctx, &query)
}

func (s *sUser) LoginWeb(ctx context.Context, in model.UserLoginWebInput) (data *serializer.User, err error) {
	ok := service.Tools().Verify(ctx, in.CaptchaId, in.Captcha, true)
	if !ok {
		return nil, errors.New("验证码错误")
	}
	password, err := gmd5.EncryptString(in.Password)
	if err != nil {
		return nil, err
	}
	query := g.Map{
		"passport":   in.Passport,
		"password":   password,
		"is_deleted": consts.CREATED,
	}
	err = dao.User.Ctx(ctx).Unscoped().Fields(model.UserFields).Where(query).Scan(&data)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, errors.New("账号或密码错误，请重试")
	}
	return data, nil
}

// Login 执行登录
func (s *sUser) Login(ctx context.Context, in model.UserLoginInput) (uid int64, err error) {
	password, err := gmd5.EncryptString(in.Password)
	if err != nil {
		return uid, err
	}
	query := g.Map{
		"passport":   in.Passport,
		"password":   password,
		"is_deleted": consts.CREATED,
	}
	var data *entity.User
	err = dao.User.Ctx(ctx).Unscoped().Fields(model.UserFields).Where(query).Scan(&data)
	if err != nil {
		return uid, err
	}
	if data == nil {
		return uid, errors.New("账号或密码错误，请重试")
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

func (s *sUser) Count(ctx context.Context) (data int, err error) {
	query := g.Map{
		"is_deleted": consts.CREATED,
	}
	return dao.User.Ctx(ctx).Count(query)
}

func (s *sUser) Create(ctx context.Context, in model.UserCreateInput) (uid int64, err error) {
	data := g.Map{
		"passport":   in.Passport,
		"password":   in.Password,
		"nickname":   in.Nickname,
		"roleId":     in.RoleId,
		"RoleIds":    in.RoleIds,
		"is_deleted": consts.CREATED,
	}
	r, err := dao.User.Ctx(ctx).Data(data).Insert()
	if err != nil {
		return uid, err
	}
	uid, err = r.LastInsertId()
	if err != nil {
		return uid, err
	}
	return uid, nil
}

func (s *sUser) Save(ctx context.Context, passport, password, nickname string, roleId int, roleIds string) (result sql.Result, err error) {
	data := &g.Map{
		"passport":   passport,
		"password":   password,
		"nickname":   nickname,
		"role_id":    roleId,
		"RoleIds":    roleIds,
		"is_deleted": consts.CREATED,
	}
	return dao.User.Ctx(ctx).Save(data)
}

func (s *sUser) update(ctx context.Context, query, data g.Map) (row int64, err error) {
	logger.Debugf(ctx, "data:%v\n", data)
	result, err := dao.User.Ctx(ctx).Where(query).Update(data)
	if err != nil {
		return 0, err
	}
	row, err = result.RowsAffected()
	if err != nil {
		return row, err
	}
	if row == 0 {
		return 0, consts.ErrUpdate
	}
	return row, err
}

// Patch partial update
func (s *sUser) Patch(ctx context.Context, in model.UserPatchInput) (uid int64, err error) {
	return s.update(ctx, in.ToWhereMap(), in.ToMap())
}

// Update 执行更新
func (s *sUser) Update(ctx context.Context, in model.UserUpdateInput) (uid int64, err error) {
	return s.update(ctx, in.ToWhereMap(), in.ToMap())
}

func (s *sUser) List(ctx context.Context, in model.UserListInput) (items []*serializer.UserInfo, err error) {
	query := g.Map{
		"is_deleted": consts.CREATED,
	}
	err = dao.User.Ctx(ctx).WithAll().Page(in.Page, in.PageSize).Where(query).Scan(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

// Detail 执行详情
func (s *sUser) Detail(ctx context.Context, in model.UserDetailInput) (data *serializer.User, err error) {
	query := g.Map{
		"id":         in.Id,
		"is_deleted": consts.CREATED,
	}
	err = dao.User.Ctx(ctx).Unscoped().Fields(model.UserFields).Where(query).Scan(&data)
	if err != nil {
		return data, err
	}
	if data == nil {
		return data, consts.ErrNotExit
	}
	return data, nil
}

// Delete 执行删除
func (s *sUser) Delete(ctx context.Context, in model.UserDeleteInput) (result sql.Result, err error) {
	query := g.Map{
		"id":         in.Id,
		"is_deleted": consts.CREATED,
	}
	result, err = dao.User.Ctx(ctx).Unscoped().Where(query).Delete()
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *sUser) SafeDelete(ctx context.Context, r *model.OrmDeleteInput) (row int64, err error) {
	row, err = s.update(ctx, r.ToWhereMap(), r.ToMap())
	if err != nil {
		logger.Error(ctx, err)
		return 0, consts.ErrDel
	}
	return row, nil
}

// Logout 注销
func (s *sUser) Logout(ctx context.Context) error {
	//return service.Session().RemoveUser(common)
	return nil
}
