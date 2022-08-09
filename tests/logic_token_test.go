package tests

import (
	jwt "admin/internal/logic/token"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
	"time"
)

func TestGenToken(t *testing.T) {
	ctx := gctx.New()
	uid := "1"
	exp := time.Now().Add(time.Duration(60*60) * time.Second).Unix()
	token, err := jwt.New().GenToken(ctx, uid, int64(exp))
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(token)
}

func TestCheckToken(t *testing.T) {
	ctx := gctx.New()
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiIxIiwiZXhwIjoxNjYwMDQ0MTAzLCJpc3MiOiIxIn0.b7m3hHtvueFLhytLOoT8alNc4GSPrA_5PHPlasAfg6A"
	uid, err := jwt.New().CheckToken(ctx, token)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(uid)
}
