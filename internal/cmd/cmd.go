package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

func (c *CMain) Http(ctx context.Context, in cMainHttpInput) (out *cMainOutput, err error) {
	s := g.Server()
	s = RegisterRouter(s, ctx, in)
	s.SetOpenApiPath("/api.json")
	s.Run()
	out = &cMainOutput{}
	return out, nil
}

func (c *CMain) MenuInit(ctx context.Context, in cMainMenuInput) (out *cMainOutput, err error) {
	menuInit(ctx)
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

func (c *CMain) RoleInit(ctx context.Context, in cMainRoleInput) (out *cMainOutput, err error) {
	RoleInit(ctx)
	out = &cMainOutput{}
	return out, nil
}

func (c *CMain) RoleMenuInit(ctx context.Context, in cMainRoleMenuInput) (out *cMainOutput, err error) {
	menuRoleInit(ctx)
	out = &cMainOutput{}
	return out, nil
}
