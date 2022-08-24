package casbinRule

import (
	"admin/internal/consts"
	"admin/internal/dao"
	"admin/internal/model"
	"admin/internal/model/entity"
	"admin/internal/model/serializer"
	"admin/internal/service"
	"context"
	"database/sql"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

type sCasbinRule struct {
}

var logger *glog.Logger

func init() {
	logger = g.Log(consts.LoggerDebug)
	instance := New()
	service.RegisterCasbinRule(instance)
}

func New() *sCasbinRule {
	return &sCasbinRule{}
}

func (s *sCasbinRule) Count(ctx context.Context) (data int, err error) {
	query := g.Map{
		"is_deleted": consts.CREATED,
	}
	return dao.CasbinRule.Ctx(ctx).Count(query)
}

func (s *sCasbinRule) List(ctx context.Context, in model.CasbinRuleListInput) (items []*serializer.CasbinRule, err error) {
	query := g.Map{}
	err = dao.CasbinRule.Ctx(ctx).Fields(model.CasbinRuleFields).Page(in.Page, in.PageSize).Where(query).Scan(&items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *sCasbinRule) InsertAndGetId(ctx context.Context, data g.Map) (id int64, err error) {
	id, err = dao.CasbinRule.Ctx(ctx).InsertAndGetId(data)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (s *sCasbinRule) Create(ctx context.Context, in *model.CasbinRuleCreateInput) (id int64, err error) {
	return s.InsertAndGetId(ctx, in.ToMap())
}

func (s *sCasbinRule) update(ctx context.Context, query, data g.Map) (row int64, err error) {
	logger.Debugf(ctx, "data:%v\n", data)
	result, err := dao.CasbinRule.Ctx(ctx).Where(query).Update(data)
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

// Update 执行更新
func (s *sCasbinRule) Update(ctx context.Context, in model.CasbinRuleUpdateInput) (row int64, err error) {
	return s.update(ctx, in.ToWhereMap(), in.ToMap())
}

// Detail 执行详情
func (s *sCasbinRule) Detail(ctx context.Context, in model.CasbinRuleDetailInput) (data *serializer.CasbinRule, err error) {
	query := g.Map{
		"id":         in.Id,
		"is_deleted": consts.CREATED,
	}
	err = dao.CasbinRule.Ctx(ctx).Unscoped().Fields(model.CasbinRuleFields).Where(query).Scan(&data)
	if err != nil {
		return data, err
	}
	if data == nil {
		return data, consts.ErrNotExit
	}
	return data, nil
}

// Delete 执行删除
func (s *sCasbinRule) Delete(ctx context.Context, in model.OrmDeleteInput) (result sql.Result, err error) {
	query := g.Map{
		"id": in.Id,
	}
	result, err = dao.CasbinRule.Ctx(ctx).Unscoped().Where(query).Delete()
	if err != nil {
		return result, err
	}
	return result, nil
}

func (s *sCasbinRule) Save(ctx context.Context, in *entity.CasbinRule) (result sql.Result, err error) {
	return dao.CasbinRule.Ctx(ctx).Save(in)
}
