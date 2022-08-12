package v1

import "github.com/gogf/gf/v2/frame/g"

type RoleListReq struct {
	g.Meta `path:"/role/list" method:"get" tags:"RoleService"`
}
type RoleListRes struct {
}
