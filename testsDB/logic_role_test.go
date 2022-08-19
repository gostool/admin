package testsDB

import (
	"admin/internal/logic/role"
	"admin/internal/model"
	"admin/internal/model/serializer"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestLogicRoleCount(t *testing.T) {
	ctx := gctx.New()
	r := role.New()
	data, err := r.Count(ctx)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)
}

func TestLogicRoleList(t *testing.T) {
	ctx := gctx.New()
	r := role.New()
	in := model.RoleListInput{
		Page:     1,
		PageSize: 10,
	}
	data, err := r.List(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)
}

func TestLogicRoleCreate(t *testing.T) {
	ctx := gctx.New()
	r := role.New()
	RoleAttr := model.RoleAttr{
		Name:   "guest",
		Router: "home",
		Pid:    -1,
	}
	in := model.RoleCreateInput{
		RoleAttr: RoleAttr,
	}
	uid, err := r.Create(ctx, &in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(uid)
}

func TestLogicRoleDetail(t *testing.T) {
	ctx := gctx.New()
	r := role.New()
	in := model.RoleDetailInput{
		Id: 5,
	}
	data, err := r.Detail(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)
}

func TestLogicRoleDelete(t *testing.T) {
	ctx := gctx.New()
	r := role.New()
	in := model.RoleDeleteInput{
		Id: 3,
	}
	data, err := r.Delete(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)
}

func TestLogicRoleUpdate(t *testing.T) {
	ctx := gctx.New()
	r := role.New()
	in := model.RoleUpdateInput{
		Id: 1,
		RoleAttr: model.RoleAttr{
			Name: "guest",
		},
	}
	id, err := r.Update(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(id)
}

func TestLogicRoleTree(t *testing.T) {
	ctx := gctx.New()
	r := role.New()
	in := model.RoleListInput{
		Page:     1,
		PageSize: 1000,
	}
	objList, err := r.GetTree(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	treeMap := make(map[int][]*serializer.RoleDetail)
	for _, v := range objList {
		treeMap[v.Pid] = append(treeMap[v.Pid], v)
	}

	// bfs map => tree
	g.Dump(treeMap)
}
