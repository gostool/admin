package testsDB

import (
	"admin/internal/logic/role_menu"
	"admin/internal/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestLogicRoleMenuCount(t *testing.T) {
	ctx := gctx.New()
	r := role_menu.New()
	data, err := r.Count(ctx)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)
}

func TestLogicRoleMenuList(t *testing.T) {
	ctx := gctx.New()
	r := role_menu.New()
	in := model.RoleMenuListInput{
		Page:     1,
		PageSize: 10,
	}
	data, err := r.List(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)
}

func TestLogicRoleMenuCreate(t *testing.T) {
	ctx := gctx.New()
	r := role_menu.New()
	RoleMenuAttr := model.RoleMenuAttr{
		RoleId: 5,
		MenuId: 4,
	}
	in := model.RoleMenuCreateInput{
		RoleMenuAttr: RoleMenuAttr,
	}
	uid, err := r.Create(ctx, &in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(uid)
}

func TestLogicRoleMenuDetail(t *testing.T) {
	ctx := gctx.New()
	r := role_menu.New()
	in := model.RoleMenuDetailInput{
		Id: 1,
	}
	data, err := r.Detail(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)
}

func TestLogicRoleMenuDelete(t *testing.T) {
	ctx := gctx.New()
	r := role_menu.New()
	in := model.OrmDeleteInput{
		Id: 1,
	}
	data, err := r.SafeDelete(ctx, &in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)
}

func TestLogicRoleMenuUpdate(t *testing.T) {
	ctx := gctx.New()
	r := role_menu.New()
	in := model.RoleMenuUpdateInput{
		Id: 1,
		RoleMenuAttr: model.RoleMenuAttr{
			RoleId: 5,
			MenuId: 5,
		},
	}
	id, err := r.Update(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(id)
}
