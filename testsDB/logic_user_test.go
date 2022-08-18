package testsDB

import (
	"admin/internal/logic/user"
	"admin/internal/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestLogicUserLogin(t *testing.T) {
	ctx := gctx.New()
	user := user.New()
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

func TestLogicUserLoginWeb(t *testing.T) {
	ctx := gctx.New()
	user := user.New()
	in := model.UserLoginWebInput{
		Name:     "guest",
		Password: "guest",
	}
	uid, err := user.LoginWeb(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(uid)
}

func TestLogicUserRegister(t *testing.T) {
	ctx := gctx.New()
	user := user.New()
	in := model.UserCreateInput{
		UserAttr: model.UserAttr{
			Name:     "alice",
			Password: "alice",
			Nickname: "shasha",
			RoleId:   1,
		},
	}
	uid, err := user.Register(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(uid)
}

func TestLogicUserCreate(t *testing.T) {
	ctx := gctx.New()
	user := user.New()
	in := model.UserCreateInput{
		UserAttr: model.UserAttr{
			Name:     "guest",
			Password: "guest",
			Nickname: "hanhan",
			RoleId:   5,
		},
	}
	uid, err := user.Create(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(uid)
}

func TestLogicUserDetail(t *testing.T) {
	ctx := gctx.New()
	user := user.New()
	in := model.UserDetailInput{
		Id: 9,
	}
	data, err := user.Detail(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)
}

func TestLogicUserDelete(t *testing.T) {
	ctx := gctx.New()
	user := user.New()
	in := model.UserDeleteInput{
		Id: 12,
	}
	data, err := user.Delete(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)
}
