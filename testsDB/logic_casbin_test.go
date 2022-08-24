package testsDB

import (
	casbinRule "admin/internal/logic/casbin_rule"
	"admin/internal/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestLogicCasbinCount(t *testing.T) {
	ctx := gctx.New()
	r := casbinRule.New()
	data, err := r.Count(ctx)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)
}

func TestLogicCasbinList(t *testing.T) {
	ctx := gctx.New()
	r := casbinRule.New()
	in := model.CasbinRuleListInput{
		Page:     1,
		PageSize: 10,
	}
	data, err := r.List(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)
}

func TestLogicCasbinCreate(t *testing.T) {
	ctx := gctx.New()
	r := casbinRule.New()
	attr := model.CasbinRuleAttr{
		Ptype: "p",
		V0:    "tp",
		V1:    "data1",
		V2:    "read",
	}
	in := model.CasbinRuleCreateInput{
		CasbinRuleAttr: attr,
	}
	uid, err := r.Create(ctx, &in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(uid)
}

func TestLogicCasbinDelete(t *testing.T) {
	ctx := gctx.New()
	r := casbinRule.New()
	in := model.OrmDeleteInput{
		Id: 2,
	}
	data, err := r.Delete(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)
}
