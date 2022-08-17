package common

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
)

func GetVarFromCtx(ctx context.Context, key string) *gvar.Var {
	//r := g.RequestFromCtx(common)
	//if r == nil {
	//	logger.Errorf(common, "g.RequestFromCtx :%v err  r => nil\n", common)
	//	err = errors.New("internal err")
	//	return res, err
	// }
	return gvar.New(ctx.Value(key))
}
