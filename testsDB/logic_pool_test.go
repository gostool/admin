package testsDB

import (
	"admin/internal/consts"
	"admin/internal/logic/pool"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"sync"
	"testing"
)

func TestLogicPool(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan g.Map, 2)
	ctx := gctx.New()
	instance := pool.New(10, "pool_test")
	instance.Add(ctx, func(ctx context.Context) {
		defer wg.Done()
		t.Log("111")
		data := g.Map{}
		data["code"] = consts.PoolWorkOk
		ch <- data
	})
	instance.Add(ctx, func(ctx context.Context) {
		defer wg.Done()
		t.Log("222")
		data := g.Map{}
		data["code"] = consts.PoolWorkErr
		ch <- data
	})
	instance.Show(ctx)
	wg.Wait()
	instance.Show(ctx)
	close(ch)
	for response := range ch {
		g.Dump(response)
		//if gconv.Int(response["code"]) != consts.PoolWorkOk {
		//	msg := gconv.String(response["msg"])
		//	t.Errorf("%v\n", msg)
		//}
	}
}
