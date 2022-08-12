package role

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
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type sRole struct {
}

func init() {
	instance := New()
	fmt.Println(instance)
	service.RegisterRole(instance)
}

func New() *sRole {
	return &sRole{}
}

func (s *sRole) Count(ctx context.Context) (data int, err error) {
	query := g.Map{
		"is_deleted": consts.CREATED,
	}
	return dao.Role.Ctx(ctx).Count(query)
}

func (s *sRole) List(ctx context.Context, in model.RoleListInput) (data gdb.Result, err error) {
	query := g.Map{
		"is_deleted": consts.CREATED,
	}
	return dao.Role.Ctx(ctx).Fields(model.RoleFields).Page(in.Page, in.PageSize).Where(query).All()
}

func (s *sRole) InsertAndGetId(ctx context.Context, data g.Map) (id int64, err error) {
	id, err = dao.Role.Ctx(ctx).InsertAndGetId(data)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (s *sRole) Create(ctx context.Context, in *model.RoleCreateInput) (id int64, err error) {
	return s.InsertAndGetId(ctx, in.ToMap())
}

func (s *sRole) update(ctx context.Context, query, data g.Map) (row int64, err error) {
	result, err := dao.Role.Ctx(ctx).Where(query).Update(data)
	if err != nil {
		return 0, err
	}
	row, err = result.RowsAffected()
	if err != nil {
		return row, err
	}
	if row == 0 {
		return 0, errors.New("update fail")
	}
	return row, err
}

// Update 执行更新
func (s *sRole) Update(ctx context.Context, in model.RoleUpdateInput) (id int64, err error) {
	return s.update(ctx, in.ToWhereMap(), in.ToMap())
}

// Detail 执行详情
func (s *sRole) Detail(ctx context.Context, in model.RoleDetailInput) (data *entity.Role, err error) {
	query := g.Map{
		"id":         in.Id,
		"is_deleted": consts.CREATED,
	}
	err = dao.Role.Ctx(ctx).Unscoped().Fields(model.RoleFields).Where(query).Scan(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// Delete 执行删除
func (s *sRole) Delete(ctx context.Context, in model.RoleDeleteInput) (result sql.Result, err error) {
	query := g.Map{
		"id":         in.Id,
		"is_deleted": consts.CREATED,
	}
	result, err = dao.Role.Ctx(ctx).Unscoped().Where(query).Delete()
	if err != nil {
		return result, err
	}
	return result, nil
}
