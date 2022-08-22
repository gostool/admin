package tests

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
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

func TestGfMap(t *testing.T) {
	data := g.Map{
		"a": 1,
	}
	g.Dump(data)
	gmap.New()
}
