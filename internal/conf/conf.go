package conf

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
)

var Logger *glog.Logger

func init() {
	// common
	Logger = g.Log("debug")
	ctx := gctx.New()
	Logger.Info(ctx, "logger init")
}
