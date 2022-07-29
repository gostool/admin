package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HelloReq struct {
	g.Meta `path:"/login" method:"get"`
	Name   string `v:"required" dc:"Your name"`
}
type HelloRes struct {
	Reply string `dc:"Reply content"`
}
