package main

import (
	_ "admin/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"admin/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
