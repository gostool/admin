package tests

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/redis-adapter/v2"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestRedis(t *testing.T) {
	// Initialize a Redis adapter and use it in a Casbin enforcer:
	a := redisadapter.NewAdapter("tcp", "127.0.0.1:6379") // Your Redis network and address.
	// Use the following if Redis has password like "123"
	//a := redisadapter.NewAdapterWithPassword("tcp", "127.0.0.1:6379", "123")
	e, _ := casbin.NewEnforcer("examples/rbac_model.conf", a)

	// Load the policy from DB.
	e.LoadPolicy()

	// Check the permission.
	ans, err := e.Enforce("alice", "data1", "read")
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(ans)

	// Modify the policy.
	//e.AddPolicy(...)
	//_, err = e.AddPolicy("alice", "data1", "read")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//

	// e.RemovePolicy(...)
	//e.RemovePolicy("alice", "data1", "read")

	//// Save the policy back to DB.
	e.SavePolicy()
}
