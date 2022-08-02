package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserReq struct {
	g.Meta `path:"/login" method:"get"`
	Name   string `v:"required" dc:"Your name"`
}
type UserRes struct {
	Reply string `dc:"Reply content"`
}
