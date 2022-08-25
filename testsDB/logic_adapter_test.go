package testsDB

import (
	_ "admin/internal/logic"
	"admin/internal/service"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

// sub, obj, act 表示经典三元组: 访问实体 (Subject)，访问资源 (Object) 和访问方法 (Action)

func testGetRoles(t *testing.T, e *casbin.Enforcer, res []string, name string, domain ...string) {
	t.Helper()
	myRes, err := e.GetRolesForUser(name, domain...)
	if err != nil {
		t.Error("Roles for ", name, " could not be fetched: ", err.Error())
	}
	t.Log("Roles for ", name, ": ", myRes)

	if !util.SetEquals(res, myRes) {
		t.Error("Roles for ", name, ": ", myRes, ", supposed to be ", res)
	}
}

func TestCasbin(t *testing.T) {
	e, _ := casbin.NewEnforcer("examples/rbac_model.conf", "examples/rbac_policy.csv")

	testGetRoles(t, e, []string{"data2_admin"}, "alice")
	testGetRoles(t, e, []string{}, "bob")
	testGetRoles(t, e, []string{}, "data2_admin")
	testGetRoles(t, e, []string{}, "non_exist")
}

func TestCasbinAdapter(t *testing.T) {
	a := service.Adapter()
	e, err := casbin.NewEnforcer("examples/rbac_model.conf", a)
	if err != nil {
		t.Fatal(err)
	}
	e.LoadPolicy()

	// Check the permission.
	ans, err := e.Enforce("alice", "data1", "read")
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(ans)

}

func TestGfMap(t *testing.T) {
	data := g.Map{
		"a": 1,
	}
	g.Dump(data)
	gmap.New()
}
