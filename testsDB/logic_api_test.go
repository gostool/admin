package testsDB

import (
	"admin/internal/logic/admin_api"
	"admin/internal/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestLogicApiList(t *testing.T) {
	ctx := gctx.New()
	r := admin_api.New()
	in := model.AdminApiListInput{
		Page:     1,
		PageSize: 10,
	}
	data, err := r.List(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)
}

func TestLogicApiCreate(t *testing.T) {
	ctx := gctx.New()
	r := admin_api.New()
	in := model.AdminApiCreateInput{
		AdminApiAttr: model.AdminApiAttr{
			Group:  "1",
			Path:   "/123",
			Detail: "1",
			Method: "get",
		},
	}
	uid, err := r.Create(ctx, &in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(uid)
}

func TestLogicApiDetail(t *testing.T) {
	ctx := gctx.New()
	r := admin_api.New()
	in := model.AdminApiDetailInput{
		Id: 3,
	}
	data, err := r.Detail(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)
}

func TestLogicApiDelete(t *testing.T) {
	ctx := gctx.New()
	r := admin_api.New()
	in := model.OrmDeleteInput{
		Id: 3,
	}
	data, err := r.SafeDelete(ctx, &in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)
}

func TestLogicApiUpdate(t *testing.T) {
	ctx := gctx.New()
	r := admin_api.New()
	in := model.AdminApiUpdateInput{
		Id: 3,
		AdminApiAttr: model.AdminApiAttr{
			Group:  "1",
			Path:   "/333",
			Detail: "1",
			Method: "get",
		},
	}
	id, err := r.Update(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(id)
}
