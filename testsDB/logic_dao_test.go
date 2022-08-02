package testsDB

import (
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

func TestList(t *testing.T) {
	ctx := gctx.New()
	query := g.Map{
		"is_deleted": 0,
	}
	res, err := dao.Users.Ctx(ctx).Page(0, 1000).Where(query).All()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}
