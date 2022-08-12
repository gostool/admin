package v1

import "github.com/gogf/gf/v2/frame/g"

type MenuListReq struct {
	g.Meta `path:"/menu/list" method:"get"`
}
type MenuListRes struct {
}
