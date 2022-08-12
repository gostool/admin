package tpl

import (
	"admin/internal/consts"
	"admin/internal/dao"
	"admin/internal/model"
	"admin/internal/model/entity"
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/frame/g"
)

type sModel struct {
	avatarUploadPath      string // 头像上传路径
	avatarUploadUrlPrefix string // 头像上传对应的URL前缀
}

func init() {
	//instance := New()
	//service.RegisterUser(instance)
}

func New() *sModel {
	return &sModel{}
}

func (s *sModel) Create(ctx context.Context, in model.UserCreateInput) (uid int64, err error) {
	data := g.Map{
		"name":       in.Name,
		"password":   in.Password,
		"nickname":   in.Nickname,
		"is_deleted": consts.CREATED,
	}
	r, err := dao.User.Ctx(ctx).Data(data).Insert()
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
func (s *sModel) Update(ctx context.Context, in model.UserUpdateInput) (uid int64, err error) {
	return
}

// Detail 执行详情
func (s *sModel) Detail(ctx context.Context, in model.UserDetailInput) (data *entity.User, err error) {
	query := g.Map{
		"id":         in.Id,
		"is_deleted": consts.CREATED,
	}
	err = dao.User.Ctx(ctx).Unscoped().Fields(model.UserFields).Where(query).Scan(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// Delete 执行删除
func (s *sModel) Delete(ctx context.Context, in model.UserDeleteInput) (result sql.Result, err error) {
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

