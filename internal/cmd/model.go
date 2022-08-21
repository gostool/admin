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
}

type cMainUserInput struct {
	g.Meta `name:"user" brief:"init user data"`
	Act    string `v:"required" name:"action" short:"act" brief:"action:[init,create],init是初始化guest，create是创建超级管理员"`
}

type cMainRoleInput struct {
	g.Meta `name:"role" brief:"init role data"`
}

type cMainRoleMenuInput struct {
	g.Meta `name:"role-menu" brief:"init role-menu data"`
}

type cMainApiInput struct {
	g.Meta `name:"api" brief:"init api data"`
}

type cMainInitInput struct {
	g.Meta `name:"init" brief:"init role->menu->user->role-menu data"`
}
