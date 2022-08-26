package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	_ "admin/internal/controller"
	_ "admin/internal/logic"
	_ "admin/internal/packed"

	"admin/internal/cmd"

	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd, err := gcmd.NewFromObject(cmd.CMain{})
	if err != nil {
		panic(err)
	}
	cmd.Run(gctx.New())
}
