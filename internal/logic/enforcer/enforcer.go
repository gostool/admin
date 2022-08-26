package enforcer

import (
	"admin/internal/consts"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

type sEnforcer struct {
}

var logger *glog.Logger

func init() {
	logger = g.Log(consts.LoggerDebug)
	//instance := New()
	//service.RegisterEnforcer(instance)
}

func New() *sEnforcer {
	return &sEnforcer{}
}
