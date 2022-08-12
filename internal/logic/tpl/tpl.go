package tpl

import (
	"admin/internal/consts"
	"admin/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

type sTpl struct {
}

var logger *glog.Logger

func init() {
	logger = g.Log(consts.LoggerDebug)
	instance := New()
	service.RegisterTpl(instance)
}

func New() *sTpl {
	return &sTpl{}
}

// Logout 注销
func (s *sTpl) Logout(ctx context.Context) error {
	//return service.Session().RemoveUser(ctx)
	return nil
}
