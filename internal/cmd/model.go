package cmd

import "github.com/gogf/gf/v2/frame/g"

type CMain struct {
	g.Meta `name:"main" brief:"start main"`
}

type cMainOutput struct{}

type cMainHttpInput struct {
	g.Meta `name:"http" brief:"start http server"`
}

type cMainMenuInput struct {
	g.Meta `name:"menu" brief:"init menu data"`
	Act    string `v:"required" name:"action" short:"act" brief:"action:[init, create]"`
}

type cMainUserInput struct {
	g.Meta `name:"user" brief:"init user data"`
	Act    string `v:"required" name:"action" short:"act" brief:"action"`
}

type cMainRoleInput struct {
	g.Meta `name:"role" brief:"init role data"`
}