package testsDB

import (
	"admin/internal/consts"
	"admin/internal/dao"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestOrmUser(t *testing.T) {
	ctx := gctx.New()
	n, err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Id, 1).Count()
	if err != nil {
		t.Fatal()
	}
	t.Log(n)
}
func TestOrmUserInsert(t *testing.T) {
	ctx := gctx.New()
	data := g.Map{
		"name":       "john",
		"password":   "123",
		"is_deleted": 0,
	}
	r, err := dao.Users.Ctx(ctx).Data(data).Insert()
	if err != nil {
		t.Fatal(err)
	}
	uid, err := r.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(uid)
}

func TestOrmUserList(t *testing.T) {
	ctx := gctx.New()
	query := g.Map{
		"is_deleted": consts.CREATED,
	}
	res, err := dao.Users.Ctx(ctx).Page(0, 1000).Where(query).All()
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(res)
}

func TestOrmUserDelete(t *testing.T) {
	ctx := gctx.New()
	query := g.Map{
		"id": 2,
	}
	res, err := dao.Users.Ctx(ctx).Unscoped().Where(query).Delete()
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(res)
}

func TestOrmUserUpdate(t *testing.T) {
	ctx := gctx.New()
	//dao.User.Data(g.Map{"is_deleted": model.DELETED}).Where("id=?", pk).Update()
	query := g.Map{
		"id": 6,
	}
	update := g.Map{"is_deleted": consts.DELETED}
	res, err := dao.Users.Ctx(ctx).Where(query).Update(update)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(res)
}
