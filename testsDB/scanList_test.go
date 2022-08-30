package testsDB

import (
	"admin/internal/dao"
	"admin/internal/model/entity"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestScan(t *testing.T) {
	ctx := gctx.New()
	// 定义用户列表
	var user entity.User
	// 查询用户基础数据
	// SELECT * FROM `user` WHERE `name`='guest'
	err := dao.User.Ctx(ctx).Scan(&user, "name", "guest")
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(user)
	// 查询role详情数据
	// SELECT * FROM `role` WHERE `id`=1
	var role entity.Role
	err = dao.Role.Ctx(ctx).Scan(&role, "id", user.RoleId)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(role)
}

// 组合模型，用户信息
type Entity struct {
	User *entity.User
	Role *entity.Role
}

func TestScanList(t *testing.T) {
	ctx := gctx.New()
	// 定义用户列表
	var users []Entity
	// 查询用户基础数据
	// SELECT * FROM `user`
	err := dao.User.Ctx(ctx).ScanList(&users, "User")
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(users)
	// 查询用户详情数据
	// SELECT * FROM `role` WHERE `id` IN(1,2)
	query := gdb.ListItemValuesUnique(users, "User", "RoleId")
	g.Dump(query)
	err = dao.Role.Ctx(ctx).
		Where("id", query).
		ScanList(&users, "Role", "User", "id:RoleId")
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(users)
}
