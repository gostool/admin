package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

type CMain struct {
	g.Meta `name:"main" brief:"start main"`
}

type cMainOutput struct{}

type cMainHttpInput struct {
	g.Meta `name:"http" brief:"start http server"`
}

type cMainMenuInput struct {
	g.Meta `name:"menu" brief:"init menu data"`
	Act    string `v:"required" name:"action" short:"act" brief:"action"`
}

type cMainUserInput struct {
	g.Meta `name:"user" brief:"init user data"`
	Act    string `v:"required" name:"action" short:"act" brief:"action"`
}

func (c *CMain) Http(ctx context.Context, in cMainHttpInput) (out *cMainOutput, err error) {
	s := g.Server()
	s = RegisterRouter(s, ctx, in)
	s.SetOpenApiPath("/api.json")
	s.Run()
	out = &cMainOutput{}
	return out, nil
}

func (c *CMain) MenuInit(ctx context.Context, in cMainMenuInput) (out *cMainOutput, err error) {
	switch in.Act {
	case "init":
		menuInit(ctx)
	case "roleMenuInit":
		menuRoleInit(ctx)
	}
	out = &cMainOutput{}
	return out, nil
}

func (c *CMain) UserInit(ctx context.Context, in cMainUserInput) (out *cMainOutput, err error) {
	switch in.Act {
	case "init":
		userInit(ctx)
	case "create":
		name := gcmd.Scan("> What's your super user name?\n")
		password := gcmd.Scan("> What's your super user password?\n")
		nickname := gcmd.Scan("> What's your super user nickname?\n")
		userCreateAdmin(ctx, name, password, nickname)
	}
	out = &cMainOutput{}
	return out, nil
}
