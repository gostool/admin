package testsDB

import (
	"admin/internal/logic/tools"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestLogicToolsCapture(t *testing.T) {
	ctx := gctx.New()
	s := tools.New()
	id, b64s, err := s.Captcha(ctx)
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(id, b64s)
}
