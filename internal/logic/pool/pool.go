package pool

import (
	"admin/internal/consts"
	"admin/internal/model"
	"admin/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/grpool"
)

type sPool struct {
	Pool *grpool.Pool
}

var logger *glog.Logger

func init() {
	logger = g.Log(consts.LoggerDebug)
	instance := New()
	service.RegisterPool(instance)
}

func New() *sPool {
	return &sPool{
		Pool: grpool.New(3),
	}
}

func (s *sPool) Invoke(ctx context.Context, r *model.LogCreateInput) {
	s.Pool.Add(ctx, func(ctx context.Context) {
		//写入日志数据
		id, err := service.Log().Create(ctx, r)
		if err != nil {
			logger.Warningf(ctx, "op write err:%v", err)
		} else {
			logger.Debugf(ctx, "op write ok:%v", id)
		}
	})
}

func (s *sPool) Jobs() int {
	opJobs := s.Pool.Jobs()
	return opJobs
}
