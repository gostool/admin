package pool

import (
	"admin/internal/consts"
	"admin/internal/model"
	"admin/internal/service"
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/grpool"
	"github.com/gogf/gf/v2/util/gconv"
)

type sPool struct {
	Name string
	Pool *grpool.Pool
}

var logger *glog.Logger

func init() {
	logger = g.Log(consts.LoggerDebug)
	instance := New(10, "pool-task")
	service.RegisterPool(instance)
}

func New(n int, name string) *sPool {
	if n < 3 {
		// 3是默认给日志中间件使用的
		n = 3
	}
	return &sPool{
		Name: name,
		Pool: grpool.New(n),
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
	return s.Pool.Jobs()
}

func (s *sPool) Cap() int {
	return s.Pool.Cap()
}

func (s *sPool) Size() int {
	return s.Pool.Size()
}

func (s *sPool) Add(ctx context.Context, f func(ctx context.Context)) error {
	return s.Pool.Add(ctx, f)
}

func (s *sPool) Idea() int {
	return s.Cap() - s.Size()
}

func (s *sPool) Show(ctx context.Context) {
	logger.Debugf(ctx, "name:%v cap:%v size:%v job:%v idea:%v", s.Name, s.Cap(), s.Size(), s.Jobs(), s.Idea())
	return
}

func jobDemo(ctx context.Context, ch chan<- g.Map, arg interface{}) {
	data := g.Map{
		"code": consts.PoolWorkOk,
		"msg":  "",
		"type": "",
		"data": nil,
	}
	defer func() {
		ch <- data
		if err := recover(); err != nil {
			logger.Error(ctx, err)
		}
	}()
	// do serve and deal err
	data["data"] = g.Map{"data": "demo"}
	return
}

func makeChannelResponse(ch <-chan g.Map) (data g.Map, err error) {
	data = g.Map{}
	for response := range ch {
		if gconv.Int(response["code"]) != consts.PoolWorkOk {
			msg := gconv.String(response["msg"])
			return data, errors.New(msg)
		}
		// data: 必须为map
		//curData := gconv.Map(response["data"])
		//jobType := gconv.Int(response["type"])
		g.Dump(response)
	}
	return data, nil
}
