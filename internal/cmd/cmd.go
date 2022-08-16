package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
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
}

func (c *CMain) Http(ctx context.Context, in cMainHttpInput) (out *cMainOutput, err error) {
	s := g.Server()
	s = RegisterRouter(s, ctx, in)
	s.SetOpenApiPath("/api.json")
	s.Run()
	return
}

func (c *CMain) MenuInit(ctx context.Context, in cMainMenuInput) (out *cMainOutput, err error) {
	return
}
