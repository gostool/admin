package testsDB

import (
	"admin/internal/dao"
	"admin/internal/model/entity"
	"context"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/database/gdb"
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
	//data := g.Map{
	//	"name":        "john",
	//	"password":    "123",
	//	"create_time": gtime.Now(),
	//	"update_time": gtime.Now(),
	//	"is_deleted":  0,
	//}
	err := dao.Users.Transaction(ctx, func(ctx context.Context, tx *gdb.TX) error {
		user := &entity.Users{}
		user.Password = "123"
		// 自动生成头像
		r, err := dao.Users.Ctx(ctx).Data(user).OmitEmpty().Save()
		t.Log(r)
		return err
	})
	t.Log(err)
	//r, err := dao.Users.Ctx(ctx).Data(data).Insert()
}
