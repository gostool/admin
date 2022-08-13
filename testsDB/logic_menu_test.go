package testsDB

import (
	"admin/internal/logic/menu"
	"admin/internal/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestLogicMenuCount(t *testing.T) {
	ctx := gctx.New()
	r := menu.New()
	data, err := r.Count(ctx)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)
}

func TestLogicMenuList(t *testing.T) {
	ctx := gctx.New()
	r := menu.New()
	in := model.MenuListInput{
		Page:     1,
		PageSize: 10,
	}
	data, err := r.List(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)
}

func TestLogicMenuCreate(t *testing.T) {
	ctx := gctx.New()
	r := menu.New()
	MenuAttr := model.MenuAttr{
		Title:     "HI",
		Icon:      "HI",
		Name:      "HI",
		Component: "a.vue",
	}
	in := model.MenuCreateInput{
		MenuAttr: MenuAttr,
	}
	uid, err := r.Create(ctx, &in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(uid)
}

func TestLogicMenuDetail(t *testing.T) {
	ctx := gctx.New()
	r := menu.New()
	in := model.MenuDetailInput{
		Id: 3,
	}
	data, err := r.Detail(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)
}

func TestLogicMenuDelete(t *testing.T) {
	ctx := gctx.New()
	r := menu.New()
	in := model.OrmDeleteInput{
		Id: 4,
	}
	data, err := r.SafeDelete(ctx, &in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)
}

func TestLogicMenuUpdate(t *testing.T) {
	ctx := gctx.New()
	r := menu.New()
	in := model.MenuUpdateInput{
		Id: 4,
		MenuAttr: model.MenuAttr{
			Title:     "HI",
			Icon:      "HI",
			Name:      "Hello",
			Component: "a.vue",
		},
	}
	id, err := r.Update(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(id)
}
