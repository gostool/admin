package testsDB

import (
	"admin/internal/logic/users"
	"admin/internal/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestLogicUserLogin(t *testing.T) {
	ctx := gctx.New()
	user := users.New()
	in := model.UserLoginInput{
		Name:     "john",
		Password: "123",
	}
	uid, err := user.Login(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(uid)
}
