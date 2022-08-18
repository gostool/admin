package tests

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"testing"
)

func TestGf(t *testing.T) {
	roleIds := []int{1, 2}
	data := gconv.String(roleIds)
	g.Dump(data)
}
