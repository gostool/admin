package main

import (
	_ "admin/internal/logic"
	_ "admin/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"admin/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
