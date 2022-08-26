package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
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
		userGuestInit(ctx)
	case "create":
		createAdmin(ctx)
	}
	out = &cMainOutput{}
	return out, nil
}

func (c *CMain) RoleInit(ctx context.Context, in cMainRoleInput) (out *cMainOutput, err error) {
	roleInit(ctx)
	out = &cMainOutput{}
	return out, nil
}

func (c *CMain) RoleMenuInit(ctx context.Context, in cMainRoleMenuInput) (out *cMainOutput, err error) {
	menuRoleInit(ctx)
	out = &cMainOutput{}
	return out, nil
}

func (c *CMain) ApiInit(ctx context.Context, in cMainApiInput) (out *cMainOutput, err error) {
	apiInit(ctx)
	out = &cMainOutput{}
	return out, nil
}

func (c *CMain) PermInit(ctx context.Context, in cMainPermInput) (out *cMainOutput, err error) {
	permInit(ctx)
	out = &cMainOutput{}
	return out, nil
}

func (c *CMain) Init(ctx context.Context, in cMainInitInput) (out *cMainOutput, err error) {
	apiInit(ctx)
	roleInit(ctx)
	menuInit(ctx)
	userGuestInit(ctx)
	createAdmin(ctx)
	menuRoleInit(ctx)
	permInit(ctx)
	out = &cMainOutput{}
	return out, nil
}
