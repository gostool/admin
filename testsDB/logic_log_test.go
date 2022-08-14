package testsDB

import (
	"admin/internal/logic/log"
	"admin/internal/model"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestLogicLogCount(t *testing.T) {
	ctx := gctx.New()
	r := log.New()
	data, err := r.Count(ctx)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)
}

func TestLogicLogList(t *testing.T) {
	ctx := gctx.New()
	r := log.New()
	in := model.LogListInput{
		Page:     1,
		PageSize: 10,
	}
	data, err := r.List(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)
}

func TestLogicLogCreate(t *testing.T) {
	ctx := gctx.New()
	r := log.New()
	LogAttr := model.LogAttr{
		UserId:   21,
		Status:   0,
		Extra:    "ex",
		Ip:       "1.1.1.1",
		Path:     "",
		Method:   "get",
		Request:  "req",
		Response: "res",
		Agent:    "agent",
		Latency:  10,
		ErrMsg:   "err",
		UserName: "tp",
	}
	in := model.LogCreateInput{
		LogAttr: LogAttr,
	}
	uid, err := r.Create(ctx, &in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(uid)
}

func TestLogicLogDetail(t *testing.T) {
	ctx := gctx.New()
	r := log.New()
	in := model.LogDetailInput{
		Id: 100,
	}
	data, err := r.Detail(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)
}

func TestLogicLogDelete(t *testing.T) {
	ctx := gctx.New()
	r := log.New()
	in := model.OrmDeleteInput{
		Id: 100,
	}
	data, err := r.SafeDelete(ctx, &in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(data)
}

func TestLogicLogUpdate(t *testing.T) {
	ctx := gctx.New()
	r := log.New()
	in := model.LogUpdateInput{
		Id: 100,
		LogAttr: model.LogAttr{
			UserId:   21,
			Status:   0,
			Extra:    "ex",
			Ip:       "1.1.1.1",
			Path:     "",
			Method:   "get",
			Request:  "request",
			Response: "res",
			Agent:    "agent",
			Latency:  10,
			ErrMsg:   "err",
			UserName: "tp",
		},
	}
	id, err := r.Update(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(id)
}
