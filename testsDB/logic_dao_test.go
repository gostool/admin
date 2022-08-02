package testsDB

import (
	"admin/internal/dao"
	"admin/internal/model"
	"admin/internal/model/entity"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
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
	//data := g.Map{
	//	"name":        "john",
	//	"password":    "123",
	//	"create_time": gtime.Now(),
	//	"update_time": gtime.Now(),
	//	"is_deleted":  0,
	//}
	in := model.UserRegisterInput{}
	user := &entity.Users{}
	if err := gconv.Struct(in, &user); err != nil {
		t.Fatal(err)
	}
	user.Password = "123"
	r, err := dao.Users.Ctx(ctx).Data(user).Insert()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
