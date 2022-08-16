package testsDB

import (
	"admin/internal/dao"
	"admin/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gmeta"
	"testing"
)

type UserInfo struct {
	gmeta.Meta `orm:"table:user"`
	Id         int          `json:"id"`
	Name       string       `json:"name"`
	RoleId     int          `json:"roleId"`
	UserRole   *entity.Role `orm:"with:id=roleId"`
}

func TestWith(t *testing.T) {
	ctx := gctx.New()
	var user *UserInfo
	//db.Model(tableUser).WithAll().Where("id", 3).Scan(&user)
	err := dao.User.Ctx(ctx).WithAll().Where("role_id", 5).Scan(&user)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(user)
}
