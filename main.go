package main

import (
	_ "admin/internal/controller"
	_ "admin/internal/logic"
	_ "admin/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gcmd"

	"github.com/gogf/gf/v2/os/gctx"

	"admin/internal/cmd"
)

func main() {
	cmd, err := gcmd.NewFromObject(cmd.CMain{})
	if err != nil {
		panic(err)
	}
	cmd.Run(gctx.New())
}
