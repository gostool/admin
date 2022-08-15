package tests

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestCasbin(t *testing.T) {
	e, err := casbin.NewEnforcer("./example.conf", "./example.csv")
	if err != nil {
		t.Fatal(err)
	}
	g.Dump(e)
	ans := e.AddNamedMatchingFunc("g", "KeyMatch2", util.KeyMatch2)
	g.Dump(ans)
}

func TestGfMap(t *testing.T) {
	data := g.Map{
		"a": 1,
	}
	g.Dump(data)
	gmap.New()
}
